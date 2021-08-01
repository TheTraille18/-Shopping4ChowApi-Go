package models

import "mime/multipart"

type Ingredient struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	S3Key           string `json:"s3Key"`
	Preferred_store string `json:"preferred_store"`
	File            multipart.File
}
