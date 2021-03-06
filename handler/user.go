package handler

import (
	"fmt"
	"net/http"
	"startup-funding/auth"
	"startup-funding/helper"
	"startup-funding/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register Account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// format response
	formatter := user.FormatUser(newUser, token)
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
	token, err := h.authService.GenerateToken(loggedUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// format response
	formatter := user.FormatUser(loggedUser, token)
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
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// input email di-mapping ke struct input
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email address has been registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"is_available": isEmailAvailable,
	}
	var metaMessage string
	metaMessage = "Email has been registerd"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	// api response
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	// return json
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar email", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// JWT ( Sementara Hardcode ,User Data )
	currentUser := c.MustGet("CurrentUser").(user.User)
	userID := currentUser.ID
	// generate file name
	path := fmt.Sprintf("public/images/%d-%s", userID, file.Filename)
	// simpan gambar ke local storage
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar email", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// parse service
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar email", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded!", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return

}
