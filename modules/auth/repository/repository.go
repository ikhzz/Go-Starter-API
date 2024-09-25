package repository

import (
	"starterapi/modules/auth/models"

	"gorm.io/gorm"
)

type authRepository struct {
	Conn *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) models.AuthRepository {
	return &authRepository{Conn: conn}
}

func (a authRepository) FindUserByEmail(param string) (res models.UserData, err error) {
	ecx := a.Conn.Table("users").Where("email = ? and deleted_at IS NULL", param).First(&res)
	if ecx.Error != nil {
		return res, ecx.Error
	}

	return res, err
}

func (a authRepository) FindUserByUid(param string) (res models.UserData, err error) {
	ecx := a.Conn.Table("users").Where("uid_user = ? and deleted_at IS NULL", param).First(&res)
	if ecx.Error != nil {
		return res, ecx.Error
	}

	return res, err
}

func (r *authRepository) PutUser(param string, toUpdate map[string]interface{}) error {
	query := r.Conn.Table("users").Where("uid_user = ? AND deleted_at IS NULL", param).Updates(&toUpdate)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
