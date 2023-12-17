package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func main() {
	r := gin.Default()


	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	
	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}