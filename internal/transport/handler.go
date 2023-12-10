package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dmytrodemianchuk/bank-transactions/docs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
	v1 "github.com/dmytrodemianchuk/bank-transactions/internal/transport/v1"
)

type IServicesBank interface {
	GetFilteredData(ctx context.Context, input domain.FilterSearchInput) ([]domain.Transaction, error)
	FetchExternTransactions(ctx context.Context, url string) (domain.Status, error)
}
type IServicesRemote interface {
	Get(ctx context.Context, from, to *int) ([]domain.Transaction, error)
}

type Handler struct {
	servicesBank   IServicesBank
	servicesRemote IServicesRemote
}

func NewHandler(servicesBank IServicesBank, servicesRemote IServicesRemote) *Handler {
	return &Handler{
		servicesBank:   servicesBank,
		servicesRemote: servicesRemote,
	}
}

func (h *Handler) Init(cfg *domain.Config) *gin.Engine {
	router := gin.Default()

	pprof.Register(router) // http://localhost:8080/debug/pprof/

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI_v1(router)

	return router
}

func (h *Handler) initAPI_v1(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.servicesBank, h.servicesRemote)
	api := router.Group("/api/v1")
	{
		handlerV1.Init(api)
	}
}
