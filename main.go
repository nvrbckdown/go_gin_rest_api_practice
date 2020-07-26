package main

import (
	"gin-gonic/v1/controllers"
	"gin-gonic/v1/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Инициализируем Gin как поумолчанию

	models.ConnectDataBase() // Подключаемся к базе данных

	// Router'ы
	router.GET("/books", controllers.FindBooks)   // router для получения всех книг
	router.POST("/books", controllers.CreateBook) // router для создание книги
	router.GET("/books/:id", controllers.FindBook) // router для получение одной книги
	router.PATCH("/books/:id", controllers.UpdateBook) // router для обновления книги
	router.DELETE("/books/:id", controllers.DeleteBook) // router для удаления книги

	// Стартуем Сервер на порту 8080
	router.Run()
}
