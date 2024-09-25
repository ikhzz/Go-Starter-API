## Project Name
# - golang startup prokect

## Description
-

## Version Running / Configuration
1. Minimum go version 1.16
2. Copy the config.json.sample file to config.json.
3. Open config.json and modify the necessary settings.

## Env structure
```
{
    "mode": "development",
    "address": ":6060",
    "database": {
        "mysql": {
            "host": "MYSQL_HOST",
            "port": "MYSQL_PORT",
            "user": "MYSQL_USER",
            "pass": "MYSQL_PASS",
            "db": "MYSQL_DBNAME",
            "parseTime": "1"
        },
        "mongo": {
            "host": "MYSQL_HOST",
            "port": "MYSQL_PORT",
            "user": "MYSQL_USER",
            "pass": "MYSQL_PASS",
            "db": "MYSQL_DBNAME",
            "dbauth": "MYSQL_DBAUTH",
            "url": "MSYQL_URL"
        }
    },
    "smtp": {
        "host": "SMTP_HOST",
        "port": "SMTP_PORT",
        "auth": "SMTP_AUTH",
        "name": "SMTP_SENDER_NAME",
        "address": "SMTP_ADDRESS",
        "username": "SMTP_USERNAME",
        "password": "SMTP_PASSWORD",
        "secure": "SMTP_SECURE"
    },
    "timeout": {
        "context": 10,
        "forget_password_token": 180000,
        "jwt":2592000
    },
    "key": {
        "jwt": "JWT_KEY",
        "admin_cred": {
            "email": "ADMIN_EMAIL",
            "username": "ADMIN_USERNAME",
            "password": "ADMIN_PASSWORD
        },
        "default_lang": "en",
        "timezone": "Asia/Jakarta"
    },
    "default_unhandled_error": "Unhandled error. Please try again.",
    "cloudinary": {
        "user":"CLOUDINARY_USER",
        "key":"CLOUDINARY_KEY",
        "secret":"CLOUDINARY_SECRET",
        "projectName": "CLOUDINARY_NAME"
    }
}
```
## Directory structure
```
.
├── config // database directory or inital data
│   └──db.go
│   └──seeder_file
├── docs // swaggo directory for easy documentation
├── common // helper/middleware module directory
├── module // module directory
│   └──────[module_name]
│          ├──delivery
│          │  └── [delivery name].go
│          ├──repository
│          │  └── [repository name].go
|          ├──model
│          │  └── [model name].go
│          └──usecase
│             └── [usecase name].go
├── config.json // json config
├── routes // init router
│   └── routes.go
├── .gitignore
├── config.json
├── go.mod
├── go.sum
└── main.go
```

## Package
- framework: gin(github.com/gin-gonic/gin)
- orm: 
  - gorm(gorm.io/gorm)
  - (gorm.io/driver/mysql)
- validator: validator(github.com/go-playground/validator/v10)
- documentation: 
  - swaggo(github.com/swaggo/swag/cmd/swag)
  - (github.com/swaggo/files)
  - (github.com/swaggo/gin-swagger)
- uuid helper: satori(github.com/satori/go.uuid)
- jwt helper: dgrijalva(github.com/dgrijalva/jwt-go)
- env setup: viper(github.com/spf13/viper)
- unit test: stretchr(github.com/stretchr/testify)
- hash helper: (golang.org/x/crypto)
- hot reload: air-verse(github.com/air-verse/air)

  

