package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tammysuryana93/helper"
	"github.com/tammysuryana93/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler (userService user.Service) *userHandler {
	return &userHandler{ userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user  ke struct register user input
	// struct di atas kita passing sebgagai parameter service


	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidasiErrot(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("gagal cuy", http.StatusUnprocessableEntity,"error euy", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("gagal maning , gagal maning ", http.StatusUnprocessableEntity,"gagal ning", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(newUser,"tokentokenkentotoketkentot")
	response := helper.APIResponse("Account has been registered", http.StatusOK,"success", formatter)
	c.JSON(http.StatusOK, response)
	return
}