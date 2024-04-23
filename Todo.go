package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string    `json:"id"`
	Item      string    `json:"item"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"time"`
}

var todos = []todo{
	{ID: "1", Item: "Study to exam", Completed: false, CreatedAt: time.Now()},
	{ID: "2", Item: "Make bed", Completed: false, CreatedAt: time.Now()},
	{ID: "3", Item: "Read the Bible", Completed: false, CreatedAt: time.Now()},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.CreatedAt = time.Now()
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found."})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)

}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found."})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:3000")
}
