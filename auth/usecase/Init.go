package usecase

import (
	"cleanArchCrud/auth/repository"
	"cleanArchCrud/models"
	"context"
	"time"
	general "cleanArchCrud/general/usecase"
)

type Authusecase interface {
	CreateUser(ctx context.Context, data *models.AuthModel) (string, error) 
	SignIn(ctx context.Context, data *models.AuthSignIn) (string, error)
	UpdateProfile(ctx context.Context, data *models.AuthUpdateProfile) error
	UpdatePassword(ctx context.Context, data *models.AuthUpdatePassword) error
}

type authUsecase struct {
	Authrepos repository.AuthRepository
	GeneralUsecase general.Usecase
	contextTimeout time.Duration
}

func NewAuthUsecase(AuthRepo repository.AuthRepository, gu general.Usecase,timeout time.Duration) Authusecase {
	return &authUsecase{ 
		Authrepos: AuthRepo,
		GeneralUsecase: gu,
		contextTimeout: timeout,
	}
}
