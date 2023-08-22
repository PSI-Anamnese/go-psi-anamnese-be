package repositories

type Entity interface {
	GetId() string
}

type Repository[K string, E Entity] interface {
	Save(E) (E, error)
	Update(E) (E, error)
	Delete(K) error
	ListAll() ([]E, error)
	ListById(string) (E, error)
}
