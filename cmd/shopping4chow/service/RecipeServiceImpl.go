package service

import (
	"shopping4chow/cmd/shopping4chow/dao"
	"shopping4chow/cmd/shopping4chow/models"
)

type RecipeServiceImpl struct {
	RecipeDao dao.RecipeDao
}

func NewRecipeService(dao dao.RecipeDao) RecipeServiceImpl {
	return RecipeServiceImpl{dao}
}

func (r RecipeServiceImpl) GetRecipe(recipe models.Recipe) []models.Recipe {
	return nil
}
func (r RecipeServiceImpl) RemoveRecipe(recipe models.Recipe) {

}
func (r RecipeServiceImpl) GetAllRecipes() []models.Recipe {
	return nil
}
func (r RecipeServiceImpl) AddRecipe(recipe models.Recipe) {
	r.AddRecipe(recipe)
}
