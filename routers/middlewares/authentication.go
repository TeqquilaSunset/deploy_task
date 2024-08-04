package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imshawan/gin-backend-starter/helpers"
	"github.com/imshawan/gin-backend-starter/infra/database"
	"github.com/imshawan/gin-backend-starter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
        authHeader := ctx.GetHeader("Authorization")

        // Check if the header contains a Bearer token
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			helpers.FormatAPIResponse(ctx, http.StatusUnauthorized, errors.New("authorization header is missing or invalid"))
            ctx.Abort()
            return
        }

        // Extract the token
        token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helpers.ValidateJWTToken(token)
		if err != nil {
			helpers.FormatAPIResponse(ctx, http.StatusUnauthorized, err)
            ctx.Abort()
            return
		}

		usersCollection := database.Mongo.Collection("users")

		var existingUser models.User
		if err := usersCollection.FindOne(context.TODO(), bson.M{"_id": claims.ID}).Decode(&existingUser); err != nil {
			helpers.FormatAPIResponse(ctx, http.StatusForbidden, errors.New("could not find user associated with this token"))
			return
		}

		ctx.Set("User", existingUser)

        // Proceed to the next middleware or handler
        ctx.Next()
	}
}