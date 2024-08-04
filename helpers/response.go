package helpers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Status struct {
    Code  string `json:"code"`
    Error bool `json:"error"`
    Route string `json:"route"`
	Message string `json:"message"`
}

type Response struct {
    Status   Status      `json:"status"`
    Response interface{} `json:"response"`
}

// APIResponse handles the response formatting and error checking
func FormatAPIResponse(context *gin.Context, code int, data any) {
    route := context.Request.URL.Path
    var statusResponse string
    var isError bool
    var message string

	err, ok := data.(error);

	if code >= 400 {
		statusResponse = GetStatusMessage(code)
		isError = true

		if ok {
			data = nil
            message = err.Error()
		} else {
            message = statusResponse
        }

	} else if ok { // Check if data contains an error
        statusResponse = "Error"
        isError = true
        message = err.Error()
        data = nil  // Set data to nil if there's an error
        code = http.StatusBadRequest  // Set HTTP status code to 400 for errors
    } else {
        statusResponse = "Ok"
        isError = false
        message = ""
    }

    // Build the response
    status := Status{
        Code:    statusResponse,
        Error:   isError,
        Route:   route,
        Message: message,
    }

    response := Response{
        Status:   status,
        Response: data,
    }

    context.JSON(code, response)
}