package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Gin58/gin-api/forms"
	"github.com/Gin58/gin-api/models"
)

var userModel = new(models.UserModel)

type UserController = struct {

}

func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H {
			"message": "Provide relevant fields",
		})
		c.Abort()
		return
	}

	err := userModel.Signup(data)

	if err != nil {
		c.JSON(400, gin.H {
			"message": "Problem creating an account",
		})
		c.Abort()
		return
	}

	c.JSON(201, gin.H {
		"message": "New user account registered",
	})
}
