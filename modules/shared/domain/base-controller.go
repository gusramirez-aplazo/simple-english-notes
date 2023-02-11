package domain

type BaseController interface {
	getAll() func(t *any) error
	getOneById() func(t *any) error
	getOneByUniqueParam() func(t *any) error
	createOne() func(t *any) error
	updateOne() func(t *any) error
	deleteOne() func(t *any) error
}
