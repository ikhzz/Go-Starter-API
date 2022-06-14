package models

import "time"

type StorageModel struct {
	Id int `json:"product_id" gorm:"primaryKey"`
	ProductName string `validate:"required" json:"product_name" gorm:"notnull;type:varchar(255)"`
	ProductDescription string `json:"product_description" gorm:"type:text"`
	ProductOwner string `json:"product_owner" gorm:"type:varchar(255)"`
	ProductStock int `validate:"required" json:"product_stock" gorm:"notnull;type:int(11)"`
	ProductPrice int `validate:"required" json:"product_price" gorm:"notnull;type:int"`
	CreatedAt time.Time `json:"-" `
	UpdatedAt time.Time `json:"-" `
	IsDeleted int `json:"-" gorm:"notnull;type:int(1);default:0"`
	ProductForeign AuthModel `validate:"-" json:"-" gorm:"references:Username;foreignKey:ProductOwner;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (StorageModel) TableName() string {
	return "Storage"
}

type StorageGetParam struct {
	Keyword string `form:"keyword"`
	Orderby string `form:"order_by"`
}

type StorageAddParam struct {
	ProductName string `validate:"required" json:"product_name"`
	ProductStock int `validate:"required" json:"product_stock"`
	ProductPrice int `validate:"required" json:"product_price"`
}