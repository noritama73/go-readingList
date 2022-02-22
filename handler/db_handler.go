package handler

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/noritama73/go-readinglist/model"
)

type SQLService struct {
	db *sql.DB
}

func NewSQLService() *SQLService {
	db, e := sql.Open("sqlite3", "../app/item.sqlite3")
	if e != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(e)
	}
	return &SQLService{
		db: db,
	}
}

func (s *SQLService) ListItems() (result model.ItemList, e error) {
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

func (s *SQLService) GetItem(id model.ID) (result model.Item, e error) {
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
	det := model.ItemDetail{
		Title: title,
		Memo:  memo,
		URL:   url,
		Tag:   tag,
	}
	result = model.Item{
		ID:     id,
		Detail: det,
	}

	return
}

func (s *SQLService) PutItemData(data []byte) error {
	db := s.db
	defer db.Close()

	log.SetFlags(log.Lshortfile)

	create_sql, e := db.Prepare(`CREATE TABLE IF NOT EXISTS item (
		id INTEGER AUTO_INCREMENT NOT NULL,
		title TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now','localtime')),
		updated_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now','localtime')),
		url TEXT,
		memo TEXT,
		tag TEXT,
		PRIMARY KEY (id)
	)`)
	if e != nil {
		log.Println(e)
		return e
	}
	defer create_sql.Close()
	create_sql.Exec()

	var DetailData model.PutDetailData
	if e := json.Unmarshal(data, &DetailData); e != nil {
		if err, ok := e.(*json.SyntaxError); ok {
			log.Println(string(data[err.Offset-7:err.Offset+7]))
			log.Println(e)
		}
		return e
	}

	insert_sql, e := db.Prepare(`INSERT INTO item (
		title, url, memo, tag
	)
	VALUES (
		$1, $2, $3, $4
	)`)
	if e != nil {
		log.Println(e)
		return e
	}
	defer insert_sql.Close()
	insert_sql.Exec(DetailData.Title, DetailData.URL, DetailData.Memo, DetailData.Tag)

	return nil
}

type FakeSQLService struct {
	*SQLService
}

func NewFakeSQLService() *FakeSQLService {
	return &FakeSQLService{
		SQLService: NewSQLService(),
	}
}

func (s *FakeSQLService) DeleteAll() {
	db := s.db
	defer db.Close()
	drop_sql, e := db.Prepare("DROP TABLE item")
	if e != nil {
		log.Println(e)
		return
	}
	defer drop_sql.Close()
	drop_sql.Exec()
}
