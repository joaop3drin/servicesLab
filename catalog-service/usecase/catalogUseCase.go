package usecase

import (
	"catolog-service/domain"
	"catolog-service/port"
)

type CatalogUseCase struct {
	CatalogPort port.CatalogPort
}

func NewCatalogUseCase(catalogPort port.CatalogPort) *CatalogUseCase {
	return &CatalogUseCase{
		CatalogPort: catalogPort,
	}
}

func (c *CatalogUseCase) FindAll() ([]domain.Catalog, error) {
	return c.CatalogPort.FindAll()
}

func (c *CatalogUseCase) FindById(id int64) (domain.Catalog, error) {
	return c.CatalogPort.FindById(id)
}

func (c *CatalogUseCase) Save(catalog domain.Catalog) (domain.Catalog, error) {
	return c.CatalogPort.Save(catalog)
}

func (c *CatalogUseCase) Delete(id int64) error {
	return c.CatalogPort.Delete(id)
}

func (c *CatalogUseCase) Update(catalog domain.Catalog) (domain.Catalog, error) {
	return c.CatalogPort.Update(catalog)
}
