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
	// token , err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser,"tokentokenkentotoketkentot")
	response := helper.APIResponse("Account has been registered", http.StatusOK,"success", formatter)
	c.JSON(http.StatusOK, response)

}
func (h *userHandler) Login (c *gin.Context){
	// user memaksujkan input (email dan password atau dengan Nomer KTA)
	// input di tanggkap handler
	// mapping dari input user
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email atau KTA x
	// memcocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidasiErrot(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("fail login ", http.StatusUnprocessableEntity,"status err", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("login failed", 422, "errort" , errorMessage)
		c.JSON(422, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, "totkoktoktoktoktok")
	response := helper.APIResponse("Success ", 200,"Yeah", formatter)
	c.JSON(200, response)
	return
}