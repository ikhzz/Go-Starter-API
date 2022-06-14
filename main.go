package main

import (
	"fmt"
	"os"
	// "bytes"
	"log"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"cleanArchCrud/db"
	authRepository "cleanArchCrud/auth/repository"
	authUsecase "cleanArchCrud/auth/usecase"
	authHandler "cleanArchCrud/auth/delivery/http"

	generalRepository "cleanArchCrud/general/repository"
	generalUsecase "cleanArchCrud/general/usecase"
	generalHandler "cleanArchCrud/general/delivery/http"
	
	storageRepository "cleanArchCrud/storage/repository"
	storageUsecase "cleanArchCrud/storage/usecase"
	storageHandler "cleanArchCrud/storage/delivery/http"
)




func main() {

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	
	f, _ := os.Create("./log/service.log")
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(f)

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	appPort := viper.GetString("address")
	if appPort == "" {
		appPort = ":6060"
	}

	connMysql := db.InitMysql()
	// connMonggo := db.InitMonggoDB()
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// router.LoadHTMLGlob("templates/**/*")
	
	// helper function
	gr := generalRepository.NewGeneralRepository(connMysql)
	gu := generalUsecase.NewGeneralUsecase(timeoutContext, gr)
	
	
	router.Use(gin.Recovery())
	router.Use(gu.PanicCatcher(mw))
	router.Use(gu.CustomLogger(mw))
	router.Use(gu.CheckRoute())
	// general handler
	generalHandler.NewGeneralHandler(router, gu)

	ar := authRepository.NewMysqlAuthRepository(connMysql)
	au := authUsecase.NewAuthUsecase(ar, gu, timeoutContext)
	authHandler.NewAuthHandler(router, au, gu)
	

	sr := storageRepository.NewMysqlStorageRepository(connMysql)
	su := storageUsecase.NewStorageUsecase(sr, gu, timeoutContext)
	storageHandler.NewStorageHandler(router, su, gu)

	fmt.Println("router run")
	fmt.Println(appPort)
	router.Run(appPort)
	fmt.Println("router exit")
}