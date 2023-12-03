package app

import (
	"os"
	"project-2/auth"
	"project-2/dto"

	"github.com/gin-gonic/gin"
)

// var PORT = "8080"

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", dto.Register)
		userRouter.POST("/login", dto.Login)
		userRouter.PUT("/", auth.Authentication(), dto.UpdateUser)
		userRouter.DELETE("/", auth.Authentication(), dto.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(auth.Authentication())
		photoRouter.POST("/", dto.CreatePhoto)
		photoRouter.GET("/", dto.GetAllPhotos)
		photoRouter.PUT("/:photoId", auth.PhotoAuthorization(), dto.UpdatePhoto)
		photoRouter.DELETE("/:photoId", auth.PhotoAuthorization(), dto.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(auth.Authentication())
		commentRouter.POST("/", dto.CreateComment)
		commentRouter.GET("/", dto.GetComment)
		commentRouter.PUT("/:commentId", auth.CommentAuthorization(), dto.UpdateComment)
		commentRouter.DELETE("/:commentId", auth.CommentAuthorization(), dto.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(auth.Authentication())
		socialMediaRouter.POST("/", dto.CreateSocialMedia)
		socialMediaRouter.GET("/", dto.GetAllSocialMedias)
		socialMediaRouter.PUT("/:socialMediaId", auth.SocialMediaAuthorization(), dto.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", auth.SocialMediaAuthorization(), dto.DeleteSocialMedia)
	}

	router.Run(":" + os.Getenv("PORT"))
}
