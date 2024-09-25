package usecase

import (
	"net/http"
	"starterapi/common/models"
	modauth "starterapi/modules/auth/models"
	"starterapi/modules/auth/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authUsecase struct {
	authRepo modauth.AuthRepository
	common   models.CommonUsecase
}

func NewAuthUsecase(conn *gorm.DB, common models.CommonUsecase) modauth.AuthUsecase {
	repo := repository.NewAuthRepository(conn)

	return authUsecase{authRepo: repo, common: common}
}

func (a authUsecase) SignIn(c *gin.Context, param *modauth.ReqPostSignin) *models.GeneralResponse {
	data, err := a.authRepo.FindUserByEmail(param.Email)
	if err != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    err.Error(),
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		a.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusNotFound,
			Status:     false,
			Message:    []string{err.Error()},
			Data:       models.EmptyResponse{},
		}
		return &res
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(param.Password))
	if errCompare != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    param.Email + " password missmatch",
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		a.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusBadRequest,
			Status:     false,
			Message:    []string{"password missmatch"},
			Data:       models.EmptyResponse{},
		}
		return &res
	}

	token := a.common.CreateToken(models.JWTData{UidAuth: data.UIDUser, Device: "WEB"})
	toUpdate := make(map[string]interface{}, 0)
	toUpdate["token"] = token
	toUpdate["token_expire"] = time.Now().Add(time.Hour * time.Duration(viper.GetInt("timeout.jwt")))
	toUpdate["updated_at"] = time.Now()
	errPut := a.authRepo.PutUser(data.UIDUser, toUpdate)
	if errPut != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    errPut.Error(),
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		a.common.CreateLog(&resLog)
	}

	res := models.GeneralResponse{
		StatusCode: http.StatusOK,
		Status:     true,
		Message:    []string{"success"},
		Data:       modauth.ResSignin{Token: token},
	}

	return &res
}

func (a authUsecase) GetProfile(c *gin.Context, param *modauth.ReqGetProfile) *models.GeneralResponse {
	id := c.GetString("user_id")
	data, err := a.authRepo.FindUserByUid(id)
	if err != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    "data not found",
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		a.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusUnauthorized,
			Status:     false,
			Message:    []string{err.Error()},
			Data:       models.EmptyResponse{},
		}
		return &res
	}

	res := models.GeneralResponse{
		StatusCode: http.StatusOK,
		Status:     true,
		Message:    []string{"success"},
		Data:       data,
	}

	return &res
}
