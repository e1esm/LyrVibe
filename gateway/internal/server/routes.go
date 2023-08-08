package server

func setUpRoutes(server ProxyServer) {
	server.Router.HandleMethodNotAllowed = true
	server.Router.POST("/v1/auth/signup", server.SignUp)
	server.Router.POST("/v1/auth/login", server.Login)
	server.Router.POST("v1/auth/logout", server.Logout)
}
