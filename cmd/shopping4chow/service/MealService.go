package service

import "shopping4chow/cmd/shopping4chow/models"

type MealService interface {
	GetMeal(findMeal models.Meal) []models.Meal
	RemoveMeal(meal models.Meal)
	GetAllMeals() []models.Meal
	AddMeal(meal models.Meal)
}
