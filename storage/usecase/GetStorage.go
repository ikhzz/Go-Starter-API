package usecase

import (
	"context"
	"cleanArchCrud/models"
)

func(s *storageUseCase) GetAll(ctx context.Context, param models.StorageGetParam) ([]models.StorageModel, error) {

	result, err := s.StorageRepo.GetAll(ctx)
	if err != nil {
		return make([]models.StorageModel, 0), err
	}

	return result,nil
} 