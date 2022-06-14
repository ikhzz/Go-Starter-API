package models

import (
	"time"
)

type AuthModel struct {
	Id  string `json:"omitempty" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email string `validate:"required,email" json:"email" gorm:"unique;notnull"`
	Username string `validate:"required" json:"username" gorm:"notnull;type:varchar(255)"`
	Password string `validate:"required,gte=6" json:"password" gorm:"notnull;size:255;type:varchar(255)"`
	ProfileImage string `json:"profile_image" gorm:"type:varchar(255)"`
	Role string `json:"-" gorm:"notnull;type:enum('seller', 'buyer')"`
	CreatedAt time.Time `json:"-" `
	IsActive string `json:"-" gorm:"notnull;default:1;type:tinyint(1)"`
}

type Tabler interface {
	TableName() string
}

func (AuthModel) TableName() string {
	return "User"
}

type AuthSignIn struct {
	Email string `validate:"required,email" json:"email"`
	Password string `validate:"required,gte=6" json:"password" `
}

type AuthUpdateProfile struct {
	Id string `json:"-"`
	Username string `validate:"required,gte=3" json:"username"`
	Email string `validate:"required,email" json:"email"`
	Image string `validate:"required" json:"image"`
}

type AuthUpdatePassword struct {
	OldPassword string `validate:"required,gte=6" json:"old_password"`
	NewPassword string `validate:"required,gte=6" json:"new_password"`
	ConfirmPassword string `validate:"required,gte=6" json:"confirm_password"`
}

type AuthResult struct {
	Id string `json:"-" `
	Email string `json:"-" `
	Username string `json:"-" `
	Password string `json:"-" `
	Role string `json:"-"`
}

type AuthProfileResult struct {
	Id string `json:"-"`
	Email string `json:"email"`
	Username string `json:"username"`
	ProfileImage string `json:"profile_image" `
}