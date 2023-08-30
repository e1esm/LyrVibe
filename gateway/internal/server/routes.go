package server

func setUpRoutes(server ProxyServer) {
	server.Router.HandleMethodNotAllowed = true
	setupAuthRoutes(&server)
	setupArtistRoutes(&server)
}

func setupAuthRoutes(server *ProxyServer) {
	authGroup := server.Router.Group("/v1/auth/")
	authGroup.POST("signup", server.SignUp)
	authGroup.POST("login", server.Login)
	authGroup.POST("logout", server.AuthMiddleware, server.Logout)
}

func setupArtistRoutes(server *ProxyServer) {
	artistGroup := server.Router.Group("/v1/artist/")
	artistGroup.Use(server.AuthMiddleware)
	artistGroup.POST("new", server.NewArtist)
	artistGroup.POST("release", server.ReleaseTrack)
}
