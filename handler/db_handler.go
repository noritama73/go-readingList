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
	db, e := sql.Open("sqlite3", "../app/item.db")
	if e != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatalln(e)
	}
	create_sql, e := db.Prepare(`CREATE TABLE IF NOT EXISTS item (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		title TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now','localtime')),
		updated_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now','localtime')),
		url TEXT,
		memo TEXT,
		tag TEXT
	)`)
	if e != nil {
		log.Fatalln(e)
	}
	defer create_sql.Close()
	create_sql.Exec()

	return &SQLService{
		db: db,
	}
}

func (s *SQLService) DestructDB() {
	s.db.Close()
}

func (s *SQLService) ListItems() (result model.ItemList, e error) {
	db := s.db

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

func (s *SQLService) GetID() (id model.ID, err error) { // テスト用
	db := s.db

	row := db.QueryRow("SELECT * FROM item LIMIT 1")

	log.Println(row)

	return
}

func (s *SQLService) GetItem(id model.ID) (result model.Item, e error) {
	db := s.db
	log.Println(id)
	
	det := result.Detail
	get_sql := db.QueryRow(`SELECT title, memo, url, tag FROM item WHERE id = $1`, id)
	e = get_sql.Scan(&det.Title, &det.Memo, &det.URL, &det.Tag)
	if e != nil {
		log.Println(e)
		return
	}
	
	result.ID = id
	result.Detail = det

	log.Println(result)

	return
}

func (s *SQLService) PutItemData(data []byte) error {
	db := s.db

	var DetailData model.PutDetailData
	if e := json.Unmarshal(data, &DetailData); e != nil {
		if err, ok := e.(*json.SyntaxError); ok {
			log.Println(string(data[err.Offset-7 : err.Offset+7]))
		}
		log.Println(e)
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
	
	var det model.ItemDetail
	get_sql := db.QueryRow(`SELECT title, memo, url, tag FROM item LIMIT 1`)
	e = get_sql.Scan(&det.Title, &det.Memo, &det.URL, &det.Tag)
	if e != nil {
		log.Println(e)
		return e
	}
	log.Println(det)
	
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
