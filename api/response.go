package api

import (
	"encoding/json"
	"io"
)

// code in response data, not http.status_code
const (
	CodeSuccess = 20000
	// 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
	CodeIllegalToken         = 50008
	CodeOtherClientsLoggedIn = 50012
	CodeTokenExpired         = 50014
)

// Code: custom code type, not http status code
// Data: data to show
// Message: message to explain code
type CustomResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseV1(code int, message string, data interface{}) (customResp CustomResponse) {
	customResp = CustomResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return
}

func DecodeResponseV1(buffer io.Reader) CustomResponse {
	var data CustomResponse
	decoder := json.NewDecoder(buffer)
	decoder.Decode(&data)
	return data
}
