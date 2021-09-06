package dao

import (
	"shopping4chow/cmd/shopping4chow/models"

	"github.com/jackc/pgx/v4"
)

type IngredientDao interface {
	GetIngredient(conn *pgx.Conn, findIngredients models.Ingredient) []models.Ingredient
	RemoveIngredient(ingredient models.Ingredient)
	GetAllIngredients() []models.Ingredient
	AddIngredient(ingredient models.Ingredient)
}
