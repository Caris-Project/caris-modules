package utils

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Caris-Project/caris-modules/locale"
)

func sendResponse(c echo.Context, httpCode int, success bool, data interface{}, message string, code int) error {
	if data == nil {
		data = echo.Map{}
	}

	return c.JSON(httpCode, echo.Map{
		"success": success,
		"data":    data,
		"message": message,
		"code":    code,
	})
}

// Response200 response success
func Response200(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = locale.CommonKeySuccess
	}

	localeData := locale.GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusOK, true, data, localeData.Message, localeData.Code)
}

// Response400 bad request
func Response400(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = locale.CommonKeyBadRequest
	}

	localeData := locale.GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusBadRequest, false, data, localeData.Message, localeData.Code)
}

// Response401 unauthorized
func Response401(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = locale.CommonKeyUnauthorized
	}

	localeData := locale.GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusUnauthorized, false, data, localeData.Message, localeData.Code)
}

// Response404 not found
func Response404(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = locale.CommonKeyNotFound
	}

	localeData := locale.GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusNotFound, false, data, localeData.Message, localeData.Code)
}

// ResponseInvalidChecksum for invalid checksum
func ResponseInvalidChecksum(c echo.Context) error {
	key := locale.CommonKeyInvalidChecksum
	localeData := locale.GetByKey(key)
	return sendResponse(c, http.StatusBadRequest, false, echo.Map{}, localeData.Message, localeData.Code)
}

// ResponseRouteValidation ...
// WARNING: only match with lib "github.com/go-ozzo/ozzo-validation/v4"
func ResponseRouteValidation(c echo.Context, err error) error {
	errors := strings.Split(err.Error(), ";")
	totalErrors := len(errors)

	// If have no errors
	if totalErrors == 0 {
		return Response400(c, nil, "")
	}

	// pretty.Println("errors[0]", errors[0])

	key := ""

	if !strings.Contains(errors[0], "(") {
		// Get message
		key = strings.Split(errors[0], ":")[1]
		// Trim unnecessary spaces
		key = strings.Trim(key, " ")
		// Replace char "." for the last value
		key = strings.Replace(key, ".", "", -1)
	} else {
		// Get message
		key = strings.Split(errors[0], ":")[2]
		// Trim unnecessary spaces
		key = strings.Trim(key, " ")
		// Replace char "."
		key = strings.Replace(key, ".", "", -1)
		// Replace char ")"
		key = strings.Replace(key, ")", "", -1)
	}

	// pretty.Println("validation error key -", key)

	// Return
	return Response400(c, nil, key)
}
