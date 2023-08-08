package server

import (
	"errors"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/e1esm/LyrVibe/gateway/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var (
	successLogOut = "Successfully logged out of the service"
	reqBodyErr    = errors.New("body of the request doesn't fullfil server requirements")
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
	setUpRoutes(proxy)
	proxy.Services = services
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
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ps *ProxyServer) SignUp(c *gin.Context) {
	var signUpRequest proto.SignUpRequest
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, reqBodyErr)
		return
	}
	resp, err := ps.Services.AuthService.SignUp(&signUpRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ps *ProxyServer) Logout(c *gin.Context) {
	var logoutRequest proto.LogoutRequest
	if err := c.BindJSON(&logoutRequest); err != nil {
		c.JSON(http.StatusBadRequest, reqBodyErr)
	}
	err := ps.Services.AuthService.Logout(&logoutRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, successLogOut)

}
