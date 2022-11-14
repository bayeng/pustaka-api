package handler

import (
	"fmt"
	"go-web-api/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookHandler struct {
	bookService book.Service
}

func NewService(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}
func converterToResponse(abook book.Book) book.BookResponse {
	bookResponse := book.BookResponse{
		Id:     abook.Id,
		Title:  abook.Title,
		Desc:   abook.Desc,
		Price:  abook.Price,
		Rating: abook.Rating,
	}

	return bookResponse
}

func (h *bookHandler) GetAllBook(ctx *gin.Context) {
	books, err := h.bookService.FindAllBook()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var booksResponse []book.BookResponse
	for _, value := range books {
		bookResponse := converterToResponse(value)

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	aBook, err := h.bookService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "Not Found",
		})

		return
	}

	bookResponnse := converterToResponse(aBook)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponnse,
	})
}

func (h *bookHandler) PostBook(ctx *gin.Context) {
	var bookInput book.BookInput

	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {

		messages := []string{}
		for _, value := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("error on field %s , tag condition %s", value.Field(), value.ActualTag())
			messages = append(messages, message)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": messages,
		})

	}

	book, err := h.bookService.CreateBook(bookInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBook(ctx *gin.Context) {
	var bookRequest book.BookInput

	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {

		messages := []string{}
		for _, value := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("error on field %s , tag condition %s", value.Field(), value.ActualTag())
			messages = append(messages, message)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": messages,
		})
		return
	}

	updateBook, err := h.bookService.UpdateBook(id, bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data_update": updateBook,
	})
}

func (h *bookHandler) DeleteBook(ctx *gin.Context) {

	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	aBook, err := h.bookService.DeleteBook(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	delBook := converterToResponse(aBook)

	ctx.JSON(http.StatusOK, gin.H{
		"data": delBook,
	})
}
