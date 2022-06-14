package usecase

import (
	"cleanArchCrud/models"
	"context"
	"errors"
	"fmt"
	"time"
)

func(s *storageUseCase) AddProduct(ctx context.Context, param models.StorageModel) (models.StorageModel, error) {
	role := fmt.Sprintf("%v", ctx.Value("payload_role"))
	if role != "seller" {
		return param, errors.New("buyer can't add product")
	}
	param.CreatedAt = time.Now()
	
	param.ProductOwner = s.GeneralUsecase.GetUsername(ctx, fmt.Sprintf("%v",ctx.Value("payload_id")))

	result, err := s.StorageRepo.AddProduct(ctx, &param)
	if err != nil {
		return result, err
	}

	return result,nil
} 