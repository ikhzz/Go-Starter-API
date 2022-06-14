package repository

import (
	"cleanArchCrud/models"
	"context"
	"errors"
	"fmt"
)

func (m *mysqlAuthRepository) Create(ctx context.Context, data *models.AuthModel) error {
	res := m.ConnSql.Create(data)
	if res.Error != nil {
		return errors.New("email is already registered")
	}
	return nil;
}

func (m *mysqlAuthRepository) SignIn(ctx context.Context, data *models.AuthSignIn) (models.AuthResult, error) {
	result := models.AuthResult{}
	res := m.ConnSql.Table("user").Where("is_active = 1").Select("id", "email", "password", "role").First(&result, fmt.Sprintf(`email = '%s'`, data.Email))
	if res.Error != nil {
		return result, errors.New("user not found")
	}
	
	return result,nil
}

func (m *mysqlAuthRepository) UpdateProfile(ctx context.Context, data *models.AuthUpdateProfile) error {
	value := map[string]interface{}{"email": data.Email, "username": data.Username, "profile_image": data.Image}
	res := m.ConnSql.Table("user").Where("id = ?", fmt.Sprintf(`%s`, ctx.Value("payload_id"))).Updates(value)
	if res.Error != nil {
		return errors.New("failed to save")
	}

	return nil
}

func (m *mysqlAuthRepository) GetPasswordById(ctx context.Context) (string, error) {
	var result string
	res := m.ConnSql.Table("user").Where("id = ?", fmt.Sprintf(`%s`, ctx.Value("payload_id"))).Select("password").Find(&result)
	if res.Error != nil {
		return result, errors.New("user not found")
	}
	return result,nil
}

func (m *mysqlAuthRepository) UpdatePassword(ctx context.Context, newpass string) error {
	res := m.ConnSql.Table("user").Where("id = ?", fmt.Sprintf(`%s`, ctx.Value("payload_id"))).Update("password", newpass)
	if res.Error != nil {
		return errors.New("failed to save password")
	}

	return nil
}