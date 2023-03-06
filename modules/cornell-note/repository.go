package cornellNote

import (
	"errors"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	ClientDB *gorm.DB
}

var repository *Repository

func getCurrentClientDB() *gorm.DB {
	return repository.ClientDB
}

func GetRepository(
	clientDB *gorm.DB,
) *Repository {
	if repository == nil {
		repository = &Repository{
			ClientDB: clientDB,
		}
	}
	return repository
}

func (repo *Repository) GetItemById(
	id uint,
) (domain.CornellNote, error) {
	resp := &domain.CornellNote{}

	getCurrentClientDB().
		Preload(clause.Associations).
		First(
			resp,
			"id=?",
			id,
		)

	if resp.Topic.Name == "" {
		return *resp, errors.New("item not found")
	}

	return *resp, nil
}

func (repo *Repository) CreateOne(
	topic domain.Topic,
	subjects []domain.Subject,
	notes []domain.Note,
) (domain.CornellNote, error) {
	var cornellNote = new(domain.CornellNote)
	cornellNote.Topic = topic
	cornellNote.Subjects = subjects
	cornellNote.Notes = notes

	dbCreationResponse := getCurrentClientDB().
		Create(cornellNote)

	if dbCreationResponse.Error != nil {
		return *cornellNote, dbCreationResponse.Error
	}

	getCurrentClientDB().Save(cornellNote)

	return *cornellNote, nil
}

func (repo *Repository) GetAllItems() ([]domain.CornellNote, error) {
	var items []domain.CornellNote

	query := getCurrentClientDB().
		Preload(clause.Associations).
		Find(&items)

	if query.Error != nil {
		return items, query.Error
	}

	return items, nil
}
