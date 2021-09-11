package dao

import (
	"context"
	"errors"
	"fmt"
	"log"
	"shopping4chow/cmd/shopping4chow/models"
	config "shopping4chow/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jackc/pgx/v4"
)

type MealDaoImpl struct {
}

func NewMealDAO() MealDaoImpl {
	return MealDaoImpl{}
}

func (m MealDaoImpl) GetMeal(conn *pgx.Conn, findMeal models.Meal) []models.Meal {
	log.Printf("Searching for %s\n", findMeal.Name)
	var meals []models.Meal

	haveMeal := make(map[string]*models.Meal)
	//rows, err := conn.Query(context.Background(), "select id,name from ingredient where name like $1", findMeal.Name+"%")
	rows, err := conn.Query(context.Background(), "select meal.name,recipe.id, recipe.amount, recipe.meal_id, recipe.name from recipe join meal on recipe.meal_id = meal.id where meal.name like $1", findMeal.Name+"%")

	if err != nil {
		fmt.Println("Error in select")
		fmt.Println(err)
	}

	for rows.Next() {

		log.Println(1)
		var meal models.Meal
		var mealName string
		var recipeName string
		var id int
		var mealId int
		var amount int
		rows.Scan(&mealName, &id, &amount, &mealId, &recipeName)
		m, mealExists := haveMeal[mealName]

		if mealExists {
			recipe := models.Recipe{Name: recipeName, Amount: amount}
			m.Recipes = append(m.Recipes, recipe)
		} else {
			meal.Name = mealName
			meal.ID = mealId
			recipe := models.Recipe{Name: recipeName, Amount: amount}
			meal.Recipes = append(meal.Recipes, recipe)
			haveMeal[mealName] = &meal
		}

	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	for key := range haveMeal {
		log.Printf("Meal Name %s\n", key)
		meals = append(meals, *haveMeal[key])
	}
	return meals
}

func (m MealDaoImpl) RemoveMeal(id int) error {
	log.Printf("Removing meal with id of %d", id)
	_, err := config.Conn.Exec(context.Background(), "delete from meal where id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m MealDaoImpl) GetAllMeals() []models.Meal {
	return nil
}

func (m MealDaoImpl) AddMeal(username string, meal models.Meal) (int, error) {
	var nameExists string
	config.Conn.QueryRow(context.Background(), "select name from meal where name=$1", meal.Name).Scan(&nameExists)
	if nameExists != "" {
		existsMsg := fmt.Sprintf("Meal %s exists\n", nameExists)
		return -1, errors.New(existsMsg)
	}

	input := &s3.PutObjectInput{
		Body:   meal.File,
		Bucket: aws.String("shopping4chow.com"),
		Key:    aws.String("local/Meal/" + meal.Name + ".png"),
	}

	_, err := config.Svc.PutObject(input)
	if err != nil {
		log.Fatal(err)
	}

	// input := &s3.PutObjectInput{
	// 	Body:   ingredient.File,
	// 	Bucket: aws.String("shopping4chow.com"),
	// 	Key:    aws.String("local/" + ingredient.S3Key),
	// }
	// _, err := config.Svc.PutObject(input)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var id int
	config.Conn.QueryRow(context.Background(), "insert into meal (name) values ($1) returning id", meal.Name).Scan(&id)
	_, err = config.Conn.Exec(context.Background(), "insert into user_join_meal(user_name,meal_id) values ($1,$2)", "DEV", id)
	if err != nil {
		log.Println(err)
	}
	return id, nil
}
