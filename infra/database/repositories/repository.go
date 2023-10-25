package repositories

import "github.com/psi-anamnese/psi-anamnese-be/infra/database/repositories/entities"

type Repository[K string, E entities.Entity] interface {
	Save(E) (E, error)
	Update(E) (E, error)
	Delete(K) error
	ListAll() ([]E, error)
	ListById(string) (E, error)
}
