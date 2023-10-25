package domain

import "context"

type CrudService interface {
	GetRepository() Repository
	CreateService(ctx context.Context, data Entity) (Entity, error)
	ReadService(ctx context.Context, query Entity) ([]Entity, error)
	FindService(ctx context.Context, id uint, entity Entity) (Entity, error)
	UpdateService(ctx context.Context, id uint, updateFields Entity) (int, Entity, error)
	DeleteService(ctx context.Context, id uint, entity Entity) (int, error)
}