package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/wensboy/sss/model"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
)

type EchoApiHandler interface {
	ScalarWrapHandler(*echo.Context) error
}

type echoApiHandler struct {
}

func NewEchoApiHandler() EchoApiHandler {
	return &echoApiHandler{}
}

func (h *echoApiHandler) ScalarWrapHandler(c *echo.Context) error {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "./api/docs/swagger.json",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "SSS API Reference",
		},
		DarkMode: true,
	})
	if err != nil {
		return model.UseEchoResponder(c).Err(http.StatusInternalServerError, 1, err.Error())
	}
	return c.HTML(http.StatusOK, htmlContent)
}
