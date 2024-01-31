package errorresponse

import mainresponse "github.com/baskararestu/transfer-money/responses/mainResponse"

func NewErrorResponse(message string) mainresponse.DefaultResponse {
	return mainresponse.DefaultResponse{
		Success: false,
		Message: message,
	}
}
