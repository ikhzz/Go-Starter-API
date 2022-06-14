package http

import (
	"cleanArchCrud/auth/usecase"
	general "cleanArchCrud/general/usecase"
	"cleanArchCrud/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthUsecase usecase.Authusecase
	GeneralUsecase general.Usecase
}

func NewAuthHandler(r *gin.Engine, au usecase.Authusecase, h general.Usecase) {
	router := &AuthHandler{
		AuthUsecase: au,
		GeneralUsecase: h,
	}
	v1 := r.Group("/v1")
	{
		v1.POST("/signup", router.SignUp)
		v1.POST("/signin", router.SignIn)
		v1.POST("/updateProfile", router.UpdateProfile)
		v1.POST("/updatePassword", router.UpdatePassword)
	}
}

func(a *AuthHandler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var authModel models.AuthModel
	c.Bind(&authModel)
	errvalid := validator.New().Struct(&authModel)
	if errvalid != nil {
		res := a.GeneralUsecase.ValidatorHelper(errvalid.(validator.ValidationErrors))
		c.JSON(403, gin.H{"status": "failed", "error": res})
		return
	}
	token, err := a.AuthUsecase.CreateUser(ctx, &authModel)
	if err != nil {
		c.JSON(403, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success", "token":token})
}

func(a *AuthHandler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()
	var authModel models.AuthSignIn
	c.Bind(&authModel)
	errvalid := validator.New().Struct(&authModel)
	if errvalid != nil {
		res := a.GeneralUsecase.ValidatorHelper(errvalid.(validator.ValidationErrors))
		c.JSON(403, gin.H{"status": "failed", "error": res})
		return
	}
	token, errToken := a.AuthUsecase.SignIn(ctx, &authModel)
	if errToken != nil {
		c.JSON(403, gin.H{"status": "failed", "error": errToken.Error()})
		return
	}
	ctx.Done()
	c.JSON(200, gin.H{"status": "success", "token":token})
}

func(a *AuthHandler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()
	var authModel models.AuthUpdateProfile
	c.Bind(&authModel)
	errvalid := validator.New().Struct(&authModel)
	if errvalid != nil {
		res := a.GeneralUsecase.ValidatorHelper(errvalid.(validator.ValidationErrors))
		c.JSON(403, gin.H{"status": "failed", "error": res})
		return
	}
	id, _ := c.Get("payload_id")
	role, _ := c.Get("payload_role")
	ctx = context.WithValue(ctx, "payload_id", id)
	ctx = context.WithValue(ctx, "payload_role", role)
	errResult := a.AuthUsecase.UpdateProfile(ctx, &authModel)
	if errResult != nil {
		c.JSON(403, gin.H{"status": "failed", "error": errResult.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": "success change profile data"})
}

func(a *AuthHandler) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	var authModel models.AuthUpdatePassword
	c.Bind(&authModel)
	errvalid := validator.New().Struct(&authModel)
	if errvalid != nil {
		res := a.GeneralUsecase.ValidatorHelper(errvalid.(validator.ValidationErrors))
		c.JSON(403, gin.H{"status": "failed", "error": res})
		return
	}
	id, _ := c.Get("payload_id")
	role, _ := c.Get("payload_role")
	ctx = context.WithValue(ctx, "payload_id", id)
	ctx = context.WithValue(ctx, "payload_role", role)
	errResult := a.AuthUsecase.UpdatePassword(ctx, &authModel)
	if errResult != nil {
		c.JSON(403, gin.H{"status": "failed", "error": errResult.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data":"success to change password"})
}