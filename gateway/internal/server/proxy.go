package server

import (
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/e1esm/LyrVibe/gateway/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Proxy interface {
	Run(address string)
	Login(c *gin.Context)
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
	c.JSON(200, "OK")
}
