package service

import (
	"shopping4chow/cmd/shopping4chow/models"

	"github.com/jackc/pgx/v4"
)

type MealService interface {
	GetMeal(conn *pgx.Conn, findMeal models.Meal) []models.Meal
	RemoveMeal(meal models.Meal)
	GetAllMeals() []models.Meal
	AddMeal(meal models.Meal)
}
