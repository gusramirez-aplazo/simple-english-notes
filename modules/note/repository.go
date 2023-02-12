package note

import (
	"errors"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"gorm.io/gorm"
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
) (domain.Note, error) {
	item := &domain.Note{}

	getCurrentClientDB().
		First(item,
			"id",
			id,
		)

	if item.Content == "" {
		return *item, errors.New("item not found")
	}

	return *item, nil
}

func (repo *Repository) CreateOne(
	content string,
	cue string,
) (domain.Note, error) {
	item := &domain.Note{
		Cue:     cue,
		Content: content,
	}

	query := getCurrentClientDB().
		Create(item)

	if query.Error != nil {
		return *item, query.Error
	}

	return *item, nil
}

func (repo *Repository) GetAllItems() ([]domain.Note, error) {
	var items []domain.Note

	query := getCurrentClientDB().
		Find(&items)

	if query.Error != nil {
		return items, query.Error
	}

	return items, nil
}

func (repo *Repository) DeleteOne(
	id uint,
) (domain.Note, error) {
	item, err := repo.GetItemById(id)

	if err != nil {
		return item, err
	}

	query := getCurrentClientDB().
		Delete(&item)

	if query.Error != nil {
		return item, query.Error
	}

	return item, nil
}

func (repo *Repository) UpdateOne(
	id uint,
	content string,
	cue string,
) (domain.Note, error) {
	item, findErr := repo.
		GetItemById(id)

	itemNotFoundedErr := findErr != nil

	if itemNotFoundedErr {
		return item, findErr
	}

	if item.Content == content && item.Cue == cue {
		return item, errors.New("nothing to update")
	}

	item.Content = content
	item.Cue = cue
	getCurrentClientDB().
		Save(&item)

	return item, nil
}
