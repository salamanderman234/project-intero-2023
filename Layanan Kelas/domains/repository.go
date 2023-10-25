package domain

import "context"

type Repository interface {
	GetConnection()
	Create(ctx context.Context, data Model) (Model, error)
	Update(ctx context.Context, id uint, updateFields Model) (int, Model,error)
	Read(ctx context.Context, query RepositoryGetQueryFunc) ([]Model, error)
	Find(ctx context.Context, id uint, model Model) (Model, error)
	Delete(ctx context.Context, id uint, model Model) (int, error)
}

type RepositoryGetQueryFunc func(ctx context.Context, conn uint, query Model) ([]Model, error)