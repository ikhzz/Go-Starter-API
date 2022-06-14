package repository

import (
	"context"

	"gorm.io/gorm"
)

type GeneralRepository interface {
	GetUsername(ctx context.Context, id string) string
}

type mysqlGeneralRepository struct {
	ConnSql *gorm.DB
	// can add nosql db
}

func NewGeneralRepository(Conn *gorm.DB) GeneralRepository {
	return &mysqlGeneralRepository{Conn}
}