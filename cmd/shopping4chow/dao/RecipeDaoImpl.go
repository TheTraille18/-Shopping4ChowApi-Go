package dao

import (
	"context"
	"fmt"
	"log"
	"shopping4chow/cmd/shopping4chow/models"
	config "shopping4chow/configs"
)

type RecipeDaoImpl struct {
}

func NewRecipeDAO() RecipeDaoImpl {
	return RecipeDaoImpl{}
}

func (r RecipeDaoImpl) GetReipe(recipe models.Recipe) []models.Recipe {
	return nil
}
func (r RecipeDaoImpl) RemoveRecipe(recipe models.Recipe) {

}
func (r RecipeDaoImpl) GetAllRecipes() []models.Recipe {
	return nil
}
func (r RecipeDaoImpl) AddRecipe(recipe models.Recipe) {
	log.Println("Adding recipe")

	_, err := config.Conn.Exec(context.Background(), "insert into recipe (amount, name, units, ingredient_id, meal_id) values ($1,$2,$3,$4,$5)", recipe.Amount,
		recipe.Name, recipe.Units, recipe.Ingredient_id, recipe.Meal_id)
	if err != nil {
		fmt.Println(err)
	}
}
