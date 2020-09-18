package main

import (
	"fmt"

	"consume-bluebird/worker/consume/services"

	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

func main() {
	port := beego.AppConfig.DefaultString("httpport", "9090")

	gin.SetMode(gin.DebugMode)

	router := gin.New()
	fmt.Println("Waiting connection... [", port, "] ")

	go services.StartServicesKafka()
	router.Use(CORSMiddleware())
	router.Run(":" + port)

}

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

