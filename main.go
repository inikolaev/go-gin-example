package main

import (
	"github.com/gin-gonic/gin"
	"github.com/modern-go/reflect2"
	"net/http"
)

type Amount struct {
	Value    int32  `json:"value" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type PaymentMethod struct {
	Type string `json:"type" binding:"required"`
	Id   string `json:"id" binding:"required"`
}

type Payment struct {
	Amount        Amount        `json:"amount" binding:"required"`
	PaymentMethod PaymentMethod `json:"payment_method" binding:"required"`
}

func problem(status int, _type string, title string, detail string) gin.H {
	return gin.H{
		"type":   _type,
		"status": status,
		"title":  title,
		"detail": detail,
	}
}

func main() {
	println("Type by name:", reflect2.TypeByName("main.PaymentMethod").Type1().Name())
	println()

	r := gin.Default()
	r.Use(gin.Logger())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, problem(
			404,
			"https://localhost/resource-not-found",
			"Resource not found",
			"Endpoint does not exist: "+c.Request.RequestURI,
		))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/payments", func(c *gin.Context) {
		var json Payment
		if err := c.BindJSON(&json); err != nil {
			c.JSON(
				http.StatusBadRequest,
				problem(
					http.StatusBadRequest,
					"https://localhost/invalid-payment-request",
					"Invalid request",
					err.Error(),
				),
			)
			return
		}
		c.JSON(200, json)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
