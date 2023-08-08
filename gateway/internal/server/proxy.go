package server

import (
	"github.com/e1esm/LyrVibe/gateway/internal/service"
	"github.com/gin-gonic/gin"
)

type Proxy interface {
}

type ProxyServer struct {
	Router   *gin.Engine
	Services service.Services
}

func NewProxyServer(authProvider service.AuthenticationProvider) *ProxyServer {
	proxy := ProxyServer{}
	proxy.Router = gin.Default()
	proxy.Services.AuthService = authProvider
	return &proxy
}
