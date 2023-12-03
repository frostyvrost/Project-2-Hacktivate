package dto

import (
	"net/http"
	"project-2/model"
	"project-2/pkg"
	"project-2/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(context *gin.Context) {
	var photo model.Photo

	if err := context.ShouldBindJSON(&photo); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	result, err := service.PhotoService.CreatePhoto(&photo, userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         result.ID,
		"title":      result.Title,
		"caption":    result.Caption,
		"photo_url":  result.PhotoURL,
		"user_id":    result.UserID,
		"created_at": result.CreatedAt,
	})
}

func UpdatePhoto(context *gin.Context) {
	var photo model.PhotoUpdate

	if err := context.ShouldBindJSON(&photo); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	photoIDInt, _ := strconv.Atoi(context.Param("photoId"))
	photoID := uint(photoIDInt)

	result, err := service.PhotoService.UpdatePhoto(&photo, photoID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"title":      result.Title,
		"caption":    result.Caption,
		"photo_url":  result.PhotoURL,
		"user_id":    result.UserID,
		"updated_at": result.CreatedAt,
	})
}

func GetAllPhotos(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	photos, err := service.PhotoService.GetAllPhotos(userId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	var photoMaps []map[string]interface{}

	for _, photo := range photos {
		photoMap := map[string]interface{}{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoURL,
			"user_id":    photo.UserID,
			"created_at": photo.CreatedAt,
			"updated_at": photo.UpdatedAt,
			"User": map[string]interface{}{
				"email":    photo.User.Email,
				"username": photo.User.Username,
			},
		}

		photoMaps = append(photoMaps, photoMap)
	}

	context.JSON(http.StatusOK, photoMaps)
}

func DeletePhoto(context *gin.Context) {
	id, _ := pkg.GetIdParam(context, "photoId")

	err := service.PhotoService.DeletePhoto(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
