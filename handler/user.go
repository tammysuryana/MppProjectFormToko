package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tammysuryana93/auth"
	"github.com/tammysuryana93/helper"
	"github.com/tammysuryana93/user"
	"net/http"
)
type userHandler struct {
	userService user.Service
	authServices auth.Servis
}
func NewUserHandler(userService user.Service, authServis auth.Servis) *userHandler {
	return &userHandler{  userService, authServis }
}
//func NewUserHandler(userService UserService) *userHandler {
//	return &userHandler{ userService}
//}
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
	token , err := h.authServices.GenerateToken(newUser.ID)

	formatter := user.FormatUser(newUser,token)
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

		response := helper.APIResponse("fail login ", 422,"status err", errorMessage)
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
	token, err := h.authServices.GenerateToken(loggedinUser.ID)

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Success ", 200,"Yeah", formatter)
	c.JSON(200, response)
	return
}
func (h *userHandler)CheckEmaiAvailability(c *gin.Context){
	// ada inpuut email dari user
	// input email di mapping ke struck inut
	// stuct input di passsing ke service
	// service akan memanggil repository - email sudah ada atay belum
	// repository db
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidasiErrot(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("Check Email  ", http.StatusUnprocessableEntity,"status err", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Eerrorr "}
		response := helper.APIResponse("Check Email  ", http.StatusUnprocessableEntity,"status err", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
			data := gin.H{
				"is_Available": IsEmailAvailable,
			}
		metaMessage := "Email has bin Registered "
			if IsEmailAvailable {
				metaMessage = "Email is Available"
			}
		response := helper.APIResponse(metaMessage, 200, "success",data)
		c.JSON(200 , response)
}
func(h *userHandler) UploadAvatar (c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is uploadeds": false}
		respose := helper.APIResponse("gagal mengupload avatar ", 402, "error", data)
		c.JSON(402,respose)
		return
	}
	userID := 36
	//path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is uploadeds": false}
		respose := helper.APIResponse("gagal mengupload avatar ", 402, "error", data)
		c.JSON(402, respose)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	 if err != nil {
		 data := gin.H{"is uploadeds": false}
		 respose := helper.APIResponse("gagal mengupload avatar ", 402, "error", data)
		 c.JSON(402, respose)
		 return
	 }
	data := gin.H{"uploaded Success": true}
	respose := helper.APIResponse(" avatar TERUPLOAD ", 200, "Suxcess", data)
	c.JSON(402, respose)
	// input dari User
	// simpan gambar nya di folder "images/"
	// di service kita panggil repository
	// JWT (sementara di buat hardcore , eakan akan user yang login ID = 1)
	// repository ambbil data yang ID = 1
	// repo Update data usern simpan di lokasi file
}
