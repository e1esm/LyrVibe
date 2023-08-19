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

func (ps *ProxyServer) Login(c *gin.Context) {
	var signInRequest proto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		c.JSON(http.StatusBadRequest, reqBodyErr)
		return
	}
	resp, err := ps.Services.AuthService.Login(&signInRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	accessTTL, err := time.ParseDuration(resp.Tokens.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TTLErr)
		return
	}
	refreshTTL, err := time.ParseDuration(resp.Tokens.RefreshTTL)
	if err != nil {
		logger.Logger.Info(err.Error())
		c.JSON(http.StatusInternalServerError, TTLErr)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    resp.Tokens.AccessToken,
		Expires:  time.Now().Add(accessTTL),
		HttpOnly: true,
		Path:     "/",
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    resp.Tokens.RefreshToken,
		Expires:  time.Now().Add(refreshTTL),
		HttpOnly: true,
		Path:     "/",
	})
	c.JSON(http.StatusOK, LogInOk)
}

func (ps *ProxyServer) SignUp(c *gin.Context) {
	signUpRequest := proto.SignUpRequest{}
	if err := c.BindJSON(&signUpRequest); err != nil {
		logger.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, reqBodyErr)
		return
	}

	resp, err := ps.Services.AuthService.SignUp(&signUpRequest)
	if err != nil {
		logger.Logger.Info(fmt.Sprintf("Password: %v", signUpRequest.Password))
		c.JSON(http.StatusInternalServerError, gin.H{
			"username": signUpRequest.Username,
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ps *ProxyServer) Logout(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, cookieRetrievingErr)
	}
	err = ps.Services.AuthService.Logout(&proto.LogoutRequest{AccessToken: token})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, successLogOut)

}

func (ps *ProxyServer) AuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil || accessToken == "" {
		refreshToken, _ := c.Cookie("refresh_token")
		if refreshToken == "" {
			logger.Logger.Error("No required tokens", zap.String("err", err.Error()))
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		resp, err := ps.Services.AuthService.Refresh(&proto.RefreshRequest{RefreshToken: refreshToken})
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, "Refresh error")
			return
		}
		ttl, err := time.ParseDuration(resp.Ttl)
		if err != nil {
			logger.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, "Refresh error")
			return
		}
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "access_token",
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
		logger.Logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	c.Set("username", resp.Username)
	c.Set("role", resp.Role)
	c.Set("id", resp.Id)
	c.Next()
}
