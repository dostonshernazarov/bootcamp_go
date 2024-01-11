package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Autor_book struct {
	ID       int
	Autor_id int
	Book_id  int
}

type Autors struct {
	ID         int
	Full_name  string
	Birth_date int
	Books      Books
}

type Books struct {
	ID    int
	Title string
	Year  int
}

func main() {

	connStr := "user=doston password=doston dbname=najot_talim sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err != nil {
		panic(err)
	}

	resAutor := []byte(`
		{
			"full_name":"Oybek",
			"bith_date":1955,
			"books": {
				"title": "A.Navoiy",
				"year":1999
			}
		}
	`)

	var autor1 Autors

	if err := json.Unmarshal(resAutor, &autor1); err != nil {
		panic(err)
	}

	var resAuto Autors

	rowAutor := db.QueryRow(`INSERT INTO autors (full_name, birth_date) VALUES ($1, $2) returning id`, autor1.Full_name, autor1.Birth_date)

	if err1 := rowAutor.Scan(&resAuto.ID); err1 != nil {
		panic(err1)
	}

	fmt.Println("Autor added in table", resAuto.ID)

	var resBook Autors

	rowBook := db.QueryRow(`INSERT INTO books (title, year) VALUES ($1,$2) returning id`, autor1.Books.Title, autor1.Books.Year)

	if err = rowBook.Scan(&resBook.Books.ID); err != nil {
		panic(err)
	}

	fmt.Println("Book added in table", resBook.Books.ID)

	var resAutorBook Autor_book

	row_book_autor := db.QueryRow(`INSERT INTO autor_book (autor_id, book_id) VALUES ($1, $2) returning autor_id`, resAuto.ID, resBook.Books.ID)

	if err = row_book_autor.Scan(&resAutorBook.Autor_id); err != nil {
		panic(err)
	}

	fmt.Println("book_id and autor_id added in table")

}
