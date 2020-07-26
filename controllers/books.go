package controllers

import (
	"gin-gonic/v1/models"
	"github.com/gin-gonic/gin"
	)

// GET /books
// GET all books
func FindBooks(c *gin.Context) {
	var books []models.Book // Переменная типа Book из модела
	models.DB.Find(&books) // Берем с базы все книги
	c.JSON(200, gin.H{"data": books}) // Отправляем как json
}

// Struct to create book
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	var input CreateBookInput // Оределяем входящие данные, Титул и Автора
	if err := c.ShouldBindJSON(&input); err != nil { // Проверяем JSON на целостьность, если нет возврашяем ошибку
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Берем данные со входа и парсим через модел Book и создаем книгу в Базе.
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(200, gin.H{"data": book}) // Возврашаем результат с id
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	var book models.Book //Определяем книгу
	//Ищем книгу в базе если не находим возврашаем ошибку
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Record not found!",
		})
		return
	}
	c.JSON(200, gin.H{"data": book}) //Возврашаем книгу
}

// Struct to update book
type UpdateBootInput struct {
	Title  string `json: "title"`
	Author string `json: "author"`
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context)  {
	var book models.Book //Определяем книгу
	// Ищем в базе
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Record not found!",
		})
		return
	}
	var input UpdateBootInput // Проверяем входяшие данные
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Model(&book).Update(input)
	c.JSON(200, gin.H{
		"data": book,
	})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context)  {
	var book models.Book // Определяем книгу
	// Ищем в базе есть ли книга если нет возврашаем ошибку
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Record not found!",
		})
		return
	}
	models.DB.Delete(&book) // Удаляем книгу с базы
	// Возврашаем ответ
	c.JSON(200, gin.H{
		"data": true,
	})
}


