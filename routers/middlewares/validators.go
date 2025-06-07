package middlewares

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/imshawan/gin-backend-starter/helpers"
)

func ValidateRequestFields(model interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Сохраняем тело запроса в буфер, чтобы можно было прочитать его несколько раз
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			helpers.FormatAPIResponse(ctx, http.StatusBadRequest, err)
			ctx.Abort()
			return
		}

		// Восстанавливаем тело запроса для последующих вызовов
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Первый вызов ShouldBind (валидация)
		if err := ctx.ShouldBind(model); err != nil {
			var validationError error

			if errs, ok := err.(validator.ValidationErrors); ok {
				var validationErrors []string
				for _, e := range errs {
					field := e.Field()
					tag := e.Tag()
					validationErrors = append(validationErrors, fmt.Sprintf("%s (%s)", field, tag))
				}
				errorMessage := fmt.Sprintf("validation failed: %s", strings.Join(validationErrors, ", "))
				validationError = errors.New(errorMessage)
			} else {
				validationError = err
			}

			helpers.FormatAPIResponse(ctx, http.StatusBadRequest, validationError)
			ctx.Abort()
			return
		}

		// Восстанавливаем тело запроса снова, чтобы следующий ShouldBind мог прочитать данные
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Если валидация прошла успешно, продолжаем
		ctx.Next()
	}
}
