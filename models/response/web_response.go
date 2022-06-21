package response

type WebResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}
