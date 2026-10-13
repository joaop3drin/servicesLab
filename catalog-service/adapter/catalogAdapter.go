package adapter

import (
	"catolog-service/domain"
	"fmt"
)

type CatalogAdapter struct {
}

func (a *CatalogAdapter) FindAll() ([]domain.Catalog, error) {
	fmt.Print("Find All Called")
	return nil, nil
}

func (a *CatalogAdapter) FindById(id int64) (domain.Catalog, error) {
	fmt.Print("Find By Id Called")
	return domain.Catalog{}, nil
}

func (a *CatalogAdapter) Save(catalog domain.Catalog) (domain.Catalog, error) {
	fmt.Print("Save Called")
	return domain.Catalog{}, nil
}

func (a *CatalogAdapter) Delete(id int64) error {
	fmt.Print("Delete Called")
	return nil
}

func (a *CatalogAdapter) Update(catalog domain.Catalog) (domain.Catalog, error) {
	fmt.Print("Update Called")
	return domain.Catalog{}, nil
}
