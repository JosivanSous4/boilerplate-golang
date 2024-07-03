package repository

import (
	"boilerplate-go/internal/domain/model"
	"context"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
    db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) UserRepository {
    return &mysqlUserRepository{db: db}
}


func (r *mysqlUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
    var user model.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
