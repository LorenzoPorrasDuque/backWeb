package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	fmt.Println("logger")
	c.Next()
}
