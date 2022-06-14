package repository

import (
	"context"
	"cleanArchCrud/models"
)

func(s *mysqlStorageRepository) GetAll(ctx context.Context) ([]models.StorageModel, error) {
	result := make([]models.StorageModel, 0)

	conn := s.ConnSql.Where("is_deleted = 0").Find(&result)
	if conn.Error != nil {
		return result, conn.Error
	}

	return result, nil
}

func(s *mysqlStorageRepository) AddProduct(ctx context.Context, data *models.StorageModel) (models.StorageModel, error) {
	res := s.ConnSql.Create(data)
	if res.Error != nil {
		return *data, res.Error
	}

	return *data, nil
}