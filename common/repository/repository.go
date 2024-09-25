package repository

import (
	"gorm.io/gorm"
	"starterapi/common/models"
)

type commonRepository struct {
	ConnSql *gorm.DB
	// can add nosql db
}

func NewCommonRepository(Conn *gorm.DB) models.CommonRepository {

	return &commonRepository{Conn}
}

func (a commonRepository) CreateLog(param *models.LogModel) error {
	err := a.ConnSql.Table("log").Create(&param)

	return err.Error
}

func (a commonRepository) FindUserByUid(param string) (res models.UserData, err error) {
	ecx := a.ConnSql.Table("users").Where("uid_user = ? and deleted_at IS NULL", param).First(&res)
	if ecx.Error != nil {
		return res, ecx.Error
	}

	return res, err
}

func (r *commonRepository) PutUser(param string, toUpdate map[string]interface{}) error {
	query := r.ConnSql.Table("users").Where("uid_user = ? AND deleted_at IS NULL", param).Updates(&toUpdate)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
