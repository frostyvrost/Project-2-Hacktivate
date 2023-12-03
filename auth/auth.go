package auth

import (
	"net/http"
	"project-2/database"
	"project-2/model"
	"project-2/pkg"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifiedToken, err := pkg.VerifyToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.Set("userData", verifiedToken)
		context.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		photoId, err := pkg.GetIdParam(context, "photoId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		photo := model.Photo{}

		errMsg := db.Select("user_id").First(&photo, photoId).Error
		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserID != userID {
			err := pkg.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		socialMediaId, err := pkg.GetIdParam(context, "socialMediaId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		socialMedia := model.SocialMedia{}

		errMsg := db.Select("user_id").First(&socialMedia, socialMediaId).Error

		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socialMedia.UserID != userID {
			err := pkg.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		commentId, err := pkg.GetIdParam(context, "commentId")
		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		comment := model.Comment{}

		errMsg := db.Select("user_id").First(&comment, uint(commentId)).Error
		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if comment.UserID != userID {
			err := pkg.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
