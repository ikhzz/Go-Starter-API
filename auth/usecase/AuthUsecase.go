package usecase

import (
	"cleanArchCrud/models"
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/google/uuid"
)

func (a *authUsecase) CreateUser(ctx context.Context, data *models.AuthModel) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	if data.Role != "seller" {
		data.Role = "buyer"
	}

	password, err := a.GeneralUsecase.PasswordHash(data.Password)
	data.CreatedAt = time.Now()
	if err != nil {
		return "",err
	}
	data.Password = password
	data.Id = uuid.NewString()
	err = a.Authrepos.Create(ctx, data)
	if err != nil {
		return "",err
	}
	fmt.Println("Create seller", data)
	tokenReq := models.TokenRequirement{Id: data.Id, Role: data.Role}
	token, errToken := a.GeneralUsecase.TokenCreate(tokenReq)
	
	if errToken != nil {
		return "", errToken	
	}

	return token, nil;
}

func (a *authUsecase) SignIn(ctx context.Context, data *models.AuthSignIn) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	result, err := a.Authrepos.SignIn(ctx, data)
	fmt.Println(result)
	if err != nil {
		return "", err
	}
	errCompare := a.GeneralUsecase.PasswordCompare(result.Password, data.Password)
	if errCompare != nil {
		return "", errors.New("wrong password")
	}
	tokenReq := models.TokenRequirement{Id: result.Id, Role: result.Role}
	token, errToken := a.GeneralUsecase.TokenCreate(tokenReq)
	if errToken != nil {
		return "", errToken	
	}
	return token, err
}



