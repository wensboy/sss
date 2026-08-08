package model

type RestResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type EmptyStruct struct{}
type EmptySlice []struct{}

type RestResponder interface {
	Success(msg string, data any)
	Fail(msg string)
	Err(code int, msg string)
}
