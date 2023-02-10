package subject

import (
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"gorm.io/gorm"
)

type Repository struct {
	ClientDB *gorm.DB
}

var repository *Repository

func getCurrentClientDB() *gorm.DB {
	return *&repository.ClientDB
}

func GetRepository(clientDB *gorm.DB) *Repository {
	if repository == nil {
		repository = &Repository{
			ClientDB: clientDB,
		}
	}
	return repository
}

func (repo *Repository) GetItemById(
	subject *entities.Subject,
) {
	getCurrentClientDB().First(&subject, "subject_id=?", subject.SubjectID)
}

func (repo *Repository) GetItemByName(
	subject *entities.Subject,
) {
	getCurrentClientDB().First(&subject, "name=?", subject.Name)
}

func (repo *Repository) CreateItem(
	subject *entities.Subject,
) error {
	dbCreationResult := repo.ClientDB.Create(&subject)

	if dbCreationResult.Error != nil {
		return dbCreationResult.Error
	}

	return nil
}

func (repo *Repository) GetAllItems(
	subjects *[]entities.Subject,
) error {
	findAllErr := getCurrentClientDB().Find(&subjects)

	if findAllErr.Error != nil {
		return findAllErr.Error
	}

	return nil
}

func (repo *Repository) DeleteItem(
	subject *entities.Subject,
) {
	getCurrentClientDB().Delete(&subject)
}

func (repo *Repository) UpdateItem(
	subject *entities.Subject,
) {
	getCurrentClientDB().Save(&subject)
}
