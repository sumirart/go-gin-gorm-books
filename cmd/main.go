package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"

	"github.com/sumirart/go-gin-gorm-books/pkg/books"
	"github.com/sumirart/go-gin-gorm-books/pkg/common/db"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbURL := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbURL)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"port":  port,
			"dbURL": dbURL,
		})
	})

	router.Run(port)
}
