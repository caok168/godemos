package e

import (
	"net/http"
)

type JSONResult struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"status bad request"`
	Data    interface{} `json:"data"`
}

func Result(data interface{}) *JSONResult {
	return &JSONResult{Code: ErrorOK, Message: "", Data: data}
}

func Error(code int) *JSONResult {
	var message string
	if code > 0 && code < 1000 {
		message = http.StatusText(code)
	} else {
		msg, ok := errorText[code]
		if ok {
			message = msg
		} else {
			message = errorText[ErrorUnknown]
		}
	}
	return &JSONResult{Code: code, Message: message, Data: nil}
}

func ErrorText(code int, message string) *JSONResult {
	return &JSONResult{Code: code, Message: message, Data: nil}
}
