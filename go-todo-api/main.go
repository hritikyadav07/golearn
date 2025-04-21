package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Todo struct
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

// In-memory todos slice
var todos = []Todo{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}

func getTodos(c *gin.Context) {
	query := c.Query("done") 
	if query == "" {
		c.IndentedJSON(http.StatusOK, todos)
	} else {
		res := []Todo{}
		done, err := strconv.ParseBool(query)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid 'done' parameter"})
			return
		}
		for _, todo := range todos {
			if todo.Done == done {
				res = append(res, todo)
			}
		}
		c.IndentedJSON(http.StatusOK, res)
	}


}

func addTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	} else {
		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)
		c.IndentedJSON(http.StatusCreated, newTodo)
	}
}
func updateTodo(c *gin.Context) {
	
	var updatedTodo Todo 
	if err := c.BindJSON((&updatedTodo)); err != nil {
		c.IndentedJSON((http.StatusBadRequest), gin.H{"error": "Shi Input de bhai"})
		return 
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "shi format me de re baba"})
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Done = updatedTodo.Done
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "nhi hai ye todo"})
}

func deleteTodo(c *gin.Context){
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "shi format me de re baba"})
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo uda diya re baba"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "nhi hai ye todo"})
}

func getTodo(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"shi format nhi hai"})
	}
	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"kidhr hai re baba"})
}



func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)

	router.GET("/todo/:id", getTodo)

	router.POST("/todo", addTodo)

	router.PUT("/todo/:id", updateTodo)

	router.DELETE("/todo/:id", deleteTodo)

	router.Run(":4000")
}
