package models

import "starterapi/common/models"

type (
	SwaggoGetNotesResSuccess struct {
		Status         bool        `json:"status" example:"true"`
		StatusCode     int         `json:"status_code" example:"200"`
		Message        string      `json:"message" example:"success"`
		TotalGroupData int         `json:"total_data" example:"0"`
		Data           []NotesData `json:"data"`
	}

	SwaggoGetNotesResBadRequest struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"valid param required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoGetNotesResUnauthorized struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"401"`
		Message        string               `json:"message" example:"token is required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoGetNotesResNotFound struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"data not found"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoPostNotesResSuccess struct {
		Status         bool                 `json:"status" example:"true"`
		StatusCode     int                  `json:"status_code" example:"200"`
		Message        string               `json:"message" example:"success"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoPostNotesResBadRequest struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"valid param required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoPostNotesResUnauthorized struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"401"`
		Message        string               `json:"message" example:"token is required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoPutNotesResSuccess struct {
		Status         bool                 `json:"status" example:"true"`
		StatusCode     int                  `json:"status_code" example:"200"`
		Message        string               `json:"message" example:"success"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoPutNotesResBadRequest struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"valid param required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoPutNotesResUnauthorized struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"401"`
		Message        string               `json:"message" example:"token is required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoDeleteNotesResSuccess struct {
		Status         bool                 `json:"status" example:"true"`
		StatusCode     int                  `json:"status_code" example:"200"`
		Message        string               `json:"message" example:"success"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoDeleteNotesResBadRequest struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"400"`
		Message        string               `json:"message" example:"valid param required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}

	SwaggoDeleteNotesResUnauthorized struct {
		Status         bool                 `json:"status" example:"false"`
		StatusCode     int                  `json:"status_code" example:"401"`
		Message        string               `json:"message" example:"token is required"`
		TotalGroupData int                  `json:"total_data" example:"0"`
		Data           models.EmptyResponse `json:"data"`
	}
)
