package api

import (
	"test/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	//	h := handler.New(store)

	r := gin.New()

	return r
}
