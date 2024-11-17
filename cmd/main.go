package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var tasks = []Task{
	{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

// Handler to get all tasks
func getTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": tasks,
	})
}

// Handler to get a task by ID
func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(200, gin.H{
				"data": task,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Task not found",
	})
}

// Handler to create a new task
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	newTask.ID = time.Now().Format("20060102150405") // Unique ID
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()
	tasks = append(tasks, newTask)

	c.JSON(201, gin.H{
		"message": "Task created successfully",
		"data":    newTask,
	})
}

// Handler to update a task
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedData Task

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedData.Title
			tasks[i].Description = updatedData.Description
			tasks[i].Completed = updatedData.Completed
			tasks[i].UpdatedAt = time.Now()

			c.JSON(200, gin.H{
				"message": "Task updated successfully",
				"data":    tasks[i],
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Task not found",
	})
}

// Handler to delete a task
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(200, gin.H{
				"message": "Task deleted successfully",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Task not found",
	})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", getTasks)
		v1.GET("/tasks/:id", getTaskByID)
		v1.POST("/tasks", createTask)
		v1.PUT("/tasks/:id", updateTask)
		v1.DELETE("/tasks/:id", deleteTask)
	}

	router.Run(":8080")
}
