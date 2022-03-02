package handler

import (
	"fmt"
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

func apiResponseErr(c echo.Context, status int, message string) error {
	return c.JSON(status, apiErrorResponse{Code: fmt.Sprintf("%d", status), Errors: []string{message}})
}
