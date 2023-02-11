package domain

type BaseRepository interface {
	CreateOne(name string) error
	GetAllItems() error
	GetItemById(id uint)
	GetItemByUniqueParam(name string)
	UpdateOne(id uint)
	DeleteOne(id uint)
}
