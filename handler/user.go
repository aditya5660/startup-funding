package handler

import (
	"net/http"
	"startup-funding/helper"
	"startup-funding/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *gin.Context) {
	// get input from form data
	var input user.RegisterUserInput
	// map input from user to struct Register User Input and validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// struct diatas akan di parse setelah service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register Account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// get token
	// token,err := h.jwtService.GenerateToken()
	// format response
	formatter := user.FormatUser(newUser, "asdasduhquwehasdkfjokjASDJKOJAORIJA;KLFSDJGOIERFIUHAKJSDNFKJNWEOIDRJOADNSLKJASL;KD")
	// api response
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	// return json
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukkan input form (email & password)
	// mapping dari input user ke input struct
	var input user.LoginUserInput
	// input menangkap handler
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// input struct passing to service
	// di service mencari bantuan repository dengan input email
	loggedUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// format response
	formatter := user.FormatUser(loggedUser, "asdasduhquwehasdkfjokjASDJKOJAORIJA;KLFSDJGOIERFIUHAKJSDNFKJNWEOIDRJOADNSLKJASL;KD")
	// api response
	response := helper.APIResponse("Login successfuly!", http.StatusOK, "success", formatter)
	// return json
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Fetch(c *gin.Context) {
	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	// format response
	formatter := user.FormatUser(currentUser, "asdjhajksdhkljahsdlkjhalkjsdhajkhsdkjhasdkljhalkds")
	// api response
	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)
	// return json
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input email dari user
	// input email di-mapping ke struct input
	// struct input di-passing ke service
	// service akan memanggil repository - email sudah ada apa belum
	// repository akan search query ke db

	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email address has been registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data:= gin.H{
		"is_available": isEmailAvailable,
	}
	var metaMessage string
	metaMessage = "Email has been registerd"
	if isEmailAvailable{
		metaMessage = "Email is available"
	}
	// api response
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	// return json
	c.JSON(http.StatusOK, response)
}
