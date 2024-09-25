package models

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommonUsecase interface {
	PanicCatcher(mw io.Writer) gin.HandlerFunc
	CustomLogger(mw io.Writer) gin.HandlerFunc
	CreateLog(param *LogModel)
	Validate(param interface{}) (error, []string)
	CreateToken(param JWTData) string
	JwtMiddleware(c *gin.Context)
}

type CommonRepository interface {
	CreateLog(param *LogModel) error
	FindUserByUid(param string) (res UserData, err error)
	PutUser(param string, toUpdate map[string]interface{}) error
}

type (
	GeneralResponse struct {
		StatusCode  int         `json:"status_code"`
		Status      bool        `json:"status"`
		Message     []string    `json:"messages"`
		Data        interface{} `json:"data"`
		PagesLength int         `json:"pages_length"`
	}

	TokenPayload struct {
		UidAuth string
		Device  string
	}

	JWTData struct {
		UidAuth string
		Device  string
	}

	EmptyResponse struct {
	}

	LogModel struct {
		IDlog      int       `json:"id" gorm:"primaryKey;autoIncrement"`
		AccessType string    `json:"access_type" gorm:"varchar(255)"`
		LogType    string    `json:"log_type" gorm:"varchar(255)"`
		EndPoint   string    `json:"end_point" gorm:"varchar(255)"`
		LogData    string    `json:"log_data" gorm:"text"`
		CreatedBy  string    `json:"-"`
		UpdatedBy  string    `json:"-"`
		DeletedBy  string    `json:"-"`
		CreatedAt  time.Time `json:"-"`
		UpdatedAt  time.Time `json:"-"`
		DeletedAt  time.Time `json:"-" gorm:"default:null"`
	}

	UserData struct {
		UIDUser     string    `json:"-" `
		Token       string    `json:"-"`
		TokenExpire time.Time `json:"-"`
	}
)

func (LogModel) TableName() string {
	return "log"
}

func (a *LogModel) BeforeCreate(stmt *gorm.DB) error {
	// a.DeletedAt = nil

	return nil
}
