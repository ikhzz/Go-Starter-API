package usecase

import (
	"cleanArchCrud/models"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (g *GeneralUsecase) ValidatorHelper(ers validator.ValidationErrors) (allerror []models.AuthResponse) {
	fmt.Println(ers)
	for _, v := range ers {
		var res models.AuthResponse
		res.Field = v.StructField()
		res.Msg = g.validatorMessage(v.ActualTag(), v.Param())
		allerror = append(allerror, res)
	}
	return 
}


func (g *GeneralUsecase) validatorMessage(tag string, param string) (str string) {
	switch tag {
		case "required":
			str += "field is required, "
		case "gte":
			str += "field is more than " +param + " character"
		case "email":
			str += "request is not a valid email"
	} 
	return
}

func (g *GeneralUsecase) TokenCreate(s models.TokenRequirement) (string, error) {
	setToken := jwt.New(jwt.SigningMethodHS256)
	claims := setToken.Claims.(jwt.MapClaims)

	claims["id"] = s.Id
	claims["role"] = s.Role

	token, err := setToken.SignedString(SignedString)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (g *GeneralUsecase) TokenDecrypt(c *gin.Context) {

	bearer := c.GetHeader("Authorization")
	strSplit := strings.Split(bearer, " ")
	// if token doesnt't have 3 parts it is unhanled error on token claims
	tokenSection := 0
	for _, t := range strSplit[1] {
		if string(t) == "." {
			tokenSection++
		}
	}
	if bearer != "" && len(strSplit) == 2 && tokenSection == 2{
		token, err := jwt.Parse(strSplit[1], func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(SignedString), nil
		})
		claims,ok := token.Claims.(jwt.MapClaims)
		if err == nil && ok && token.Valid {
			c.Set("payload_id",  claims["id"])
			c.Set("payload_role", claims["role"]) 
		} else {
			c.JSON(403, gin.H{
				"status":  false,
				"message": "route required valid token",
			})
			c.AbortWithStatus(403)	
		}
	} else {
		c.JSON(403, gin.H{
			"status":  false,
			"message": "route required auth token",
		})
		c.AbortWithStatus(403)
	}
}

func(g *GeneralUsecase) PasswordHash(s string) (string, error) {
	passbyte, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return "", err
	}
	return string(passbyte), nil
}

func(g *GeneralUsecase) PasswordCompare(pass string, compare string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(compare))
	fmt.Println(err)
	return
}