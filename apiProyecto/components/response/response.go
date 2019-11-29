package response

import (
	"net/http"

	"github.com/labstack/echo"
)

func Success(c echo.Context, payload interface{}) error {
	return SuccessWithCode(c, 200, payload)
}

func SuccessWithCode(c echo.Context, statusCode int, payload interface{}) error {
	finalPayload := map[string]interface{}{"Resultado": payload}
	return c.JSON(statusCode, finalPayload)
}

func Error(c echo.Context, statusCode int, message string) error {
	errorPlayload := map[string]interface{}{"code": statusCode, "message": message}
	finalPayload := map[string]interface{}{"error": errorPlayload}
	return c.JSON(statusCode, finalPayload)
}

func ErrorKey(c echo.Context, statusCode int, key string) error {
	return Error(c, statusCode, "Error in Headers")

}

func ErrorBadRequestWithKey(c echo.Context, key string) error {
	return Error(c, http.StatusBadRequest, key)
}
