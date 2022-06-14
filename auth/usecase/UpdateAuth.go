package usecase

import (
	"cleanArchCrud/models"
	"context"
	"errors"
	"fmt"
)


func(a *authUsecase) UpdateProfile(ctx context.Context, data *models.AuthUpdateProfile) (error) {
	err := a.Authrepos.UpdateProfile(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func(a *authUsecase) UpdatePassword(ctx context.Context, data *models.AuthUpdatePassword) (error) {
	if data.NewPassword != data.ConfirmPassword {
		return errors.New("new password and confirmation password is not equal")
	}
	getPassword, errGetPassword := a.Authrepos.GetPasswordById(ctx)
	if errGetPassword != nil {
		return errGetPassword
	}
	fmt.Println(getPassword)
	errCompare := a.GeneralUsecase.PasswordCompare(getPassword, data.OldPassword)
	if errCompare != nil {
		return errors.New("wrong old password")
	}
	
	data.NewPassword, errCompare = a.GeneralUsecase.PasswordHash(data.NewPassword)
	if errCompare != nil {
		return errors.New("failed to create password")
	}
	errCompare = a.Authrepos.UpdatePassword(ctx, data.NewPassword) 
	if errCompare != nil {
		return errCompare
	}
	return nil
}