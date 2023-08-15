package server

func setUpRoutes(server ProxyServer) {
	server.Router.HandleMethodNotAllowed = true
	authGroup := server.Router.Group("/v1/auth/", server.AuthMiddleware)
	authGroup.POST("signup", server.SignUp)
	authGroup.POST("login", server.Login)
	authGroup.POST("logout", server.Logout)
}
