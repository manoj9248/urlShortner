package endpoints

import (
	"errors"
	"fmt"
	"net/http"

	svc "URLSHORTNER/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service svc.ShortenURL
}

func NewShortenHandler(router *gin.Engine, svc svc.ShortenURL) {
	urlHandler := &Handler{
		Service: svc,
	}
	router.POST("/v1/", urlHandler.Shortenurl)
	router.GET("/v1/url/:shortenurl", urlHandler.Redirect)
	router.GET("/v1/getmetrics", urlHandler.Metrics)
}

type CreateShortUrl struct {
	Url string `json:"url"`
}

func (h *Handler) Metrics(c *gin.Context) {
	domainResp := h.Service.DomainCount()
	if len(domainResp) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("non of domain are registered "),
		})
		return
	}
	c.JSON(http.StatusOK, &domainResp)

}
func (h *Handler) Shortenurl(c *gin.Context) {
	url := CreateShortUrl{}
	if err := c.BindJSON(&url); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if url.Url == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("URL should be valid"))
	}
	resp, _ := h.Service.GetShorternURL(url.Url)
	c.JSON(http.StatusOK, gin.H{
		"shortUrl": fmt.Sprintf("http://localhost:8020/v1/%s", resp),
	})

}
func (h *Handler) Redirect(c *gin.Context) {
	shortUrl := c.Param("shortenurl")
	resp, isExists := h.Service.GetURL(shortUrl)
	if !isExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusFound, resp)
}
