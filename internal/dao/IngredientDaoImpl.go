package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"shopping4chow/internal/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jackc/pgx/v4"
)

type IngredientDaoImpl struct {
}

var svc *s3.S3
var conn *pgx.Conn
var connErr error

func init() {
	svc = s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
	host := os.Getenv("S4C_HOST")
	database := os.Getenv("S4C_DATABASE")
	user := os.Getenv("S4C_USERNAME")
	password := os.Getenv("S4C_PASSWORD")
	fmt.Printf("Host %s\n", host)
	fmt.Printf("Database %s\n", database)
	url := "postgresql://" + user + ":" + password + "@" + host + ":5432/" + database
	conn, connErr = pgx.Connect(context.Background(), url)
	if connErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", connErr)
	}
	//defer conn.Close(context.Background())
}

func NewDAO() *IngredientDaoImpl {
	return &IngredientDaoImpl{}
}

func (i *IngredientDaoImpl) GetIngredient(ingredient models.Ingredient) []models.Ingredient {
	fmt.Printf("search for %s\n", ingredient.Name)
	var ingredients []models.Ingredient
	rows, err := conn.Query(context.Background(), "select id,name,s3key from ingredient where name like $1", ingredient.Name+"%")
	if err != nil {
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
	_, err := svc.DeleteObject(input)
	if err != nil {
		log.Fatal(err)
	}
	_, sqlErr := conn.Exec(context.Background(), "delete from ingredient where id=$1", ingredient.Id)
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
}

func (i *IngredientDaoImpl) GetAllIngredients() []models.Ingredient {
	return nil
}

func (i *IngredientDaoImpl) AddIngredient(ingredient models.Ingredient) {
	var nameExists string
	conn.QueryRow(context.Background(), "select name from ingredient where name=$1", ingredient.Name).Scan(&nameExists)
	if nameExists != "" {
		fmt.Printf("Name %s exists\n", nameExists)
	}

	input := &s3.PutObjectInput{
		Body:   ingredient.File,
		Bucket: aws.String("shopping4chow.com"),
		Key:    aws.String("local/" + ingredient.S3Key),
	}
	_, err := svc.PutObject(input)
	if err != nil {
		log.Fatal(err)
	}
	flag, sqlErr := conn.Exec(context.Background(), "insert into ingredient (name,s3key) values ($1,$2)", ingredient.Name, ingredient.S3Key)
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
	fmt.Println(flag.RowsAffected())

	//fmt.Printf("Result: %s", ingredient.GetName())
}
