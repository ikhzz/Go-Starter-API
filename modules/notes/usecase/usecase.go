package usecase

import (
	"net/http"
	"starterapi/common/models"
	nomod "starterapi/modules/notes/models"
	"starterapi/modules/notes/repository"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type notesUsecase struct {
	notesRepo nomod.NotesRepository
	common    models.CommonUsecase
}

func NewNotesUsecase(conn *gorm.DB, common models.CommonUsecase) nomod.NotesUsecase {
	repo := repository.NewNotesRepository(conn)

	return notesUsecase{notesRepo: repo, common: common}
}

func (nu notesUsecase) GetNotes(c *gin.Context, param *nomod.ReqGetNotes) *models.GeneralResponse {
	id := c.GetString("user_id")
	data, count, err := nu.notesRepo.GetNotes(param, &id)
	if err != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    err.Error(),
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		nu.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusNotFound,
			Status:     false,
			Message:    []string{err.Error()},
			Data:       models.EmptyResponse{},
		}
		return &res
	}

	res := models.GeneralResponse{
		StatusCode:  http.StatusOK,
		Status:      true,
		Message:     []string{"success"},
		Data:        data,
		PagesLength: int(count),
	}

	return &res
}

func (nu notesUsecase) PostNotes(c *gin.Context, param *nomod.ReqPostNotes) *models.GeneralResponse {
	id := c.GetString("user_id")
	param.CreatedBy = id
	param.UpdatedBy = id
	err := nu.notesRepo.PostNotes(param)
	if err != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    err.Error(),
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		nu.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusBadRequest,
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
		Data:       models.EmptyResponse{},
	}

	return &res
}

func (nu notesUsecase) PutNotes(c *gin.Context, param *nomod.ReqPutNotes) *models.GeneralResponse {
	id := c.GetString("user_id")
	toUpdate := make(map[string]interface{}, 0)
	toUpdate["title"] = param.Title
	if param.Description != "" {
		toUpdate["description"] = param.Description
	}
	toUpdate["updated_by"] = id
	toUpdate["updated_at"] = time.Now()

	err := nu.notesRepo.PutNotes(param.UIDNotes, toUpdate)
	if err != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    err.Error(),
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		nu.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusBadRequest,
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
		Data:       models.EmptyResponse{},
	}

	return &res
}

func (nu notesUsecase) DeleteNotes(c *gin.Context, param *nomod.ReqDeleteNotes) *models.GeneralResponse {
	id := c.GetString("user_id")
	toUpdate := make(map[string]interface{}, 0)
	toUpdate["deleted_at"] = time.Now()
	toUpdate["deleted_by"] = id

	err := nu.notesRepo.PutNotes(param.UIDNotes, toUpdate)
	if err != nil {
		resLog := models.LogModel{
			AccessType: "error",
			LogData:    err.Error(),
			LogType:    c.Request.Method,
			EndPoint:   c.Request.URL.Path,
		}
		nu.common.CreateLog(&resLog)

		res := models.GeneralResponse{
			StatusCode: http.StatusBadRequest,
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
		Data:       models.EmptyResponse{},
	}

	return &res
}
