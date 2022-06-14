package usecase

import (
	"cleanArchCrud/storage/repository"
	"cleanArchCrud/models"
	"context"
	"time"
	general "cleanArchCrud/general/usecase"
)

type StorageUsecase interface {
	// CreateUser(ctx context.Context, data *models.Auth) (string, error) 
	// SignIn(ctx context.Context, data *models.AuthSignIn) (string, error)
	GetAll(ctx context.Context, param models.StorageGetParam) ([]models.StorageModel, error)
	AddProduct(ctx context.Context, param models.StorageModel) (models.StorageModel, error)
}

type storageUseCase struct {
	StorageRepo repository.StorageRepository
	GeneralUsecase general.Usecase
	contextTimeout time.Duration
}

func NewStorageUsecase(StorageRepo repository.StorageRepository, gu general.Usecase,timeout time.Duration) StorageUsecase {
	return &storageUseCase{ 
		StorageRepo: StorageRepo,
		GeneralUsecase: gu,
		contextTimeout: timeout,
	}
}
