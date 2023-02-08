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

func GetSubjectRepository(clientDB *gorm.DB) *Repository {
	if repository == nil {
		repository = &Repository{
			ClientDB: clientDB,
		}
	}
	return repository
}

func (repo *Repository) GetItem(
	subject *entities.Subject,
) {
	getCurrentClientDB().First(&subject, "name=?", subject.Name)
}

func (repo *Repository) GetItemOrCreate(
	subject *entities.Subject,
) error {
	repo.GetItem(subject)
	if subject.SubjectID != 0 {
		return nil
	}

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
