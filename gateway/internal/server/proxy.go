package server

import (
	"errors"
	"github.com/e1esm/LyrVibe/gateway/internal/service"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
