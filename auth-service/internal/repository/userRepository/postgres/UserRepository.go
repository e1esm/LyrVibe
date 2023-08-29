package postgres

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"github.com/e1esm/LyrVibe/auth-service/pkg/hash"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"time"
)

const (
	timeoutTime = 5 * time.Second
)

type UserStorage interface {
	Add(context.Context, *models.User) error
	GetOne(context.Context, string, string) *models.User
	UpdateRole(context.Context, uuid.UUID, models.Role) error
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(config config.Config) UserStorage {

	if err := godotenv.Load("postgres.env"); err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?pool_max_conns=%d",
		pgUser,
		pgPassword,
		config.UsersStorage.ContainerName,
		pgPort,
		pgDB,
		config.UsersStorage.MaxConnectionPool)
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		logger.GetLogger().Error(err.Error(), zap.String("url", databaseUrl))
		return nil
	}

	return &UserRepository{pool: pool}
}

func (ur *UserRepository) Add(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutTime)
	defer cancel()

	_, err := ur.pool.Exec(ctx, "INSERT INTO users VALUES($1, $2, $3, $4, $5);",
		user.ID,
		user.Username,
		user.Password,
		user.Role,
		user.ProfilePicture,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetOne(ctx context.Context, username, password string) *models.User {
	ctx, cancel := context.WithTimeout(ctx, timeoutTime)
	defer cancel()
	var user models.User
	resultedRow := ur.pool.QueryRow(ctx, "SELECT * FROM users WHERE username = $1", username)
	if err := resultedRow.Scan(&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.ProfilePicture); err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}

	if !hash.ComparePasswords(password, []byte(user.Password)) {
		return nil
	}
	return &user
}

func (ur *UserRepository) UpdateRole(ctx context.Context, id uuid.UUID, role models.Role) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutTime)
	defer cancel()
	_, err := ur.pool.Exec(ctx, "UPDATE users SET role = $1 WHERE id = $2", role, id)
	if err != nil {
		logger.GetLogger().Error("UpdateRole:UserRepository", zap.String("", err.Error()))
		return err
	}
	return nil
}
