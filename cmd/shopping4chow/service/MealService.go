package service

import (
	"shopping4chow/cmd/shopping4chow/models"

	"github.com/jackc/pgx/v4"
)

type MealService interface {
	GetMeal(conn *pgx.Conn, findMeal models.Meal) []models.Meal
	RemoveMeal(id int) error
	GetAllMeals() []models.Meal
	AddMeal(user string, meal models.Meal) error
}
