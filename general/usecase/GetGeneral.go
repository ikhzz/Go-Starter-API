package usecase

import "context"


func (g *GeneralUsecase) GetUsername(ctx context.Context, id string) string {
	return g.Generalrepos.GetUsername(ctx, id)
}