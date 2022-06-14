package http

import (
	general "cleanArchCrud/general/usecase"
	"cleanArchCrud/models"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GeneralHandler struct {
	GeneralUsecase general.Usecase
}

func NewGeneralHandler(r *gin.Engine, h general.Usecase) {
	router := &GeneralHandler{
		GeneralUsecase: h,
	}
	v1 := r.Group("/v1")
	{
		v1.POST("/uploadFile", router.UploadFile)
	}
}

func(g *GeneralHandler) UploadFile(c *gin.Context) {
	ctx := c.Request.Context()
	id, _ := c.Get("payload_id")
	role, _ := c.Get("payload_role")
	ctx = context.WithValue(ctx, "payload_id", id)
	ctx = context.WithValue(ctx, "payload_role", role)
	var fileModel models.FileModel
	c.Bind(&fileModel)
	fmt.Println(fileModel)
	fileResponse, err := g.GeneralUsecase.UploadFile(ctx, fileModel)
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data":fileResponse})
}