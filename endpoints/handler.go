package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service ShortenURL
}

func NewHandler(Service ShortenURLService) Handler {
	return Handler{Service: Service}
}
func (h *Handler) Metrics(c *gin.Context) {

}
func (h *Handler) Shortenurl(c *gin.Context) {

}
func (h *Handler) Redirect(c *gin.Context) {
	shortUrl := c.Param("shortenurl")
	resp := h.Service.GetURL(shortUrl)
	c.JSON(http.StatusOK, &resp)
}
