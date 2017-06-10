package main

import (
	"encoding/json"
	"github.com/rctyler/todoapp/src/data"
	"github.com/rctyler/todoapp/src/services"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http"
	"os"
)

type config struct {
	RedisAddr string
	RedisPwd  string
	RedisDb   int
}

func main() {
	router := gin.Default()

	config, err := getConfigSettings("../config.json")
	if err != nil {
		os.Exit(1)
	}

	// Inject dependencies
	redisCacheService := data.NewRedisCacheService(config.RedisAddr, config.RedisPwd, config.RedisDb)
	todoService := services.NewTodoService(redisCacheService)

	router.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	v1 := router.Group("/v1")
	{
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

func getConfigSettings(location string) (config, error) {
	var config config

	configFile, err := ioutil.ReadFile(location)
	if err != nil {
		return config, err
	}

	json.Unmarshal(configFile, &config)
	return config, nil
}
