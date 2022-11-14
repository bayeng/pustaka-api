package book

import "time"

type Book struct {
	Id       int
	Title    string
	Desc     string
	Price    int
	Rating   int
	CreateAt time.Time
	UpdateAt time.Time
}
