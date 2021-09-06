package service

import "shopping4chow/cmd/shopping4chow/models"

type RecipeService interface {
	GetRecipe(recipe models.Recipe) []models.Recipe
	RemoveRecipe(recipe models.Recipe)
	GetAllRecipes() []models.Recipe
	AddRecipe(ingredient models.Recipe)
}
