package topic

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
	topic *entities.Topic,
) {
	getCurrentClientDB().First(&topic, "topic_id=?", topic.TopicID)
}

func (repo *Repository) GetItemByName(
	topic *entities.Topic,
) {
	getCurrentClientDB().First(&topic, "name=?", topic.Name)
}

func (repo *Repository) CreateItem(
	topic *entities.Topic,
) error {

	dbCreationResult := getCurrentClientDB().Create(&topic)

	if dbCreationResult.Error != nil {
		return dbCreationResult.Error
	}

	return nil
}

func (repo *Repository) GetAllItems(
	notes *[]entities.Topic,
) error {
	findAllErr := getCurrentClientDB().Find(&notes)

	if findAllErr.Error != nil {
		return findAllErr.Error
	}

	return nil
}

func (repo *Repository) DeleteItem(
	topic *entities.Topic,
) {
	getCurrentClientDB().Delete(&topic)
}

func (repo *Repository) UpdateItem(
	topic *entities.Topic,
) {
	getCurrentClientDB().Save(&topic)
}
