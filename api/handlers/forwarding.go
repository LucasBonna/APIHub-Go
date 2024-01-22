package handlers

import (
	"github.com/gin-gonic/gin"
)

func FWR(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ForwardRequest(apiURL)(c.Writer, c.Request)
	}
}
