package repository

import (
	"context"
	"cleanArchCrud/models"
	"gorm.io/gorm"
)

type StorageRepository interface {
	// Create(ctx context.Context,data *models.Auth) error
	// SignIn(ctx context.Context, data *models.AuthSignIn) (models.AuthResult, error)
	GetAll(ctx context.Context) ([]models.StorageModel, error)
	AddProduct(ctx context.Context, param *models.StorageModel) (models.StorageModel, error)
}

type mysqlStorageRepository struct {
	ConnSql *gorm.DB
	// can add nosql db
}

func NewMysqlStorageRepository(Conn *gorm.DB) StorageRepository {
	return &mysqlStorageRepository{Conn}
}