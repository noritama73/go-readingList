package handler

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLService struct {
	db *sql.DB
}

func NewSQLService() *SQLService {
	db, _ := sql.Open("sqlite3", "../app/item.sql")
	return &SQLService{
		db: db,
	}
}

func ListItems() (result ItemList, e error) {

	db, _ := sql.Open("sqlite3", "../app/item.sql")
	defer db.Close()

	stmt, e := db.Prepare(`CREATE TABLE IF NOT EXISTS item (
		id int AUTO_INCREMENT NOT NULL,
		title text NOT NULL,
		created_at datetime default current_timestamp,
		updated_at timestamp default current_timestamp pn update current_timestamp,
		url text,
		memo text,
		tag text,
		PRIMARY KEY (id)
	)`)
	defer stmt.Close()
	if e != nil {
		return
	}

	rows, e := db.Query("SELECT * FROM item")
	defer db.Close()
	if e != nil {
		return
	}
	var id, title string

	for rows.Next() {
		e := rows.Scan(&id, &title)
		if e != nil {
			return nil, e
		}
	}

	return
}
