package dao

import "shopping4chow/cmd/shopping4chow/models"

type MealDao interface {
	GetMeal(findMeal models.Meal) []models.Meal
	RemoveMeal(Meal models.Meal)
	GetAllMeals() []models.Meal
	AddMeal(Meal models.Meal) int
}
