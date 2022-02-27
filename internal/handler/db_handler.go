package handler

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/noritama73/go-readinglist/internal/model"
)

type SQLService struct {
	db *sql.DB
}

func NewSQLService() *SQLService {
	db, e := sql.Open("sqlite3", "./item.db")
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

func (s *SQLService) GetItem(id model.ID) (result model.Item, e error) {
	db := s.db

	det := result.Detail
	get_sql := db.QueryRow(`SELECT title, updated_at, memo, url, tag FROM item WHERE id = $1`, id)
	e = get_sql.Scan(&det.Title, &det.Updated_at, &det.Memo, &det.URL, &det.Tag)
	if e != nil {
		log.Println(e)
		return
	}

	result.ID = id
	result.Detail = det

	return
}

func (s *SQLService) ListItems() (result model.ItemList, e error) {
	db := s.db

	rows, e := db.Query("SELECT id, title, updated_at, tag FROM item")
	if e != nil {
		return
	}
	defer rows.Close()
	var item model.ItemThumbnail

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Title, &item.Updated_at, &item.Tag)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
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

	return nil
}

func (s *SQLService) UpdateItemData(id model.ID, data []byte) error {
	db := s.db

	var idCheck model.ID
	exist_sql := db.QueryRow(`SELECT id FROM item WHERE id = $1`, id)
	e := exist_sql.Scan(&idCheck)
	if e != nil {
		log.Println(e)
		return e
	}

	var DetailData model.PutDetailData
	if e := json.Unmarshal(data, &DetailData); e != nil {
		if err, ok := e.(*json.SyntaxError); ok {
			log.Println(string(data[err.Offset-7 : err.Offset+7]))
		}
		log.Println(e)
		return e
	}

	update_sql, e := db.Prepare(`UPDATE item SET
		title = $1,
		url = $2,
		memo = $3,
		tag = $4
		WHERE id = $5
	`)
	if e != nil {
		log.Println(e)
		return e
	}
	defer update_sql.Close()
	update_sql.Exec(DetailData.Title, DetailData.URL, DetailData.Memo, DetailData.Tag, id)

	return nil
}

func (s *SQLService) DeleteItemData(id model.ID) error {
	db := s.db

	delete_sql, e := db.Prepare(`DELETE FROM item WHERE id = $1`)
	if e != nil {
		log.Println(e)
		return e
	}
	defer delete_sql.Close()
	delete_sql.Exec(id)

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
	drop_sql, e := db.Prepare("DROP TABLE item")
	if e != nil {
		log.Println(e)
		return
	}
	defer drop_sql.Close()
	drop_sql.Exec()

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
}
