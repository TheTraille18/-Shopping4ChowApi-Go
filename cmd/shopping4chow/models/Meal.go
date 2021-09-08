package models

import "mime/multipart"

type Meal struct {
	ID         int      `json:"id"`
	User       string   `json:"user"`
	Name       string   `json:"name"`
	Recipes    []Recipe `json:"recipes"`
	PicId      string   `json:"picId"`
	RecipeText string   `json:"RecipeText"`
	Website    string   `json:"website"`
	File       multipart.File
}
