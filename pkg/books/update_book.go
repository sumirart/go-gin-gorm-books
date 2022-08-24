package books

import (
	"github.com/gin-gonic/gin"
	"github.com/sumirart/go-gin-gorm-books/pkg/common/models"
	"net/http"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)
	ctx.JSON(http.StatusOK, &book)
}
