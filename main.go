package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	run()
}
func run() {
	port := "8080"
	router := gin.Default()
	router.Run(port)
}
