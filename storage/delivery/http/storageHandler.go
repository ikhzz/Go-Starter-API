package http

import (
	general "cleanArchCrud/general/usecase"
	"cleanArchCrud/models"
	"cleanArchCrud/storage/usecase"
	"context"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StorageHandler struct {
	StorageUsecase usecase.StorageUsecase
	GeneralUsecase general.Usecase
}

func NewStorageHandler(r *gin.Engine, au usecase.StorageUsecase, h general.Usecase) {
	router := &StorageHandler{
		StorageUsecase: au,
		GeneralUsecase: h,
	}
	v1 := r.Group("/v1")
	{
		v1.GET("/getProducts", router.GetProducts)
		v1.POST("/addProduct", router.AddProduct)
	}
}

func(a *StorageHandler) GetProducts(c *gin.Context) {
	ctx := c.Request.Context()
	// id, _ := c.Get("payload_id")
	// role, _ := c.Get("payload_role")
	// fmt.Println(ctx)
	// fmt.Println(id)
	// fmt.Println(role)
	var param models.StorageGetParam
	errBinding := c.Bind(&param)

	if errBinding != nil {
		c.JSON(403, gin.H{"status": "failed", "error": "request parameter is not valid"})
		return
	}
	
	res, err := a.StorageUsecase.GetAll(ctx, param)
	if err != nil {
		c.JSON(403, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data":res})
}

func(a *StorageHandler) AddProduct(c *gin.Context) {
	ctx := c.Request.Context()
	id, _ := c.Get("payload_id")
	role, _ := c.Get("payload_role")
	ctx = context.WithValue(ctx, "payload_id", id)
	ctx = context.WithValue(ctx, "payload_role", role)

	// fmt.Println(ctx.Value("payload_id"))
	// fmt.Println(id)
	// fmt.Println(role)
	var param models.StorageModel
	c.Bind(&param)
	errvalid := validator.New().Struct(&param)
	if errvalid != nil {
		res := a.GeneralUsecase.ValidatorHelper(errvalid.(validator.ValidationErrors))
		c.JSON(403, gin.H{"status": "failed", "error": res})
		return
	}
	res, err := a.StorageUsecase.AddProduct(ctx, param)
	ctx.Done()
	if err != nil {
		c.JSON(403, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data":res})
	
}