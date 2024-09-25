package models

import "starterapi/common/models"

type (
	SwaggoSigninResSuccess struct {
		Status         bool      `json:"status" example:"true"`
		StatusCode     int       `json:"status_code" example:"200"`
		Message        string    `json:"message" example:"success"`
		TotalGroupData int       `json:"total_data" example:"0"`
		Data           ResSignin `json:"data"`
	}

	SwaggoSigninResBadRequest struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"valid param required, password missmatch"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoSigninResNotFound struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"404"`
		Message        string               `json:"message" example:"error get user"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoGetProfileResSuccess struct {
		Status         bool     `json:"status" example:"true"`
		StatusCode     int      `json:"status_code" example:"200"`
		Message        string   `json:"message" example:"success"`
		TotalGroupData int      `json:"total_data" example:"0"`
		Data           UserData `json:"data"`
	}

	SwaggoGetProfileResBadRequest struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"valid param required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoGetProfileResUnauthorized struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"401"`
		Message        string               `json:"message" example:"data not found, token is required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}
)
