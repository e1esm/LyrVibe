package server

import (
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	refreshError  = errors.New("refresh error")
	unauthorized  = errors.New("unauthorized")
	verfification = errors.New("verification error")
)

const (
	accessTokenName  = "access_token"
	refreshTokenName = "refresh_token"
)

func (ps *ProxyServer) Login(c *gin.Context) {
	var signInRequest proto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		c.JSON(http.StatusBadRequest, reqBodyErr)
		return
	}
	resp, err := ps.Services.AuthService.Login(&signInRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	accessTTL, err := time.ParseDuration(resp.Tokens.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TTLErr)
		return
	}
	refreshTTL, err := time.ParseDuration(resp.Tokens.RefreshTTL)
	if err != nil {
		logger.GetLogger().Info(err.Error())
		c.JSON(http.StatusInternalServerError, TTLErr)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     accessTokenName,
		Value:    resp.Tokens.AccessToken,
		Expires:  time.Now().Add(accessTTL),
		HttpOnly: true,
		Path:     "/",
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     refreshTokenName,
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
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, reqBodyErr)
		return
	}

	resp, err := ps.Services.AuthService.SignUp(&signUpRequest)
	if err != nil {
		logger.GetLogger().Info(fmt.Sprintf("Password: %v", signUpRequest.Password))
		c.JSON(http.StatusInternalServerError, gin.H{
			"username": signUpRequest.Username,
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ps *ProxyServer) Logout(c *gin.Context) {
	token, err := c.Cookie(refreshTokenName)
	if err != nil {
		c.JSON(http.StatusBadRequest, cookieRetrievingErr)
	}
	err = ps.Services.AuthService.Logout(&proto.LogoutRequest{RefreshToken: token})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     accessTokenName,
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     refreshTokenName,
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})

	c.JSON(http.StatusOK, successLogOut)

}
