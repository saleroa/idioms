package utils

import (
	"github.com/gin-gonic/gin"
)

func RespDiy(c *gin.Context, status int, s string) {
	c.JSON(status, gin.H{
		"status": status,
		"info":   s})
}
