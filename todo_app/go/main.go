package main

import (
	"https://github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Run(':8080')
}
