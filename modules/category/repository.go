package category

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
) (domain.Category, error) {
	resp := &domain.Category{}

	getCurrentClientDB().
		First(
			resp,
			"id=?",
			id,
		)

	if resp.Name == "" {
		return *resp, errors.New("item not found")
	}

	return *resp, nil
}

func (repo *Repository) GetItemByUniqueParam(
	name string,
) (domain.Category, error) {
	resp := &domain.Category{}

	getCurrentClientDB().
		First(
			resp,
			"name=?",
			name,
		)

	if resp.ID == 0 {
		return *resp, errors.New("item not found")
	}

	return *resp, nil
}

func (repo *Repository) CreateOne(
	name string,
) (domain.Category, error) {
	resp := &domain.Category{
		Name: name,
	}

	query := getCurrentClientDB().
		Create(resp)

	if query.Error != nil {
		return *resp, query.Error
	}

	return *resp, nil
}

func (repo *Repository) GetAllItems() ([]domain.Category, error) {
	var items []domain.Category

	query := getCurrentClientDB().
		Find(&items)

	if query.Error != nil {
		return items, query.Error
	}

	return items, nil
}

func (repo *Repository) DeleteOne(
	id uint,
) (domain.Category, error) {
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
) (domain.Category, error) {
	item, findErr := repo.
		GetItemById(id)

	itemNotFoundByIdErr := findErr != nil

	if itemNotFoundByIdErr {
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
