package controllers

import (
	"auther/models"
	"auther/util"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Register user
func Register(context *gin.Context) {
	var input models.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}

// User Login
func Login(context *gin.Context) {
	var input models.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := models.GetUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := util.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Auth", jwt, 3600*24, "", "", false, true)
	context.JSON(http.StatusOK, gin.H{"username": input.Username, "message": "Successfully logged in"})

}

// get all users
func GetUsers(context *gin.Context) {
	var user []models.User
	err := models.GetUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// get user by id
func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user models.User
	err := models.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// update user
func UpdateUser(c *gin.Context) {
	//var input model.Update
	var User models.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := models.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&User)
	err = models.UpdateUser(&User)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}

// logout
func Logout(context *gin.Context) {
	context.SetCookie("Auth", "", -1, "/", "localhost", false, true)
	context.JSON(200, gin.H{"success": "user logged out"})
}
