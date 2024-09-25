package config

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	cmmod "starterapi/common/models"
	aumod "starterapi/modules/auth/models"
	nomod "starterapi/modules/notes/models"
)

func InitMysql() *gorm.DB {
	dbHost := viper.GetString(`database.mysql.host`)
	dbPort := viper.GetString(`database.mysql.port`)
	dbUser := viper.GetString(`database.mysql.user`)
	dbPass := viper.GetString(`database.mysql.pass`)
	dbName := viper.GetString(`database.mysql.db`)

	basedsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", viper.GetString("key.timezone"))
	dsn := fmt.Sprintf("%s?%s", basedsn, val.Encode())

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	conn.AutoMigrate(&cmmod.LogModel{})
	conn.AutoMigrate(&aumod.UsersModel{})
	conn.AutoMigrate(&nomod.NotesModel{})
	var count int64
	conn.Table("users").Count(&count)
	if count == 0 {
		passbyte, _ := bcrypt.GenerateFromPassword([]byte(viper.GetString("key.admin_cred.password")), bcrypt.DefaultCost)

		superAdmin := aumod.UsersModel{
			Username: viper.GetString("key.admin_cred.username"),
			Password: string(passbyte),
			UIDUser:  uuid.NewV4(),
			Email:    viper.GetString("key.admin_cred.email"),
		}
		conn.Table("users").Create(&superAdmin)
	}

	return conn
}
