package repository

import (
	"context"
	"math"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
	"gorm.io/gorm"
)

type classSubjectRepository struct {
	db *gorm.DB
}

func NewClassSubjectRepository(db *gorm.DB) domain.ClassSubjectRepository {
	return &classSubjectRepository{
		db: db,
	}
}

func (cs *classSubjectRepository) Create(c context.Context, data domain.ClassSubjectModel) (domain.ClassSubjectModel, error) {
	_, err := basicCreateFunc(c, cs.db, &data)
	return data, err
}

func (cs *classSubjectRepository) Read(c context.Context, id uint, studentID uint, teacherID uint, classID uint, year uint, page uint, orderBy string, orderWith string) ([]domain.ClassSubjectModel, uint, error) {
	var results []domain.ClassSubjectModel
	var maxPage int64
	queryDB := cs.db.Scopes(orderScope(domain.ClassSubjectModel{}, orderBy, orderWith)).Model(&domain.ClassSubjectModel{}).WithContext(c)
	if year != 0 && studentID == 0{
		queryDB = queryDB.Where("year = ?", year)
	}
	if id != 0{
		queryDB = queryDB.Where("id = ?", id)
	}else if studentID != 0 {
		var studentClasses []domain.StudentClassModel
		cs.db.Model(&domain.StudentClassModel{}).
			WithContext(c).
			Where("student_id = ?", studentID).
			Find(&studentClasses)
		if len(studentClasses) <= 0 {
			return results, 0, domain.ErrResourceNotFound
		}
		for _, studentClass := range studentClasses {
			queryDB = queryDB.Or(cs.db.Where("class_id = ?", studentClass.ClassID).Where("year = ?", studentClass.Year))
		}

	} else if teacherID != 0{
		queryDB = queryDB.Where("teacher_id = ?", teacherID)
	} else if classID != 0 {
		queryDB = queryDB.Where("class_id = ?", classID)
	}
	_ = *queryDB.Count(&maxPage)
	result := queryDB.Scopes(paginateScope(page)).Find(&results)
	if result.RowsAffected <= 0 {
		return results,0, domain.ErrResourceNotFound
	}
	maxPage = int64(math.Ceil(float64(maxPage)/DATA_PERPAGE))
	return results, uint(maxPage), handleRepositoryError(result)
}
func (cs *classSubjectRepository) Update(c context.Context, id uint, data domain.ClassSubjectModel) (int64, domain.ClassSubjectModel, error) {
	result, err := basicUpdateFunc(c, cs.db, id, &data)
	return result.RowsAffected, data, err
}
func (cs *classSubjectRepository) Delete(c context.Context, id uint) (int64, error) {
	result, err := basicDeleteFunc(c, cs.db, id, &domain.ClassSubjectModel{})
	return result.RowsAffected, err
}