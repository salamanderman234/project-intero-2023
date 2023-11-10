package repository

import (
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

const (
	DATA_PERPAGE = 10
)

func paginateScope(page uint) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		offset := DATA_PERPAGE*(page - 1)
		return db.Offset(int(offset)).Limit(DATA_PERPAGE)
	}
}

var (
	orderWith = struct{
		DESC string
		ASC string
	}{
		DESC: "DESC",
		ASC: "ASC",
	}
)

func orderScope(model any, by string, sortWith string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
		if err != nil {
			return db
		}
		column := clause.Column{Name: s.Table + ".updated_at"}
		sortStatement := clause.OrderByColumn{Column: column, Desc: true, }
		if by != ""{
			valid := false
			for _, field := range s.Fields {
				if by == field.Name {
					valid = true
					break
				}
			}
			if valid {
				column.Name = s.Table + "." + by
			}
		}
		if sortWith == orderWith.DESC || sortWith == orderWith.ASC {
			sortStatement.Desc = sortWith == orderWith.DESC
		}
		return db.Order(sortStatement)
	}
}