package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"shopping4chow/cmd/shopping4chow/dao"
	"shopping4chow/cmd/shopping4chow/models"
	"shopping4chow/cmd/shopping4chow/service"
	config "shopping4chow/configs"
)

var ingredientDao dao.IngredientDao
var ingredientSvc service.IngredientService

var mealDao dao.MealDao
var mealSvc service.MealService

var recipeDao dao.RecipeDao
var recipeSvc service.RecipeService

func init() {

	ingredientDao = dao.NewIngredientDAO()
	ingredientSvc = service.NewIngredientService(ingredientDao)

	recipeDao = dao.NewRecipeDAO()
	recipeSvc = service.NewRecipeService(recipeDao)

	mealDao = dao.NewMealDAO()
	mealSvc = service.NewMealService(mealDao)
}

func addIngredient(w http.ResponseWriter, r *http.Request) {
	//Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	err := r.ParseMultipartForm(200000000)
	if err != nil {
		fmt.Println(err)
	}
	name := r.MultipartForm.Value["name"][0]
	fileHeader := r.MultipartForm.File["file"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println(err)
	}
	ingredient := models.Ingredient{}
	ingredient.Name = name
	ingredient.File = file

	s3Key := ingredient.Name + ".png"
	ingredient.S3Key = s3Key

	err = ingredientSvc.AddIngredient(ingredient)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Ingredient already Exists"))
	}
	w.WriteHeader(200)
}

func removeIngredient(w http.ResponseWriter, r *http.Request) {
	//Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var removeIngredient models.Ingredient
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		err = json.Unmarshal(body, &removeIngredient)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf("Remove Ingredient with id %x with key %s", removeIngredient.Id, removeIngredient.S3Key)
		ingredientSvc.RemoveIngredient(removeIngredient)
	}
}

func getIngredient(w http.ResponseWriter, r *http.Request) {

	//Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" { //Request sends both an OPTIONS and POST
		var searchIngredient models.Ingredient
		//fmt.Println(r.Body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		fmt.Printf("Body: %s\n", string(body))
		err = json.Unmarshal(body, &searchIngredient)
		if err != nil {
			fmt.Println(err)
		}
		ingredients := ingredientSvc.GetIngredient(config.Conn, searchIngredient)
		ingredientsJson, errJson := json.Marshal(ingredients)
		if errJson != nil {
			log.Println(errJson)
		}
		log.Println(string(ingredientsJson))
		w.Write(ingredientsJson)
	}
	//fmt.Println(ingredient.GetName())
}

func addMeal(w http.ResponseWriter, r *http.Request) {

	//Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	err := r.ParseMultipartForm(200000000)
	if err != nil {
		fmt.Println(err)
	}

	//Struct to unmarshal into
	var meal models.Meal
	user := struct {
		User string `json:"user"`
	}{}

	m := r.MultipartForm.Value["meal"][0]
	fileHeader := r.MultipartForm.File["file"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(m), &meal)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(m), &user)
	if err != nil {
		log.Println(err)
	}

	meal.File = file
	err = mealSvc.AddMeal(user.User, meal)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	//var meal models.Meal
	/*
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body",
					http.StatusInternalServerError)
			}
			fmt.Printf("Body %s\n", string(body))

			err = json.Unmarshal(body, &meal)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Meal created %+v\n", meal)

			mealSvc.AddMeal(meal)

		}
	*/
}

func removeMeal(w http.ResponseWriter, r *http.Request) {

	//Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		MealID := struct {
			Id int `json:"id"`
		}{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		}

		err = json.Unmarshal([]byte(body), &MealID)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		}

		err = mealSvc.RemoveMeal(MealID.Id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(200)
	}
}

func getMeal(w http.ResponseWriter, r *http.Request) {

	//Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" { //Request sends both an OPTIONS and POST
		var searchMeal models.Meal
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		fmt.Printf("Meal Body: %s\n", string(body))
		err = json.Unmarshal(body, &searchMeal)
		if err != nil {
			fmt.Println(err)
		}
		meals := mealSvc.GetMeal(config.Conn, searchMeal)
		mealsJson, errJson := json.Marshal(meals)
		if errJson != nil {
			log.Println(errJson)
		}
		log.Printf("Meals %s", string(mealsJson))
		w.Write(mealsJson)
	}
	//fmt.Println(ingredient.GetName())
}

func handleRequest() {
	http.HandleFunc("/addingredient", addIngredient)
	http.HandleFunc("/getIngredients", getIngredient)
	http.HandleFunc("/removeIngredient", removeIngredient)
	http.HandleFunc("/addmeal", addMeal)
	http.HandleFunc("/getMeals", getMeal)
	http.HandleFunc("/removeMeal", removeMeal)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	log.Println("Running")
	handleRequest()
}
