package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	serverErrMsg = "エラーが発生しました"
	clientErrMsg = "不正なリクエストです"
)

type apiErrorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func apiResponseOK(c echo.Context, data interface{}) error {
	return c.JSONPretty(http.StatusOK, data, " ")
}

func apiResponseErr(c echo.Context, data interface{}, mode string) error {
	switch mode {
	case "server":
		return c.String(http.StatusInternalServerError, serverErrMsg)
	case "client":
		return c.String(http.StatusBadRequest, clientErrMsg)
	default:
		return c.String(http.StatusBadGateway, serverErrMsg)
	}
}
