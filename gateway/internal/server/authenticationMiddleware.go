package server

import (
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func (ps *ProxyServer) AuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie(accessTokenName)
	if err != nil || accessToken == "" {
		refreshToken, _ := c.Cookie(refreshTokenName)
		if refreshToken == "" {
			logger.GetLogger().Error("No required tokens", zap.String("err", err.Error()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized.Error())
			return
		}
		resp, err := ps.Services.AuthService.Refresh(&proto.RefreshRequest{RefreshToken: refreshToken})
		if err != nil {
			logger.GetLogger().Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, refreshError.Error())
			return
		}
		ttl, err := time.ParseDuration(resp.Ttl)
		if err != nil {
			logger.GetLogger().Error(err.Error(), zap.String("ttl", fmt.Sprintf("%v", ttl)))
			c.AbortWithStatusJSON(http.StatusInternalServerError, refreshError.Error())
			return
		}
		logger.GetLogger().Info(resp.AccessToken)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     accessTokenName,
			Value:    resp.AccessToken,
			Expires:  time.Now().Add(ttl),
			HttpOnly: true,
			Path:     "/",
		})
		accessToken = resp.AccessToken
	}
	resp, err := ps.Services.AuthService.Verify(&proto.VerificationRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, verfification.Error())
		return
	}
	c.Set("username", resp.Username)
	c.Set("role", resp.Role)
	c.Set("id", resp.Id)
	c.Next()
}
