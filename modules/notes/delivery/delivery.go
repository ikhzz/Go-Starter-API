package delivery

import (
	"encoding/json"
	"net/http"
	cmmd "starterapi/common/models"
	"starterapi/modules/notes/models"
	"starterapi/modules/notes/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NotesDelivery struct {
	Usecase models.NotesUsecase
	Helper  cmmd.CommonUsecase
}

func NewNotesDelivery(c *gin.Engine, conn *gorm.DB, common cmmd.CommonUsecase) {
	nouc := usecase.NewNotesUsecase(conn, common)

	handler := &NotesDelivery{
		Usecase: nouc,
		Helper:  common,
	}

	v1 := c.Group("/v1/notes")
	v1.Use(common.JwtMiddleware)
	{
		v1.GET("/", handler.GetNotes)
		v1.POST("/", handler.PostNotes)
		v1.PUT("/", handler.PutNotes)
		v1.DELETE("/", handler.DeleteNotes)
	}
}

// @Summary Get All Notes
// @Schemes
// @Description api to get all notes
// @Tags Notes
// @Security BearerAuth
// @Produce json
// @Param GetNotes query models.ReqGetNotes true "Request Param Get Notes"
// @Success 200 {object} models.SwaggoGetNotesResSuccess
// @Failure 400 {object} models.SwaggoGetNotesResBadRequest
// @Failure 401 {object} models.SwaggoGetNotesResUnauthorized
// @Failure 404 {object} models.SwaggoGetNotesResNotFound
// @Router /notes [get]
func (a NotesDelivery) GetNotes(c *gin.Context) {
	var reqParam models.ReqGetNotes
	c.Bind(&reqParam)

	err, msg := a.Helper.Validate(reqParam)
	if err != nil {
		res := cmmd.GeneralResponse{
			StatusCode: http.StatusBadRequest,
			Status:     false,
			Message:    msg,
			Data:       cmmd.EmptyResponse{},
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	b, _ := json.Marshal(reqParam)
	logParam := cmmd.LogModel{
		AccessType: "request",
		LogType:    c.Request.Method,
		EndPoint:   c.Request.URL.Path,
		LogData:    string(b),
	}

	a.Helper.CreateLog(&logParam)

	res := a.Usecase.GetNotes(c, &reqParam)
	c.JSON(res.StatusCode, res)
}

// @Summary Post Notes
// @Schemes
// @Description api to create notes
// @Tags Notes
// @Security BearerAuth
// @Produce json
// @Param PostNotes body models.ReqPostNotes true "Request Param Post Notes"
// @Success 200 {object} models.SwaggoPostNotesResSuccess
// @Failure 400 {object} models.SwaggoPostNotesResBadRequest
// @Failure 401 {object} models.SwaggoPostNotesResUnauthorized
// @Router /notes [post]
func (a NotesDelivery) PostNotes(c *gin.Context) {
	var reqParam models.ReqPostNotes
	c.Bind(&reqParam)

	err, msg := a.Helper.Validate(reqParam)
	if err != nil {
		res := cmmd.GeneralResponse{
			StatusCode: http.StatusBadRequest,
			Status:     false,
			Message:    msg,
			Data:       cmmd.EmptyResponse{},
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	b, _ := json.Marshal(reqParam)
	logParam := cmmd.LogModel{
		AccessType: "request",
		LogType:    c.Request.Method,
		EndPoint:   c.Request.URL.Path,
		LogData:    string(b),
	}

	a.Helper.CreateLog(&logParam)

	res := a.Usecase.PostNotes(c, &reqParam)
	c.JSON(res.StatusCode, res)
}

// @Summary Put Notes
// @Schemes
// @Description api to update notes
// @Tags Notes
// @Security BearerAuth
// @Produce json
// @Param PutNotes body models.ReqPutNotes true "Request Param Put Notes"
// @Success 200 {object} models.SwaggoPutNotesResSuccess
// @Failure 400 {object} models.SwaggoPutNotesResBadRequest
// @Failure 401 {object} models.SwaggoPutNotesResUnauthorized
// @Router /notes [put]
func (a NotesDelivery) PutNotes(c *gin.Context) {
	var reqParam models.ReqPutNotes
	c.Bind(&reqParam)

	err, msg := a.Helper.Validate(reqParam)
	if err != nil {
		res := cmmd.GeneralResponse{
			StatusCode: http.StatusBadRequest,
			Status:     false,
			Message:    msg,
			Data:       cmmd.EmptyResponse{},
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	b, _ := json.Marshal(reqParam)
	logParam := cmmd.LogModel{
		AccessType: "request",
		LogType:    c.Request.Method,
		EndPoint:   c.Request.URL.Path,
		LogData:    string(b),
	}

	a.Helper.CreateLog(&logParam)

	res := a.Usecase.PutNotes(c, &reqParam)
	c.JSON(res.StatusCode, res)
}

// @Summary Delete Notes
// @Schemes
// @Description api to delete notes
// @Tags Notes
// @Security BearerAuth
// @Produce json
// @Param DeleteNotes body models.ReqDeleteNotes true "Request Param Delete Notes"
// @Success 200 {object} models.SwaggoDeleteNotesResSuccess
// @Failure 400 {object} models.SwaggoDeleteNotesResBadRequest
// @Failure 401 {object} models.SwaggoDeleteNotesResUnauthorized
// @Router /notes [delete]
func (a NotesDelivery) DeleteNotes(c *gin.Context) {
	var reqParam models.ReqDeleteNotes
	c.Bind(&reqParam)

	err, msg := a.Helper.Validate(reqParam)
	if err != nil {
		res := cmmd.GeneralResponse{
			StatusCode: http.StatusBadRequest,
			Status:     false,
			Message:    msg,
			Data:       cmmd.EmptyResponse{},
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	b, _ := json.Marshal(reqParam)
	logParam := cmmd.LogModel{
		AccessType: "request",
		LogType:    c.Request.Method,
		EndPoint:   c.Request.URL.Path,
		LogData:    string(b),
	}

	a.Helper.CreateLog(&logParam)

	res := a.Usecase.DeleteNotes(c, &reqParam)
	c.JSON(res.StatusCode, res)
}
