package domain

import "context"

type CrudViewSet interface {
	GetCrudService() CrudService
	Create(c context.Context) error
	Read(c context.Context) error
	Find(c context.Context) error
	Update(c context.Context) error
	Delete(c context.Context) error
}