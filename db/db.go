package db

import  (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/spf13/viper"
	"fmt"
	"context"
	"time"
	"net/url"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/sqlite"
	"cleanArchCrud/models"
)

func InitMysql() *gorm.DB {
	// get database variale
	dbHost := viper.GetString(`database.mysql.host`)
	dbPort := viper.GetString(`database.mysql.port`)
	dbUser := viper.GetString(`database.mysql.user`)
	dbPass := viper.GetString(`database.mysql.pass`)
	dbName := viper.GetString(`database.mysql.db`)

	// set dsn url and try connection
	basedsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", viper.GetString(`database.mysql.parseTime`))
	val.Add("loc", viper.GetString("timezone"))
	dsn := fmt.Sprintf("%s?%s", basedsn, val.Encode())
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// advance setting
	// sqldb,err :=db.DB()
	// sqldb.SetMaxOpenConns(10)
	// sqldb.SetMaxIdleConns(10)
	// sqldb.SetConnMaxLifetime(10)
	conn.AutoMigrate(&models.AuthModel{}, &models.StorageModel{})
	return conn
}

func InitMonggoDB() *mongo.Database {
	// get databse url
	user := viper.GetString("database.mongo.user")
	pass := viper.GetString("database.mongo.pass")
	host := viper.GetString("database.mongo.host")
	port := viper.GetString("database.mongo.port")
	dbAuth := viper.GetString("database.mongo.dbauth")
	defaultUrl := viper.GetString("database.mongo.url")

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// set url and try connection
	url := "mongodb://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbAuth +"?ssl=false"
	fmt.Println(url)
	client, err := mongo.NewClient(options.Client().ApplyURI(defaultUrl))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	errConnect := client.Connect(ctx)
	if errConnect != nil {
		panic(errConnect)
	}
	return client.Database("cleanArchOne")
}

func InitSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("routes.db"), &gorm.Config{})
	// db, err := sql.Open("sqlite3", "file:routes.db?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	return db
}