package mainresponse

type DataResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type DefaultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
