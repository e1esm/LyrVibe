package server

import "github.com/e1esm/LyrVibe/auth-service/internal/service"

type Server struct {
	AuthService service.Service
}
