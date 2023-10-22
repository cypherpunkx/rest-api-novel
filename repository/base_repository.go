package repository

type BaseRepository[T any] interface {
	Create(payload T) error
	ListPaginate(page, perPage int, code, title, publisher, year, author string) ([]T, error)
	List() ([]T, error)
	Get(id string) (T, error)
	Update(id string, payload T) (*T, error)
	Delete(id string) error
	Count(code, title, publisher, year, author string) (int, error)
}
