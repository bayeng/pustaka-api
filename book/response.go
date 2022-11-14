package book

type BookResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Desc   string `json:"desc"`
	Rating int    `json:"rating"`
}
