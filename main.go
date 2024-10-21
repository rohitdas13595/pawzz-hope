package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rohitdas13595/pawzz-hope/db"
	"github.com/rohitdas13595/pawzz-hope/docs"
	"github.com/rohitdas13595/pawzz-hope/utils"
	"github.com/rohitdas13595/pawzz-hope/zlog"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	//logger
	logger := zlog.InitLogger()
	defer logger.Sync()

	utils.InitSettings()
	// migration
	db.AutoMigrate()

	server := gin.Default()
	server.Use(zlog.GinLogger())

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})
	server.Use(corsConfig)

	v1 := server.Group("/api/v1")

	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})

	}

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Http Server Running on 8080",
		})
	})

	server.GET("/swagger/*any", gin.WrapH(docs.SwaggerHandler()))
	server.Run(":8080")

}
