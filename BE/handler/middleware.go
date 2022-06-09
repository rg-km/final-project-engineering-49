package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckUser(c *gin.Context) {
	ok := h.repo.IsLogin(c.Request.Header["Token"])
	if !ok {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	} else{
		c.Next()
	}
}  

func (h *Handler) CheckAdmin(c *gin.Context) {
	ok := h.repo.IsAdmin(c.Request.Header["Token"])
	if !ok {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	} else{
		c.Next()
	}
}  