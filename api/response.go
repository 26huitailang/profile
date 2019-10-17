package api

import (
	"encoding/json"
	"io"
)

type CustomResponse struct {
	Data interface{} `json:"data"`
	Info string      `json:"info"`
}

func ResponseV1(info string, data interface{}) CustomResponse {
	msg := CustomResponse{
		Info: info,
		Data: data,
	}
	return msg
}

func DecodeResponseV1(buffer io.Reader) CustomResponse {
	var data CustomResponse
	decoder := json.NewDecoder(buffer)
	decoder.Decode(&data)
	return data
}
