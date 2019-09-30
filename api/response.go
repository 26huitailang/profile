package api

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
