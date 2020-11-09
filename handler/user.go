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

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get input from form data
	var input user.RegisterUserInput
	// map input from user to struct Register User Input and validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account failer", http.StatusUnprocessableEntity, "error", errorMessage)
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
	// transform response
	formatter := user.FormatUser(newUser, "asdasduhquwehasdkfjokjASDJKOJAORIJA;KLFSDJGOIERFIUHAKJSDNFKJNWEOIDRJOADNSLKJASL;KD")
	// api response
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	// return json
	c.JSON(http.StatusOK, response)
}
