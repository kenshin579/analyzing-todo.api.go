package main

import (
	"encoding/json"
	"github.com/rctyler/todo.api.go/src/data"
	"github.com/rctyler/todo.api.go/src/dtos"
	"github.com/rctyler/todo.api.go/src/services"
	"github.com/rctyler/todo.api.go/src/shared/errors"
	"github.com/rctyler/todo.api.go/src/shared/types"
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
		v1.POST("/todo", func(c *gin.Context) {
			var data dtos.Todo

			c.BindJSON(&data)

			todo := types.Todo{
				Author: data.Author,
				TODO:   data.TODO,
				When:   data.When,
			}

			todo, err := todoService.Add(todo)

			if err == nil {
				c.JSON(http.StatusCreated, todo)
			} else if err.Code() == customErrors.ServiceUnavailableErrorCode {
				c.JSON(http.StatusServiceUnavailable, err)
			} else if err.Code() == customErrors.InternalServerErrorCode {
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.Status(http.StatusInternalServerError)
			}
		})

		v1.GET("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")

			todo, err := todoService.Get(id)

			if err == nil {
				c.JSON(http.StatusOK, todo)
			} else if err.Code() == customErrors.NotFoundErrorCode {
				c.JSON(http.StatusNotFound, err)
			} else if err.Code() == customErrors.ServiceUnavailableErrorCode {
				c.JSON(http.StatusServiceUnavailable, err)
			} else if err.Code() == customErrors.InternalServerErrorCode {
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.Status(http.StatusInternalServerError)
			}
		})

		v1.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")

			err := todoService.Delete(id)

			if err == nil {
				c.Status(http.StatusNoContent)
			} else if err.Code() == customErrors.ServiceUnavailableErrorCode {
				c.JSON(http.StatusServiceUnavailable, err)
			} else {
				c.Status(http.StatusInternalServerError)
			}
		})
	}

	router.Run(":8080")
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
