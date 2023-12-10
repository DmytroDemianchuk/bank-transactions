package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
)

//go:generate mockgen -source=evo.go -destination=mocks/mock.go

type IServicesBank interface {
	GetFilteredData(ctx context.Context, input domain.FilterSearchInput) ([]domain.Transaction, error)
	FetchExternTransactions(ctx context.Context, url string) (domain.Status, error)
}
type IServicesRemote interface {
	Get(ctx context.Context, from, to *int) ([]domain.Transaction, error)
}

type Handler struct {
	servicesEVO    IServicesBank
	servicesRemote IServicesRemote
}

func NewHandler(servicesEVO IServicesBank, servicesRemote IServicesRemote) *Handler {
	return &Handler{
		servicesEVO:    servicesEVO,
		servicesRemote: servicesRemote,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	api.GET("/get_csv_mock_remote_service", h.getSourceFileCSV_as_MockRemoteService)

	api.POST("/download_remote_transactions", h.downloadRemoteTransactionsCSV)

	filtered := api.Group("/filtered")
	{
		filtered.POST("/csv", h.getFilteredFileCSV)
		filtered.POST("/json", h.getFilteredDataJSON)
	}
}

func buildCSV(transactions []domain.Transaction) *strings.Reader {
	str, _ := gocsv.MarshalString(&transactions)
	return strings.NewReader(str)
}

// @Summary Request filtered csv file
// @Tags Services
// @ID get-filtered-csv
// @Param   input body domain.FilterSearchInput true " "
// @Success 200
// @Success 204
// @Failure 400   {object} domain.ErrorResponse
// @Failure 500   {object} domain.ErrorResponse
// @Router /api/v1/filtered/csv/ [post]
func (h *Handler) getFilteredFileCSV(ctx *gin.Context) {
	var input domain.FilterSearchInput

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	v := validator.New()
	if err := v.Struct(input); err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("data validation error"))
		return
	}

	transactions, err := h.servicesEVO.GetFilteredData(context.Background(), input)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, errors.New("service failure"))
		return
	}

	if len(transactions) == 0 {
		logrus.Error(err)
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	CSV := buildCSV(transactions)

	headers := map[string]string{
		"Content-Disposition": `attachment; filename="source.csv"`,
	}

	ctx.DataFromReader(http.StatusOK, -1, "text/html; charset=UTF-8", CSV, headers)
}

// @Summary Request filtered json
// @Tags Services
// @ID get-filtered-json
// @Param   input body     domain.FilterSearchInput true " "
// @Success 200   {object} []domain.Transaction
// @Success 204
// @Failure 400   {object} domain.ErrorResponse
// @Failure 500   {object} domain.ErrorResponse
// @Router /api/v1/filtered/json/ [post]
func (h *Handler) getFilteredDataJSON(ctx *gin.Context) {
	var input domain.FilterSearchInput

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	v := validator.New()
	if err := v.Struct(input); err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("data validation error"))
		return
	}

	transactions, err := h.servicesEVO.GetFilteredData(context.Background(), input)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, errors.New("service failure"))
		return
	}

	if len(transactions) == 0 {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

// @Summary Request to download remote transactions
// @Description  The request runs fake ~40 seconds. If url=” or download in progress, returns the status of the download.
// @Description
// @Description  The amount of memory consumed depends on the transactionCount and workerCount in the internal\services\worker_pool.go file.
// @Tags Services
// @ID request-download-remote-transactions
// @Param   input body     domain.UrlInput true " "
// @Success 200   {object} domain.StatusResponse
// @Success 204
// @Failure 400   {object} domain.ErrorResponse
// @Failure 500   {object} domain.ErrorResponse
// @Router /api/v1/download_remote_transactions/ [post]
func (h *Handler) downloadRemoteTransactionsCSV(ctx *gin.Context) {
	var input domain.UrlInput

	if err := ctx.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	status, err := h.servicesEVO.FetchExternTransactions(context.Background(), *input.URL)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, errors.New("service failure"))
		return
	}

	ctx.JSON(http.StatusOK, domain.StatusResponse{Status: status})
}

// @Summary Test service: Gives a CSV file with initial transactions
// @Tags Mock remote service
// @ID getSourceFileCSV_as_MockRemoteService-csv
// @Param from	query integer  	false 	"From transaction, example: 1"
// @Param to	query integer  	false 	"To transaction, example: 5 (TO must be greater than FROM, if both are present at the same time)"
// @Success 200
// @Success 204
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/v1/get_csv_mock_remote_service/ [get]
func (h *Handler) getSourceFileCSV_as_MockRemoteService(ctx *gin.Context) {

	query := ctx.Request.URL.Query()
	from := query.Get("from")
	to := query.Get("to")

	var fromPtr, toPtr *int

	f, err := strconv.Atoi(from)
	if err == nil {
		fromPtr = &f
	}

	t, err := strconv.Atoi(to)
	if err == nil {
		toPtr = &t
	}

	if f < 0 || t < 0 || (fromPtr != nil && toPtr != nil && f > t) {
		logrus.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, errors.New("data validation error"))
		return
	}

	transactions, err := h.servicesRemote.Get(ctx, fromPtr, toPtr)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if len(transactions) == 0 {
		newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("no data"))
		return
	}

	CSV := buildCSV(transactions)

	headers := map[string]string{
		"Content-Disposition": `attachment; filename="source.csv"`,
	}

	ctx.DataFromReader(http.StatusOK, -1, "text/html; charset=UTF-8", CSV, headers)
}
