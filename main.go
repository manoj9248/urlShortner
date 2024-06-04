package main

import (
	"URLSHORTNER/endpoints"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlHandler endpoints.Handler
}

func main() {
	registerRoutes()
}
func registerRoutes() {
	port := "8020"
	router := gin.Default()
	endpoints.NewShortenURLServicer()
	urlHandler := urlHandler{
		urlHandler: endpoints.NewHandler(endpoints.ShortenURLService{}),
	}
	router.POST("/v1/", urlHandler.urlHandler.Shortenurl)
	router.GET("/v1/url/:shortenurl", urlHandler.urlHandler.Redirect)
	router.GET("/v1/getmetrics", urlHandler.urlHandler.Metrics)
	router.Run(":" + port)
}
