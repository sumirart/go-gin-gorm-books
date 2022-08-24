package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumirart/go-gin-gorm-books/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h handler) AddBook(ctx *gin.Context) {
	body := AddBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
