package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckUser(c *gin.Context) {

	tknStr := c.Request.Header["Token"]
	if tknStr == nil {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}
	claims := &Claims{}
	var jwtKey = []byte("key")
	tkn, err := jwt.ParseWithClaims(tknStr[0], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}

	if !tkn.Valid {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}
	c.Next()
}  

func (h *Handler) CheckAdmin(c *gin.Context) {
	
	tknStr := c.Request.Header["Token"]
	if tknStr == nil {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}
	claims := &Claims{}
	var jwtKey = []byte("key")
	tkn, err := jwt.ParseWithClaims(tknStr[0], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}

	if !tkn.Valid {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}

	if claims.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}

	c.Next()
}   

