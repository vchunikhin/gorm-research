package controllers

import (
	"github.com/gin-gonic/gin"
)

func failResponse(message string) *gin.H {
	return &gin.H{"status": "fail", "message": message}
}

func errorResponse(message string) *gin.H {
	return &gin.H{"status": "error", "message": message}
}

func successResponse(data any) *gin.H {
	return &gin.H{"status": "success", "data": data}
}

func paginationResponse(size int, data any) *gin.H {
	return &gin.H{"status": "success", "size": size, "data": data}
}
