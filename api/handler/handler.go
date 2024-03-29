package handler

import (
	"connected/api/models"
	"connected/service"
	"connected/storage"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
   storage  storage.IStorage
	services service.IServiseManger
}

func New(services service.IServiseManger, store storage.IStorage) Handler {
	return Handler{
		storage:  store,
		services: services,
	}
}

func handleResponse(c *gin.Context, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "success"
	case code < 500:
		resp.Description = "bad request"
		fmt.Println("BAD REQUEST: "+msg, "reason: ", data)
	default:
		resp.Description = "internal server error"
		fmt.Println("INTERNAL SERVER ERROR: "+msg, "reason: ", data)
	}
	fmt.Println("data: ", data)

	resp.StatusCode = statusCode
	resp.Data = data
	fmt.Println("resp ", resp)

	c.JSON(resp.StatusCode, resp)
}
