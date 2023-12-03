package dto

import (
	"net/http"
	"project-2/model"
	"project-2/pkg"
	"project-2/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateComment(context *gin.Context) {
	var comment model.Comment

	if err := context.ShouldBindJSON(&comment); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	createComment, err := service.CommentService.CreateComment(&comment, userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         createComment.ID,
		"message":    createComment.Message,
		"photo_id":   createComment.PhotoID,
		"user_id":    createComment.UserID,
		"created_at": createComment.CreatedAt,
	})
}

func GetComment(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	results, err := service.CommentService.GetComment(userId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	comments := make([]gin.H, 0, len(results))

	for _, result := range results {
		comment := gin.H{
			"id":         result.ID,
			"message":    result.Message,
			"photo_id":   result.PhotoID,
			"user_id":    result.UserID,
			"updated_at": result.UpdatedAt,
			"created_at": result.CreatedAt,
			"User": gin.H{
				"id":       result.User.ID,
				"email":    result.User.Email,
				"username": result.User.Username,
			},
			"Photo": gin.H{
				"id":        result.Photo.ID,
				"title":     result.Photo.Title,
				"caption":   result.Photo.Caption,
				"photo_url": result.Photo.PhotoURL,
				"user_id":   result.Photo.UserID,
			},
		}

		comments = append(comments, comment)
	}

	context.JSON(http.StatusOK, comments)
}

func UpdateComment(context *gin.Context) {
	var message model.CommentUpdate

	if err := context.ShouldBindJSON(&message); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	ID, _ := pkg.GetIdParam(context, "commentId")

	commentID := uint(ID)

	Update, err := service.CommentService.UpdateComment(&message, commentID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         Update.ID,
		"message":    Update.Message,
		"user_id":    Update.UserID,
		"updated_at": Update.CreatedAt,
	})
}

func DeleteComment(context *gin.Context) {
	ID, _ := pkg.GetIdParam(context, "commentId")
	commentID := uint(ID)

	_, err := service.CommentService.DeleteComment(commentID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
