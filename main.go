package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitdas13595/pawzz-hope/db"
	"github.com/rohitdas13595/pawzz-hope/utils"
	"github.com/rohitdas13595/pawzz-hope/zlog"
)

// var logger *zap.Logger

// func initLogger() {
// 	var err error
// 	logger, err = zap.NewProduction()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func GinLogger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()
// 		path := c.Request.URL.Path
// 		query := c.Request.URL.RawQuery

// 		c.Next()

// 		end := time.Now()
// 		latency := end.Sub(start)

// 		logger.Info("Request",
// 			zap.Int("status", c.Writer.Status()),
// 			zap.String("method", c.Request.Method),
// 			zap.String("path", path),
// 			zap.String("query", query),
// 			zap.String("ip", c.ClientIP()),
// 			zap.String("user-agent", c.Request.UserAgent()),
// 			zap.Duration("latency", latency),
// 		)
// 	}
// }

func main() {
	//logger
	logger := zlog.InitLogger()
	defer logger.Sync()

	utils.InitSettings()
	// migration
	db.AutoMigrate()

	server := gin.Default()
	server.Use(zlog.GinLogger())
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Http Server Running on 8080",
		})
	})
	server.Run(":8080")

}
