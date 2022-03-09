package main

import (
	"context"
	"net/http"
	"time"

	prisma "github.com/Duzbee/TodoApp/server/generated/prisma-client"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var client = prisma.New(nil)
var ctx = context.TODO()

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/todos", getTodos)
	router.POST("/create-todo", createTodo)
	router.PATCH("/edit-todo/:id", editTodo)
	router.PATCH("/done-todo/:id", doneTodo)
	router.DELETE("/delete-todo/:id", deleteTodo)

	router.Run("localhost:3000")
}

func getTodos(c *gin.Context) {
	todos, err := client.Todoes(nil).Exec(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error while getting todos."})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "todos": todos})
}

func createTodo(c *gin.Context) {
	var body Todo

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid body JSON structure."})
		return
	}

	todo, err := client.CreateTodo(prisma.TodoCreateInput{
		Text: body.Text,
		Done: &body.Done,
	}).Exec(ctx)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error while creating todo."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "todo": todo})
}

func editTodo(c *gin.Context) {
	var body Todo
	id := c.Param("id")

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid body JSON structure."})
		return
	}

	updatedTodo, err := client.UpdateTodo(prisma.TodoUpdateParams{
		Where: prisma.TodoWhereUniqueInput{
			ID: &id,
		},
		Data: prisma.TodoUpdateInput{
			Text: &body.Text,
		},
	}).Exec(ctx)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error while updating todo."})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "todo": updatedTodo})
}

func doneTodo(c *gin.Context) {
	var body Todo
	id := c.Param("id")

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Invalid body JSON structure."})
		return
	}

	updatedTodo, err := client.UpdateTodo(prisma.TodoUpdateParams{
		Where: prisma.TodoWhereUniqueInput{
			ID: &id,
		},
		Data: prisma.TodoUpdateInput{
			Done: &body.Done,
		},
	}).Exec(ctx)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error while updating todo."})

	}

	c.IndentedJSON(http.StatusOK, updatedTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	_, err := client.DeleteTodo(prisma.TodoWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found."})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "id": id})
}
