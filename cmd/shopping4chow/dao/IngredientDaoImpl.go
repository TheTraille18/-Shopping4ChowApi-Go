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

type IngredientDaoImpl struct {
}

func init() {

	//defer conn.Close(context.Background())

}

func NewIngredientDAO() *IngredientDaoImpl {
	return &IngredientDaoImpl{}
}

func (i *IngredientDaoImpl) GetIngredient(conn *pgx.Conn, ingredient models.Ingredient) []models.Ingredient {
	//fmt.Printf()
	log.Printf("Searching for for %s\n", ingredient.Name)
	var ingredients []models.Ingredient
	rows, err := conn.Query(context.Background(), "select id,name,s3key from ingredient where name like $1", ingredient.Name+"%")
	if err != nil {
		fmt.Println("Error in select")
		fmt.Println(err)
	}
	for rows.Next() {
		var ingredient models.Ingredient
		var id int
		var name, s3key string
		rows.Scan(&id, &name, &s3key)
		ingredient.Name = name
		ingredient.Id = id
		ingredient.S3Key = s3key
		ingredients = append(ingredients, ingredient)
	}
	//var val []string
	return ingredients
}

func (i *IngredientDaoImpl) RemoveIngredient(ingredient models.Ingredient) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String("shopping4chow.com"),
		Key:    aws.String("local/" + ingredient.S3Key),
	}
	_, err := config.Svc.DeleteObject(input)
	if err != nil {
		log.Fatal(err)
	}
	_, sqlErr := config.Conn.Exec(context.Background(), "delete from ingredient where id=$1", ingredient.Id)
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
}

func (i *IngredientDaoImpl) GetAllIngredients() []models.Ingredient {
	return nil
}

func (i *IngredientDaoImpl) AddIngredient(ingredient models.Ingredient) error {
	var nameExists string
	config.Conn.QueryRow(context.Background(), "select name from ingredient where name=$1", ingredient.Name).Scan(&nameExists)
	if nameExists != "" {
		fmt.Printf("Name %s exists\n", nameExists)
		return errors.New("Ingredient Already Exists")
	}

	input := &s3.PutObjectInput{
		Body:   ingredient.File,
		Bucket: aws.String("shopping4chow.com"),
		Key:    aws.String("local/" + ingredient.S3Key),
	}
	_, err := config.Svc.PutObject(input)
	if err != nil {
		log.Fatal(err)
		return err
	}

	flag, sqlErr := config.Conn.Exec(context.Background(), "insert into ingredient (name,s3key) values ($1,$2)", ingredient.Name, ingredient.S3Key)
	if sqlErr != nil {
		return sqlErr
	}
	fmt.Println(flag.RowsAffected())
	return nil
	//fmt.Printf("Result: %s", ingredient.GetName())
}
