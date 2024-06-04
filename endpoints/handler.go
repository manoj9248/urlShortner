package endpoints

import (
	"errors"
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
	router.GET("/v1/:shortenurl", urlHandler.Redirect)
	router.GET("/v1/getmetrics", urlHandler.Metrics)
}

type CreateShortUrl struct {
	Url string `json:"url"`
}

func (h *Handler) Metrics(c *gin.Context) {
	domainResp := h.Service.DomainCount()
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
	c.JSON(http.StatusOK, &resp)

}
func (h *Handler) Redirect(c *gin.Context) {
	shortUrl := c.Param("shortenurl")
	resp := h.Service.GetURL(shortUrl)
	c.JSON(http.StatusOK, &resp)
}
