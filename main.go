package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер
	r := gin.Default()

	// Определяем обработчик для корневого маршрута
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
