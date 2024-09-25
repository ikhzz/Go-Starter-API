package routes

import (
	"starterapi/common/models"
	authdev "starterapi/modules/auth/delivery"
	nodev "starterapi/modules/notes/delivery"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// add dependency required on all modules
func InitRoutes(c *gin.Engine, conn *gorm.DB, cmnu models.CommonUsecase) {

	authdev.NewAuthDelivery(c, conn, cmnu)
	nodev.NewNotesDelivery(c, conn, cmnu)
}
