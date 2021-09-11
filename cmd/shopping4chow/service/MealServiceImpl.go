package service

import (
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

func (m MealServiceImpl) RemoveMeal(id int) error {
	err := m.MealDao.RemoveMeal(id)
	if err != nil {
		return err
	}
	return err

}
func (m MealServiceImpl) GetAllMeals() []models.Meal {
	return nil
}
func (m MealServiceImpl) AddMeal(user string, meal models.Meal) error {
	id, err := m.MealDao.AddMeal(user, meal)
	if err != nil {
		return err
	}
	recipeDao := dao.NewRecipeDAO()
	recipeSVC := NewRecipeService(recipeDao)

	for _, recipe := range meal.Recipes {
		recipe.SetUnits()
		recipe.Meal_id = id
		recipeSVC.AddRecipe(recipe)
	}
	return nil
}
