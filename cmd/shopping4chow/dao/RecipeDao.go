package dao

import "shopping4chow/cmd/shopping4chow/models"

type RecipeDao interface {
	GetReipe(recipe models.Recipe) []models.Recipe
	RemoveRecipe(recipe models.Recipe)
	GetAllRecipes() []models.Recipe
	AddRecipe(recipe models.Recipe)
}
