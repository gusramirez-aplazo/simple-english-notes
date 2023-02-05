package interfaces

import "gorm.io/gorm"

type Repository interface {
	getCurrentClientDB() *gorm.DB
	GetCurrentRepository(clientDB *gorm.DB) *Repository
	GetItem()
	GetItemOrCreate()
	GetAllItems()
	DeleteItem()
	UpdateItem()
}

type Repositories interface {
}
