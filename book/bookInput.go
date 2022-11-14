package book

import (
	"encoding/json"
)

type BookInput struct {
	Title  string      `json:"title" binding:"required"`
	Price  json.Number `json:"price" binding:"required"`
	Desc   string      `json:"desc" binding:"required"`
	Rating json.Number `json:"rating" binding:"required"`
}
