package models

import (
	"mime/multipart"
)

type TokenRequirement struct {
	Id string
	Role string
}

type AuthResponse struct {
	Field string `json:"field"`
	Msg string `json:"message"`
}

type FileModel struct {
	Category string `form:"category"`
	File *multipart.FileHeader `form:"file"`	
}

type FileResponse struct {
	Filename string `json:"file_name"`
	Fileurl string `json:"file_url"`
}