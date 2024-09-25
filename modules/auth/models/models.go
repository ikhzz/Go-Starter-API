package models

import (
	"starterapi/common/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	SignIn(c *gin.Context, reqParam *ReqPostSignin) *models.GeneralResponse
	GetProfile(c *gin.Context, param *ReqGetProfile) *models.GeneralResponse
}

type AuthRepository interface {
	FindUserByEmail(param string) (res UserData, err error)
	FindUserByUid(param string) (res UserData, err error)
	PutUser(param string, toUpdate map[string]interface{}) error
}

type (
	UsersModel struct {
		IDUser      int       `json:"id" gorm:"primaryKey;autoIncrement"`
		UIDUser     uuid.UUID `json:"uid" gorm:"type:char(36)"`
		Email       string    `json:"email" gorm:"default:null;type:varchar(50)"`
		Username    string    `gorm:"default:null;type:varchar(30)"`
		Password    string    `json:"-" gorm:"default:null;type:varchar(255)"`
		IsActive    int       `json:"-" gorm:"tinyint(1);default:1"`
		Token       string    `json:"-" gorm:"type:text;default:null"`
		TokenExpire time.Time `json:"-" gorm:"default:null"`
		CreatedBy   string    `json:"-"`
		UpdatedBy   string    `json:"-"`
		DeletedBy   string    `json:"-"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
		DeletedAt   time.Time `json:"-" gorm:"default:null"`
	}

	ReqPostSignin struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=6"`
	}

	ReqGetProfile struct {
		Device string `form:"device" validate:"required,oneof=WEB IPHONE ANDROID"`
	}

	UserData struct {
		IDUser   int    `json:"-" `
		UIDUser  string `json:"uid" example:"uid"`
		Email    string `json:"email" example:"example@mail.com"`
		Username string `json:"username" example:"user_example"`
		Password string `json:"-"`
	}

	ResSignin struct {
		Token string `json:"token" example:"token"`
	}
)

func (UsersModel) TableName() string {
	return "users"
}

func (a *UsersModel) BeforeCreate(stmt *gorm.DB) error {
	a.UIDUser = uuid.NewV4()

	return nil
}
