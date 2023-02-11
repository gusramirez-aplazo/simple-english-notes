package note

import (
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
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

func (repo *Repository) GetItem(
	note *entities.Note,
) {
	getCurrentClientDB().First(&note)
}

func (repo *Repository) CreateItem(
	note *entities.Note,
) error {

	dbCreationResult := getCurrentClientDB().Create(&note)

	if dbCreationResult.Error != nil {
		return dbCreationResult.Error
	}

	return nil
}

func (repo *Repository) GetAllItems(
	notes *[]entities.Note,
) error {
	findAllErr := getCurrentClientDB().Find(&notes)

	if findAllErr.Error != nil {
		return findAllErr.Error
	}

	return nil
}

func (repo *Repository) DeleteItem(
	note *entities.Note,
) {
	getCurrentClientDB().Delete(&note)
}

func (repo *Repository) UpdateItem(
	note *entities.Note,
) {
	getCurrentClientDB().Save(&note)
}
