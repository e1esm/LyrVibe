package server

import (
	"errors"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/internal/service"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

var (
	LogInOk             = "Successfully logged in"
	TTLErr              = errors.New("error while setting TTL to the cookie")
	successLogOut       = "Successfully logged out of the service"
	reqBodyErr          = errors.New("body of the request doesn't fulfill server requirements")
	cookieRetrievingErr = errors.New("couldn't have retrieved cookie")
)

type Proxy interface {
	Run(address string)
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type ProxyServer struct {
	Router   *gin.Engine
	Services service.Services
}

func NewProxyServer(services service.Services) Proxy {
	proxy := ProxyServer{}
	proxy.Router = gin.Default()
	proxy.Services = services
	setUpRoutes(proxy)
	return &proxy
}

func (ps *ProxyServer) Run(address string) {
	if err := ps.Router.Run(address); err != nil {
		logger.Logger.Fatal("Couldn't have started the server",
			zap.String("err", err.Error()))
	}
}

func (ps *ProxyServer) Login(c *gin.Context) {
	var signInRequest proto.SignInRequest
	c.Header("Content-Type", "application/json")
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
	signUpRequest := &proto.SignUpRequest{}
	c.Header("Content-Type", "application/json")

	if err := c.BindJSON(signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, reqBodyErr)
	}

	resp, err := ps.Services.AuthService.SignUp(signUpRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Status)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ps *ProxyServer) Logout(c *gin.Context) {
	token, err := c.Cookie("access-token")
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
	if strings.Contains(c.Request.RequestURI, "logout") {
		if token, err := c.Cookie("access-token"); err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
	}
	c.Next()
}
