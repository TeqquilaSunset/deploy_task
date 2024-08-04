package services

import (
	"context"
	"errors"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imshawan/gin-backend-starter/helpers"
	"github.com/imshawan/gin-backend-starter/infra/database"
	"github.com/imshawan/gin-backend-starter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignIn (ctx *gin.Context) {
	var userReq models.UserRequest
	var field string = "username"

	// Bind and validate the request body
	if err := ctx.ShouldBind(&userReq); err != nil {
		helpers.FormatAPIResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if ok := helpers.IsEmail(userReq.Username); ok {
		field = "email"
	}

	usersCollection := database.Mongo.Collection("users")

	var existingUser models.User
	if err := usersCollection.FindOne(context.TODO(), bson.M{field: userReq.Username}).Decode(&existingUser); err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.FormatAPIResponse(ctx, http.StatusUnauthorized, errors.New("invalid credentials"))
			return
		}
		helpers.FormatAPIResponse(ctx, http.StatusConflict, err)
		return
	}

	match, compareErr := helpers.ComparePassword(userReq.Password, existingUser.PasswordHash) 
	if compareErr != nil {
		helpers.FormatAPIResponse(ctx, http.StatusBadRequest, compareErr)
	}

	if !match {
		helpers.FormatAPIResponse(ctx, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}
	token, err := helpers.SignJWTToken(existingUser)
	if err != nil {
		helpers.FormatAPIResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user": existingUser,
	}

	helpers.FormatAPIResponse(ctx, http.StatusOK, response)
}