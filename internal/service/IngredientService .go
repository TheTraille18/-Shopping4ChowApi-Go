package service

import (
	"shopping4chow/internal/models"
)

type IngredientService interface {
	GetIngredient(findIngredients models.Ingredient) []models.Ingredient
	RemoveIngredient(ingredient models.Ingredient)
	GetAllIngredients() []models.Ingredient
	AddIngredient(ingredient models.Ingredient)
}
