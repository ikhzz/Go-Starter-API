package repository

import (
	"log"
	"math"
	"starterapi/modules/notes/models"

	"gorm.io/gorm"
)

type notesRepository struct {
	Conn *gorm.DB
}

func NewNotesRepository(conn *gorm.DB) models.NotesRepository {
	return &notesRepository{Conn: conn}
}

func (nr notesRepository) GetNotes(param *models.ReqGetNotes, id *string) ([]models.NotesData, int64, error) {
	result := make([]models.NotesData, 0)
	var count int64 = 1

	query := nr.Conn.Table("notes").Where("deleted_at IS NULL AND created_by = ?", id)

	if param.Keyword != "" {
		param.Keyword = "%" + param.Keyword + "%"
		query = query.Where("title LIKE ? OR description LIKE ? ", param.Keyword, param.Keyword)
	}

	if param.All == 0 {
		query.Count(&count)
		if param.Limit < 1 {
			param.Limit = 10
		}
		if param.Offset > 1 {
			query = query.Offset((param.Offset - 1) * param.Limit)
		}
		query = query.Limit(param.Limit)
		log.Println(count)

		count = int64(math.Ceil(float64(count) / float64(param.Limit)))

	}

	query = query.Find(&result)
	if query.Error != nil {
		return result, count, query.Error
	}

	return result, count, nil
}

func (r *notesRepository) PostNotes(param *models.ReqPostNotes) error {
	query := r.Conn.Table("notes").Create(param)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *notesRepository) PutNotes(param string, toUpdate map[string]interface{}) error {
	query := r.Conn.Table("notes").Where("uid_notes = ?", param).Updates(&toUpdate)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
