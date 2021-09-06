package models

type Unit int

const (
	None Unit = iota
	Each
	Peice
	Bag
	Bottle
	Box
	Pack
	Jar
	Can
	Bunch
	Roll
	Dozen
	Small
	Large
	Lbs
	Qt
	Oz
	Cup
	Dallon
	Tbsp
	Tsp
	G
	Kg
	Liter
	Milliliter
	Pis
)

type Recipe struct {
	Id            int    `json:"id"`
	Meal_id       int    `json:"meal"`
	Name          string `json:"name"`
	Ingredient_id int    `json:"ingredients"`
	Amount        int    `json:"amount"`
	Units         Unit   `json:"unit"`
}
