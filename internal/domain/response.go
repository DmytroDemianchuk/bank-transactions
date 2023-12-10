package domain

type Status string

const (
	Unknown       Status = "unknown"
	Processing    Status = "in progress"
	Skip          Status = "skip request: download in progress"
	DownloadOk    Status = "successfully"
	DownloadError Status = "error"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type StatusResponse struct {
	Status Status `json:"last_download_status" swaggertype:"string" enums:"unknown,in progress,skip request: download in progress,successfully,error" example:"unknown"`
}
