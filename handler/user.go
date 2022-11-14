package handler

import (
	"fmt"
	"go-web-api/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type userHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUsers(ctx *gin.Context) {

	users, err := h.userService.FindUsers()
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var usersResponse []user.UserResponse
	for _, value := range users {
		userResponse := converterUser(value)
		usersResponse = append(usersResponse, userResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": usersResponse,
	})

}
func (h *userHandler) GetUser(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	user, err := h.userService.FindUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	convert := converterUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"data": convert,
	})
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	var inputUser user.InputUser

	err := ctx.ShouldBindJSON(&inputUser)
	if err != nil {

		var messages = []string{}
		for _, value := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on Field %s in tag %s", value.Field(), value.ActualTag())
			messages = append(messages, message)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": messages,
		})

		return
	}
	user, err := h.userService.CreateUser(inputUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) Deleteuser(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	deleteUser, err := h.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	convert := converterUser(deleteUser)

	ctx.JSON(http.StatusOK, gin.H{
		"data": convert,
	})
}

func converterUser(users user.User) user.UserResponse {
	var response = user.UserResponse{

		Id:       users.Id,
		Name:     users.Name,
		Username: users.Username,
		Password: users.Password,
		CreateAt: users.CreateAt,
	}

	return response

}
