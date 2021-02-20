package AuthController

import (
	"github.com/gin-gonic/gin"
	"net/http"
  // "strconv"

	"blog/api/models"
)

func UserList(context *gin.Context) {
	var User models.User;
	result, err := User.FindAll();

	if err != nil {
		context.JSON(400, gin.H{
			"statusCode": 400,
			"message": "not found",
		});
		return ;
	}

	context.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"data": result,
	});
}