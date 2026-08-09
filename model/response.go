package model

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type RestResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type EmptyStruct struct{}
type EmptySlice []struct{}

type RestResponder interface {
	Success(msg string, data any) error
	Fail(msg string) error
	Err(httpCode, code int, msg string) error
}

type EchoResponder struct {
	ctx *echo.Context
}

func UseEchoResponder(ctx *echo.Context) EchoResponder {
	return EchoResponder{ctx: ctx}
}

func (r EchoResponder) Success(msg string, data any) error {
	return r.ctx.JSON(http.StatusOK, RestResponse{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func (r EchoResponder) Fail(msg string) error {
	return r.ctx.JSON(http.StatusOK, RestResponse{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
}

func (r EchoResponder) Err(httpCode, code int, msg string) error {
	return r.ctx.JSON(httpCode, RestResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
