package handler

import (
	"database/sql"
	"encoding/json"

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

func (s *SQLService) ListItems() (result ItemList, e error) {
	db := s.db
	defer db.Close()

	rows, e := db.Query("SELECT * FROM item")
	if e != nil {
		return
	}
	defer rows.Close()
	var id, title string

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			return
		}
	}

	return
}

func (s *SQLService) GetItem(id ID) (result Item, e error) {
	db := s.db
	defer db.Close()

	get_sql, e := db.Query(`SELECT * FROM item WHERE id = $1`, id)
	if e != nil {
		return
	}
	defer get_sql.Close()

	var title, memo, url, tag string
	for get_sql.Next() {
		err := get_sql.Scan(&title, &memo, &url, &tag)
		if err != nil {
			return
		}
	}
	det := ItemDetail{
		Title: title,
		Memo:  memo,
		URL:   url,
		Tag:   tag,
	}
	result = Item{
		ID:     id,
		Detail: det,
	}

	return
}

func (s *SQLService) PutItemData(data []byte) error {
	db := s.db
	defer db.Close()

	create_sql, e := db.Prepare(`CREATE TABLE IF NOT EXISTS item (
		id int AUTO_INCREMENT NOT NULL,
		title text NOT NULL,
		created_at datetime default current_timestamp,
		updated_at timestamp default current_timestamp pn update current_timestamp,
		url text,
		memo text,
		tag text,
		PRIMARY KEY (id)
	)`)
	if e != nil {
		return e
	}
	defer create_sql.Close()
	create_sql.Exec()

	var DetailData ItemDetail
	if e := json.Unmarshal(data, &DetailData); e != nil {
		return e
	}

	insert_sql, e := db.Prepare(`INSERT INTO item (
		title, url, memo, tag
	)
	VALUE (
		$1, $2, $3, $4
	)`)
	if e != nil {
		return e
	}
	defer insert_sql.Close()
	insert_sql.Exec(DetailData.Title, DetailData.URL, DetailData.Memo, DetailData.Tag)

	return nil
}
