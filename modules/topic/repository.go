package topic

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
	return *&repository.ClientDB
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
) (domain.Topic, error) {
	item := &domain.Topic{}

	getCurrentClientDB().
		First(item,
			"id",
			id,
		)

	if item.Name == "" {
		return *item, errors.New("item not found")
	}

	return *item, nil
}

func (repo *Repository) GetItemByUniqueParam(
	name string,
) (domain.Topic, error) {
	item := &domain.Topic{}

	getCurrentClientDB().
		First(
			item,
			"name=?",
			name,
		)

	if item.ID == 0 {
		return *item, errors.New("item not found")
	}

	return *item, nil
}

func (repo *Repository) CreateOne(
	name string,
) (domain.Topic, error) {
	item := &domain.Topic{
		Name: name,
	}

	query := getCurrentClientDB().
		Create(item)

	if query.Error != nil {
		return *item, query.Error
	}

	return *item, nil
}

func (repo *Repository) GetAllItems() ([]domain.Topic, error) {
	var items []domain.Topic

	query := getCurrentClientDB().
		Find(&items)

	if query.Error != nil {
		return items, query.Error
	}

	return items, nil
}

func (repo *Repository) DeleteOne(
	id uint,
) (domain.Topic, error) {
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
	name string,
) (domain.Topic, error) {
	item, findErr := repo.
		GetItemById(id)

	itemNotFoundedErr := findErr != nil

	if itemNotFoundedErr {
		return item, findErr
	}

	if item.Name == name {
		return item, errors.New("nothing to update")
	}

	_, findByNameErr := repo.
		GetItemByUniqueParam(name)

	isItemFoundByName := findByNameErr == nil

	if isItemFoundByName {
		return item, errors.New("the name you intend to update is already taken")
	}

	item.Name = name
	getCurrentClientDB().
		Save(&item)

	return item, nil
}
