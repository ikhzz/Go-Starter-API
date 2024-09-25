package models

import (
	"starterapi/common/models"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type NotesUsecase interface {
	GetNotes(c *gin.Context, param *ReqGetNotes) *models.GeneralResponse
	PostNotes(c *gin.Context, param *ReqPostNotes) *models.GeneralResponse
	PutNotes(c *gin.Context, param *ReqPutNotes) *models.GeneralResponse
	DeleteNotes(c *gin.Context, param *ReqDeleteNotes) *models.GeneralResponse
}

type NotesRepository interface {
	GetNotes(param *ReqGetNotes, id *string) ([]NotesData, int64, error)
	PostNotes(param *ReqPostNotes) error
	PutNotes(param string, toUpdate map[string]interface{}) error
}

type (
	NotesModel struct {
		IDNotes     int       `json:"id" gorm:"primaryKey;autoIncrement"`
		UIDNotes    uuid.UUID `json:"uid" gorm:"type:char(36)"`
		Title       string    `json:"title" gorm:"default:null;type:text"`
		Description string    `json:"description" gorm:"default:null;type:text"`
		CreatedBy   string    `json:"-"`
		UpdatedBy   string    `json:"-"`
		DeletedBy   string    `json:"-"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
		DeletedAt   time.Time `json:"-" gorm:"default:null"`
	}

	ReqGetNotes struct {
		Keyword string `form:"keyword" example:"title"`
		Limit   int    `form:"limit" validate:"number" example:"10"`
		Offset  int    `form:"offset" validate:"number" example:"0"`
		All     int    `form:"all" validate:"number" example:"0"`
		Order   string `form:"order" example:"title"`
		OrderBy string `form:"order_by" validate:"omitempty,oneof=asc desc" example:"asc"`
	}

	ReqPostNotes struct {
		UIDNotes    uuid.UUID `form:"-" json:"-"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description"`
		CreatedBy   string    `json:"-"`
		UpdatedBy   string    `json:"-"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
	}

	ReqPutNotes struct {
		UIDNotes    string `form:"uid_notes" json:"uid_notes" validate:"required"`
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
	}

	ReqDeleteNotes struct {
		UIDNotes string `form:"uid_notes" json:"uid_notes" validate:"required"`
	}

	NotesData struct {
		UIDNotes    uuid.UUID `json:"uid" example:"uuid"`
		Title       string    `json:"title" example:"title"`
		Description string    `json:"description" example:"description"`
		CreatedAt   time.Time `json:"created_at" example:"timestamp"`
	}
)

func (NotesModel) TableName() string {
	return "notes"
}

func (a *NotesModel) BeforeCreate(stmt *gorm.DB) error {
	a.UIDNotes = uuid.NewV4()

	return nil
}

func (a *ReqPostNotes) BeforeCreate(stmt *gorm.DB) error {
	a.UIDNotes = uuid.NewV4()

	return nil
}
