package main

type Book struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Price int64  `db:"price" json:"price"`
}
