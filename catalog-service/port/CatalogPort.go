package port

import "catolog-service/domain"

type CatalogPort interface {
	FindAll() ([]domain.Catalog, error)
	FindById(id int64) (domain.Catalog, error)
	Save(catalog domain.Catalog) (domain.Catalog, error)
	Delete(id int64) error
	Update(catalog domain.Catalog) (domain.Catalog, error)
}
