package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"shopping4chow/internal/dao"
	"shopping4chow/internal/models"
	"shopping4chow/internal/service"
)

var ingredientDao dao.IngredientDao
var ingredientSvc service.IngredientService

func init() {
	ingredientDao = dao.NewDAO()
	ingredientSvc = service.NewService(ingredientDao)
}

func addIngredient(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w)

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
	ingredientSvc.AddIngredient(ingredient)
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
		ingredients := ingredientSvc.GetIngredient(searchIngredient)
		ingredientsJson, errJson := json.Marshal(ingredients)
		if errJson != nil {
			fmt.Println(errJson)
		}
		fmt.Println(string(ingredientsJson))
		w.Write(ingredientsJson)
	}
	//fmt.Println(ingredient.GetName())

}

func handleRequest() {
	http.HandleFunc("/addingredient", addIngredient)
	http.HandleFunc("/getIngredients", getIngredient)
	http.HandleFunc("/removeIngredient", removeIngredient)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("Running")
	handleRequest()
}
