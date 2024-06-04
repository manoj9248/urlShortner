package main

import (
	"URLSHORTNER/endpoints"
	svc "URLSHORTNER/services"

	"github.com/gin-gonic/gin"
)

func main() {
	registerRoutes()
}
func registerRoutes() {
	port := "8020"
	router := gin.Default()
	urlSvc := svc.NewShortenURLServicer()
	endpoints.NewShortenHandler(router, urlSvc)
	router.Run(":" + port)
}
