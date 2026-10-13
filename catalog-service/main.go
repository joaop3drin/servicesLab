package main

import (
	"catolog-service/adapter"
	"catolog-service/usecase"
	"fmt"
)

func main() {

	adapter := &adapter.CatalogAdapter{}
	useCase := usecase.NewCatalogUseCase(adapter)

	fmt.Println("Find All Called by main.go")
	useCase.FindAll()
}
