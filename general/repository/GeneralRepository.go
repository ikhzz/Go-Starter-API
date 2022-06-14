package repository

import (
	"context"
)

func (g *mysqlGeneralRepository) GetUsername(ctx context.Context,id string) (name string) {
	g.ConnSql.Table("user").Where("id = ?", id).Where("is_active = 1").Select("username").Find(&name)
	return
}