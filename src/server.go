package main

import "gopkg.in/gin-gonic/gin.v1"
import "github.com/rctyler/todoapp/src/services"
import "github.com/rctyler/todoapp/src/data"
import "net/http"

func main() {
	router := gin.Default()

	// Inject dependencies
	redisCacheService := data.NewRedisCacheService()
	todoService := services.NewTodoService(redisCacheService)

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		v1.GET("/todo", func(c *gin.Context) {
			todos := todoService.GetAll()
			c.JSON(200, todos)
		})

		v1.POST("/todo", func(c *gin.Context) {
			todo := todoService.Add()
			c.JSON(200, todo)
		})

		v1.GET("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			todo := todoService.Get(id)
			c.JSON(200, todo)
		})

		v1.PUT("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			todo := todoService.Update(id)
			c.JSON(200, todo)
		})

		v1.PATCH("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			todo := todoService.Patch(id)
			c.JSON(200, todo)
		})

		v1.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			todo := todoService.Delete(id)
			c.JSON(200, todo)
		})
	}

	router.Run()
}
