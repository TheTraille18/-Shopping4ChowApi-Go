package service

import (
	"fmt"
	"shopping4chow/cmd/shopping4chow/dao"
	"shopping4chow/cmd/shopping4chow/models"

	"github.com/jackc/pgx/v4"
)

type IngredientServiceImpl struct {
	IngredientDAO dao.IngredientDao
}

func NewIngredientService(dao dao.IngredientDao) *IngredientServiceImpl {
	return &IngredientServiceImpl{IngredientDAO: dao}
}

//Get Ingredient from text entered in search bar
func (i IngredientServiceImpl) GetIngredient(conn *pgx.Conn, ingredient models.Ingredient) []models.Ingredient {
	ingredients := i.IngredientDAO.GetIngredient(conn, ingredient)
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
