package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(context *gin.Context, data gin.H) {
	context.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Error(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, gin.H{
		"message": message,
	})
}

func BadRequest(context *gin.Context, message string) {
	Error(context, http.StatusBadRequest, message)
}

func NotFound(context *gin.Context, message string) {
	Error(context, http.StatusNotFound, message)
}

func InternalServerError(context *gin.Context, message string) {
	Error(context, http.StatusInternalServerError, message)
}
