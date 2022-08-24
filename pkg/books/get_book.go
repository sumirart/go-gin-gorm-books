package books

import (
	"github.com/gin-gonic/gin"
	"github.com/sumirart/go-gin-gorm-books/pkg/common/models"
	"net/http"
)

func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	// Jo said he rarely found something like this in Go source code
	// Usually (in Go source code), create and fedine result and error first, then do if error..
	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
