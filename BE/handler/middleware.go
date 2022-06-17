package handler

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckUser(c *gin.Context) {

	tknStr := c.Request.Header["Authorization"]
	if tknStr == nil {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}
	fixTkn := strings.Split(tknStr[0]," ")
	jwtTkn := fixTkn[1]
	claims := &Claims{}
	var jwtKey = []byte("key")
	tkn, err := jwt.ParseWithClaims(jwtTkn, claims, func(token *jwt.Token) (interface{}, error) {
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
	
	tknStr := c.Request.Header["Authorization"]
	if tknStr == nil {
		c.AbortWithStatusJSON(http.StatusForbidden,"Forbidden")
		return
	}
	fixTkn := strings.Split(tknStr[0]," ")
	jwtTkn := fixTkn[1]
	claims := &Claims{}
	var jwtKey = []byte("key")
	tkn, err := jwt.ParseWithClaims(jwtTkn, claims, func(token *jwt.Token) (interface{}, error) {
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

