package service

import (
	"fmt"
	"shopping4chow/internal/dao"
	"shopping4chow/internal/models"
)

type IngredientServiceImpl struct {
	IngredientDAO dao.IngredientDao
}

func NewService(dao dao.IngredientDao) *IngredientServiceImpl {
	return &IngredientServiceImpl{IngredientDAO: dao}
}

//Get Ingredient from text entered in search bar
func (i IngredientServiceImpl) GetIngredient(ingredient models.Ingredient) []models.Ingredient {
	ingredients := i.IngredientDAO.GetIngredient(ingredient)
	return ingredients
}

func (i IngredientServiceImpl) RemoveIngredient(ingredient models.Ingredient) {
	i.IngredientDAO.RemoveIngredient(ingredient)
}

func (i IngredientServiceImpl) GetAllIngredients() []models.Ingredient {
	return nil
}

func (i IngredientServiceImpl) AddIngredient(ingredient models.Ingredient) {
	fmt.Println(ingredient.Name)
	i.IngredientDAO.AddIngredient(ingredient)
}
