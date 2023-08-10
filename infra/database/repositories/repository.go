package repositories

type Entity interface {
	GetId() string
}

type Repository[K, E Entity] interface {
	Save(K) (E, error)
	Update(K) (E, error)
	Delete(K) error
	ListAll() ([]E, error)
	ListById(string) (E, error)
}
