package usecase

import (
	"log"
	"starterapi/common/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (h *commonUsecase) CreateToken(param models.JWTData) string {
	setToken := jwt.New(jwt.SigningMethodHS512)
	claims := setToken.Claims.(jwt.MapClaims)

	// depend on data on token use model
	claims["id"] = param.UidAuth
	claims["device"] = param.Device
	claims["iat"] = time.Now()
	token, err := setToken.SignedString([]byte(h.jwtKey))
	if err != nil {
		log.Println("create token:", err)
		return ""
	}

	return token
}

func (g *commonUsecase) CreateLog(param *models.LogModel) {
	err := g.cr.CreateLog(param)
	if err != nil {
		log.Println(err)
	}
}
