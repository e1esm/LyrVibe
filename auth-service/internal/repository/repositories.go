package repository

import (
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/sessionRepository/redis"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/userRepository/postgres"
)

type Repositories struct {
	SessionRepository redis.SessionStorage
	MainRepository    postgres.UserStorage
}

type RepositoriesBuilder struct {
	repositories Repositories
}

func NewRepositoriesBuilder() *RepositoriesBuilder {
	return &RepositoriesBuilder{}
}

func (rb *RepositoriesBuilder) WithMainRepo(mainStorage postgres.UserStorage) *RepositoriesBuilder {
	rb.repositories.MainRepository = mainStorage
	return rb
}

func (rb *RepositoriesBuilder) WithSessionsStorage(sessionStorage redis.SessionStorage) *RepositoriesBuilder {
	rb.repositories.SessionRepository = sessionStorage
	return rb
}

func (rb *RepositoriesBuilder) Build() *Repositories {
	return &rb.repositories
}
