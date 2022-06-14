package repository

import (
	"context"
	"cleanArchCrud/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(ctx context.Context,data *models.AuthModel) error
	SignIn(ctx context.Context, data *models.AuthSignIn) (models.AuthResult, error)
	UpdateProfile(ctx context.Context, data *models.AuthUpdateProfile) error
	GetPasswordById(ctx context.Context) (string, error)
	UpdatePassword(ctx context.Context, newpass string) error
}

type mysqlAuthRepository struct {
	ConnSql *gorm.DB
	// can add nosql db
}

func NewMysqlAuthRepository(Conn *gorm.DB) AuthRepository {
	return &mysqlAuthRepository{Conn}
}