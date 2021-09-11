package models

type Unit int

var unitMap = map[int]string{
	0:  "none",
	1:  "each",
	2:  "peice",
	3:  "bag",
	4:  "bottle",
	5:  "box",
	6:  "case",
	7:  "pack",
	8:  "jar",
	9:  "can",
	10: "bunch",
	11: "roll",
	12: "dozen",
	13: "small",
	14: "large",
	15: "lbs",
	16: "qt",
	17: "oz",
	18: "cup",
	19: "gallon",
	20: "tbsp",
	21: "tsp",
	22: "g",
	23: "kg",
	24: "liter",
	25: "milliliter",
	26: "pis",
}

type Recipe struct {
	Id            int    `json:"id"`
	Meal_id       int    `json:"meal"`
	Name          string `json:"name"`
	Ingredient_id int    `json:"ingredientId"`
	Amount        int    `json:"amount"`
	UnitInt       int    `json:"unit"`
	Units         string
}

func (recipe *Recipe) SetUnits() {
	recipe.Units = unitMap[recipe.UnitInt]
}
