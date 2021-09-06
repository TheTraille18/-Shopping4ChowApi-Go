package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"shopping4chow/cmd/shopping4chow/models"
	"testing"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func TestMain(m *testing.M) {
	// 0. flag.Parse() if you need flags

	url := "postgresql://" + "postgres" + ":" + "postgres" + "@" + "192.168.1.114" + ":5432/" + "postgres"
	pgxConn, connErr := pgx.Connect(context.Background(), url)
	if connErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database2: %v\n", connErr)
		os.Exit(1)
	}
	conn = pgxConn
	importSQL()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func importSQL() {
	log.Println("Running TestDump......")

	//Create Schema
	b, err := ioutil.ReadFile("shopping4chow.sql")
	if err != nil {
		log.Fatalf("Error reading file %v\n", err)
	}
	sql := string(b)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		log.Fatalf("Error in running query %v\n", err)
	}

	//Load Teat Ingredient Data
	testData, err := ioutil.ReadFile("../test_data/FakeDataTest.sql")
	if err != nil {
		log.Fatalf("Error reading test data file %v\n", err)
	}
	sqlTest := string(testData)
	_, err = conn.Exec(context.Background(), sqlTest)
	if err != nil {
		log.Fatalf("Error loading Test Data %v\n", err)
	}
}

// func TestUserStore(t *testing.T) {
// 	db, err := sql.Open("postgres", "host=192.168.1.114 port=5432 user=postgres password=postgres sslmode=disable dbname=test_user_store")
// 	if err != nil {
// 		panic(fmt.Errorf("sql.Open() err = %s", err))
// 	}
// 	defer db.Close()
// 	us := &UserStore{
// 		sql: db,
// 	}
// 	t.Run("Find", testUserStore_Find(us))
// 	t.Run("Create", testUserStore_Find(us))
// 	t.Run("Delete", testUserStore_Find(us))
// 	t.Run("Subscribe", testUserStore_Find(us))
// 	// teardown
// }

func TestIngredient_Get(t *testing.T) {
	arg1 := "Fake"
	want1 := "[{\"id\":1,\"name\":\"FakeIngredient1\",\"s3Key\":\"FakeKey\",\"preferred_store\":\"\",\"File\":null},{\"id\":2,\"name\":\"FakeIngredient2\",\"s3Key\":\"FakeKey\",\"preferred_store\":\"\",\"File\":null},{\"id\":3,\"name\":\"FakeIngredient3\",\"s3Key\":\"FakeKey\",\"preferred_store\":\"\",\"File\":null},{\"id\":4,\"name\":\"FakesIngredient4\",\"s3Key\":\"FakeKey\",\"preferred_store\":\"\",\"File\":null}]"

	arg2 := "Ingredient"
	want2 := "null"

	arg3 := "ZZZ"
	want3 := "null"

	arg4 := "Fakes"
	want4 := "[{\"id\":4,\"name\":\"FakesIngredient4\",\"s3Key\":\"FakeKey\",\"preferred_store\":\"\",\"File\":null}]"

	t.Run("Get Test 1", IngredientGet(arg1, want1))
	t.Run("Get Test 2", IngredientGet(arg2, want2))
	t.Run("Get Test 3", IngredientGet(arg3, want3))
	t.Run("Get Test 4", IngredientGet(arg4, want4))

}

func IngredientGet(arg string, want string) func(t *testing.T) {
	f := func(t *testing.T) {

		ingredientDao := NewIngredientDAO()

		ingredient := models.Ingredient{Name: arg}
		ingredients := ingredientDao.GetIngredient(conn, ingredient)
		ingredientJson, err := json.Marshal(ingredients)
		if err != nil {
			log.Fatalf("Marshal Error %v\n", err)
		}
		got := string(ingredientJson)
		if got != want {
			log.Fatalf("GetIngredient with name %s, got = %s, want %s", arg, got, want)
		}
	}

	return f
}

func insert() {
	flag, sqlErr := conn.Exec(context.Background(), "insert into ingredient (name) values (FakeIngrdeint1)")
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
	fmt.Println(flag.RowsAffected())
}
