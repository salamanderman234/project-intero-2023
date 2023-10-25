package domain

import (
	"context"

	"gorm.io/gorm"
)

type Form interface {
	ConvertToEntity() Entity
}

type Model interface {
	GetID() uint
	GetPreloadStatement() QueryFunc
	GetAdditionalStatement() QueryFunc
}

type Entity interface {
	GetID() uint
	SetIDToNull()
	GetModel() Model
}

type QueryFunc func(ctx context.Context, query *gorm.DB) *gorm.DB
type AfterQueryFunc func(ctx context.Context, object Model, conn *gorm.DB) (error)