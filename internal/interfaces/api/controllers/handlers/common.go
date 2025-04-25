package handlers

import "localgems/internal/app/usecases"

type CoffeeHandler struct {
	coffeeUsecase usecases.CoffeeUsecase
}

func NewCoffeeHandler(usecase usecases.CoffeeUsecase) *CoffeeHandler {
	return &CoffeeHandler{
		coffeeUsecase: usecase,
	}
}
