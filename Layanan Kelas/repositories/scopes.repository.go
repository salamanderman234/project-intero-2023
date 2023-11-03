package repository

import "gorm.io/gorm"

const (
	DATA_PERPAGE = 10
)

func PaginateScope(page uint) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		offset := DATA_PERPAGE*(page - 1)
		return db.Offset(int(offset)).Limit(DATA_PERPAGE)
	}
}