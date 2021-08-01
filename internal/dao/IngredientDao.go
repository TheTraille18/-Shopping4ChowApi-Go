package dao

import (
	"shopping4chow/internal/models"
)

type IngredientDao interface {
	GetIngredient(findIngredients models.Ingredient) []models.Ingredient
	RemoveIngredient(ingredient models.Ingredient)
	GetAllIngredients() []models.Ingredient
	AddIngredient(ingredient models.Ingredient)
}
