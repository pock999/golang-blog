package routers

import (
	"github.com/gin-gonic/gin"

	"blog/api/controllers/AuthController"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1");
	{
		v1.GET("/users", AuthController.UserList);
	}

	// router.POST("/user", Store)

	// router.PUT("/user/:id", Update)

	// router.DELETE("/user/:id", Destroy)

	return router
}