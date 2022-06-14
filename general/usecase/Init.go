package usecase

import (
	"cleanArchCrud/general/repository"
	"cleanArchCrud/models"
	"context"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	SignedString []byte
)

type Usecase interface {
	ValidatorHelper(ers validator.ValidationErrors) (allerror []models.AuthResponse)
	PanicCatcher(mw io.Writer) gin.HandlerFunc
	CustomLogger(mw io.Writer) gin.HandlerFunc
	TokenCreate(s models.TokenRequirement) (string, error)
	PasswordHash(s string) (string, error)
	PasswordCompare(pass string, compare string) error
	CheckRoute() gin.HandlerFunc
	GetUsername(ctx context.Context, id string) string
	UploadFile(ctx context.Context, file models.FileModel) (models.FileResponse, error)
}

type GeneralUsecase struct {
	contextTimeout time.Duration
	Generalrepos repository.GeneralRepository
}

func NewGeneralUsecase(timeout time.Duration, GeneralRepo repository.GeneralRepository) Usecase {
	getKey := viper.GetString("jwt.key")
	SignedString = []byte(getKey)
	if getKey == "" {
		panic("jwt key is missing")
	}

	return &GeneralUsecase{ 
		contextTimeout: timeout,
		Generalrepos: GeneralRepo,
	}
}