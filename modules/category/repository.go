package category

import (
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	ClientDB *gorm.DB
}

var repository *Repository

func getCurrentClientDB() *gorm.DB {
	return *&repository.ClientDB
}

func GetCategoryRepository(
	clientDB *gorm.DB,
) *Repository {
	if repository == nil {
		repository = &Repository{
			ClientDB: clientDB,
		}
	}
	return repository
}

func (repo *Repository) AsyncGet(
	categories chan<- *entities.Category,
	category *entities.Category,
) {

	transaction := getCurrentClientDB().Begin()
	transaction.FirstOrCreate(category)
	transaction.Commit()

	log.Println(category)

	categories <- category
}

func (repo *Repository) GetItem(
	category *entities.Category,
) {
	getCurrentClientDB().First(&category, "name=?", category.Name)
}

func (repo *Repository) GetItemOrCreate(
	category *entities.Category,
) error {
	repo.GetItem(category)

	if category.CategoryID != 0 {
		return nil
	}

	dbCreationResult := getCurrentClientDB().Create(&category)

	if dbCreationResult.Error != nil {
		return dbCreationResult.Error
	}

	return nil
}

func (repo *Repository) GetAllItems(
	categories *[]entities.Category,
) error {
	findAllErr := getCurrentClientDB().Find(&categories)

	if findAllErr.Error != nil {
		return findAllErr.Error
	}

	return nil
}

func (repo *Repository) DeleteItem(
	category *entities.Category,
) {
	getCurrentClientDB().Delete(&category)
}

func (repo *Repository) UpdateItem(
	category *entities.Category,
) {
	getCurrentClientDB().Save(&category)
}
