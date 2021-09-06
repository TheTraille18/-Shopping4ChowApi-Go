package dao

import (
	"context"
	"fmt"
	"shopping4chow/cmd/shopping4chow/models"
	config "shopping4chow/configs"
)

type MealDaoImpl struct {
}

func NewMealDAO() MealDaoImpl {
	return MealDaoImpl{}
}

func (m MealDaoImpl) GetMeal(findMeal models.Meal) []models.Meal {
	return nil
}

func (m MealDaoImpl) RemoveMeal(Meal models.Meal) {

}

func (m MealDaoImpl) GetAllMeals() []models.Meal {
	return nil
}

func (m MealDaoImpl) AddMeal(meal models.Meal) int {
	var nameExists string
	config.Conn.QueryRow(context.Background(), "select name from meal where name=$1", meal.Name).Scan(&nameExists)
	if nameExists != "" {
		fmt.Printf("Name %s exists\n", nameExists)
		return -1
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
	fmt.Println(id)
	fmt.Println("***********************************************")
	return id
}
