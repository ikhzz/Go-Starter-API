package delivery

import (
	"encoding/json"
	"net/http"
	cmmd "starterapi/common/models"
	"starterapi/modules/auth/models"
	"starterapi/modules/auth/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthDelivery struct {
	Usecase models.AuthUsecase
	Helper  cmmd.CommonUsecase
}

func NewAuthDelivery(c *gin.Engine, conn *gorm.DB, common cmmd.CommonUsecase) {
	usecaseAuth := usecase.NewAuthUsecase(conn, common)

	handler := &AuthDelivery{
		Usecase: usecaseAuth,
		Helper:  common,
	}

	v1 := c.Group("/v1/auth")
	{
		v1.POST("/signin", handler.Signin)
	}
	v1.Use(common.JwtMiddleware)
	{
		v1.GET("/profile", handler.GetProfile)
	}
}

// @Summary Signin
// @Schemes
// @Description sign in api
// @Tags Authentication
// @Produce json
// @Param PostSignin body models.ReqPostSignin true "Request Body signin"
// @Success 200 {object} models.SwaggoSigninResSuccess
// @Failure 400 {object} models.SwaggoSigninResBadRequest
// @Failure 404 {object} models.SwaggoSigninResNotFound
// @Router /auth/signin [post]
func (a AuthDelivery) Signin(c *gin.Context) {
	var reqParam models.ReqPostSignin
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

	res := a.Usecase.SignIn(c, &reqParam)
	c.JSON(res.StatusCode, res)
}

// @Summary GetProfile
// @Schemes
// @Description api to get profile user
// @Tags Authentication
// @Security BearerAuth
// @Produce json
// @Param GetProfile query models.ReqGetProfile true "Request Body profile"
// @Success 200 {object} models.SwaggoGetProfileResSuccess
// @Failure 400 {object} models.SwaggoGetProfileResBadRequest
// @Failure 401 {object} models.SwaggoGetProfileResUnauthorized
// @Router /auth/profile [get]
func (a AuthDelivery) GetProfile(c *gin.Context) {
	var reqParam models.ReqGetProfile
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

	device := c.GetString("device")
	if reqParam.Device != device {
		res := cmmd.GeneralResponse{
			StatusCode: http.StatusBadRequest,
			Status:     false,
			Message:    []string{"invalid device"},
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

	res := a.Usecase.GetProfile(c, &reqParam)
	c.JSON(res.StatusCode, res)
}
