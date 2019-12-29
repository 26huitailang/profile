package api

import (
	"encoding/json"
	"io"
)

const (
	CodeSuccess = 20000
	// 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
	CodeIllegalToken         = 50008
	CodeOtherClientsLoggedIn = 50012
	CodeTokenExpired         = 50014
)

type CustomResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseV1(code int, info string, data interface{}) CustomResponse {
	msg := CustomResponse{
		Code:    code,
		Message: info,
		Data:    data,
	}
	return msg
}

func DecodeResponseV1(buffer io.Reader) CustomResponse {
	var data CustomResponse
	decoder := json.NewDecoder(buffer)
	decoder.Decode(&data)
	return data
}
