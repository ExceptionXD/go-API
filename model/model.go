package model

type Books struct {
	Id        string `form:"id" json:"id"`
	Title     string `form:"title" json:"title"`
	Isbn      string `form:"isbn" json:"isbn"`
	Firstname string `form:"firstname" json:"firstname"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Books
}
