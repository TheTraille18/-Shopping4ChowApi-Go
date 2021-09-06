package service

import (
	"fmt"
	"shopping4chow/cmd/shopping4chow/dao"
	"shopping4chow/cmd/shopping4chow/models"
)

type MealServiceImpl struct {
	MealDao dao.MealDao
}

func NewMealService(mealdao dao.MealDao) MealServiceImpl {
	return MealServiceImpl{mealdao}
}

func (m MealServiceImpl) GetMeal(findMeal models.Meal) []models.Meal {
	return nil
}

func (m MealServiceImpl) RemoveMeal(meal models.Meal) {

}
func (m MealServiceImpl) GetAllMeals() []models.Meal {
	return nil
}
func (m MealServiceImpl) AddMeal(meal models.Meal) {
	id := m.MealDao.AddMeal(meal)
	recipeDao := dao.NewRecipeDAO()
	recipeSVC := NewRecipeService(recipeDao)

	fmt.Printf("Create Meal with id %x", id)
	for _, recipe := range meal.Recipes {
		recipeSVC.AddRecipe(recipe)
	}
}
