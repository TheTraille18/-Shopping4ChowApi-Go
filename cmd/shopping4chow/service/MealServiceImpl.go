package service

import (
	"fmt"
	"shopping4chow/cmd/shopping4chow/dao"
	"shopping4chow/cmd/shopping4chow/models"

	"github.com/jackc/pgx/v4"
)

type MealServiceImpl struct {
	MealDao dao.MealDao
}

func NewMealService(mealdao dao.MealDao) MealServiceImpl {
	return MealServiceImpl{mealdao}
}

func (m MealServiceImpl) GetMeal(conn *pgx.Conn, findMeal models.Meal) []models.Meal {
	meals := m.MealDao.GetMeal(conn, findMeal)
	return meals
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

	for _, recipe := range meal.Recipes {
		fmt.Println(recipe)
		recipe.Meal_id = id
		recipeSVC.AddRecipe(recipe)
	}
}
