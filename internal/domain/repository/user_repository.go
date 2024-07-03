package repository

import (
	"boilerplate-go/internal/domain/model"
	"context"
)

type UserRepository interface {
    GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
