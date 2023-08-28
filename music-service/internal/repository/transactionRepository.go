package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	TMTransactionNotFound = errors.New("no transaction was found")
)

type TransactionManager interface {
	AddTx(uuid.UUID, pgx.Tx)
	Delete(uuid.UUID)
	Get(uuid uuid.UUID) (pgx.Tx, error)
}

type TransactionRepository struct {
	repo map[uuid.UUID]pgx.Tx
}

func NewTransactionRepository() TransactionManager {
	return &TransactionRepository{
		repo: make(map[uuid.UUID]pgx.Tx),
	}
}

func (t *TransactionRepository) AddTx(id uuid.UUID, tx pgx.Tx) {
	t.repo[id] = tx
}

func (t *TransactionRepository) Delete(id uuid.UUID) {
	delete(t.repo, id)
}

func (t *TransactionRepository) Get(id uuid.UUID) (pgx.Tx, error) {
	tx, ok := t.repo[id]
	if !ok {
		return tx, TMTransactionNotFound
	}
	return tx, nil
}
