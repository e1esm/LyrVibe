package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var noRightsErr = errors.New("no right to access resource")

func (ps *ProxyServer) RoleMiddleware(c *gin.Context) {
	switch {
	case strings.Contains(c.Request.URL.String(), "artist") && !strings.Contains(c.Request.URL.String(), "new"):
		if c.GetString("role") == "Artist" {
			c.Next()
		} else {
			c.AbortWithStatusJSON(403, fmt.Sprintf("%s:%s", noRightsErr, c.Request.URL))
		}
	default:
		c.Next()
	}
}
