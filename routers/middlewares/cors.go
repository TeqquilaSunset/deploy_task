package middlewares

import (
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// CORSMiddleware returns a gin.HandlerFunc that handles CORS policies.
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")

		if isAllowedOrigin(origin) {
			context.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			// Optionally, you can block the request by not setting the header
			// This is a more restrictive approach, which might involve handling the request differently
			context.Writer.Header().Set("Access-Control-Allow-Origin", "null") // or some default value
		}

		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, api-key, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Cache-Control", "no-cache")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
		} else {
			context.Next()
		}
	}
}

// isAllowedOrigin checks if the provided origin is allowed
func isAllowedOrigin(origin string) bool {
	var allowedOrigins []string 
	var origins string = viper.GetString("ALLOWED_ORIGINS")

	if origins != "" {
		allowedOrigins = strings.Split(origins, ",")
	} else {
		return false
	}

	if origin == "" {
		return false
	}

	// Allow all origins if wildcard (*) is present
	for _, o := range allowedOrigins {
		if o == "*" {
			return true
		}
	}

	// Check if the origin is in the allowed list
	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}

	return false
}