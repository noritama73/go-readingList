package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/noritama73/go-readinglist/internal/model"
)

type SQLService struct {
	db *sql.DB
}

func NewSQLService() *SQLService {
	log.SetFlags(log.Lshortfile)
	db, e := sql.Open(os.Getenv("DRIVER"), os.Getenv("DSN"))
	if e != nil {
		log.Fatalln(e)
	}

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
	log.Println(id)
	get_sql := db.QueryRow(`SELECT title, updated_at, memo, url, tag FROM item WHERE id = ?`, id)
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
		log.Println(e)
		return
	}
	defer rows.Close()
	var item model.ItemThumbnail

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Title, &item.Updated_at, &item.Tag)
		if err != nil {
			log.Println(err)
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
			log.Println(string(data[err.Offset-3 : err.Offset+3]))
		}
		log.Println(e)
		return e
	}

	insert_sql, e := db.Prepare(`INSERT INTO item (
		id, title, url, memo, tag
	)
	VALUES (
		?, ?, ?, ?, ?
	)`)
	if e != nil {
		log.Println(e)
		return e
	}
	defer insert_sql.Close()
	u, e := uuid.NewRandom()
	if e != nil {
		log.Println(e)
		return e
	}
	uu := u.String()
	insert_sql.Exec(uu, DetailData.Title, DetailData.URL, DetailData.Memo, DetailData.Tag)

	return nil
}

func (s *SQLService) UpdateItemData(id model.ID, data []byte) error {
	db := s.db

	var idCheck model.ID
	exist_sql := db.QueryRow(`SELECT id FROM item WHERE id = ?`, id)
	e := exist_sql.Scan(&idCheck)
	if e != nil {
		log.Println(e)
		return e
	}

	var DetailData model.PutDetailData
	if e := json.Unmarshal(data, &DetailData); e != nil {
		if err, ok := e.(*json.SyntaxError); ok {
			log.Println(string(data[err.Offset-3 : err.Offset+3]))
		}
		log.Println(e)
		return e
	}

	update_sql, e := db.Prepare(`UPDATE item SET
		title = ?,
		url = ?,
		memo = ?,
		tag = ?
		WHERE id = ?
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

	delete_sql, e := db.Prepare(`DELETE FROM item WHERE id = ?`)
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
	log.SetFlags(log.Lshortfile)
	db, e := sql.Open(os.Getenv("DRIVER"), os.Getenv("DSN"))
	if e != nil {
		log.Fatalln(e)
	}

	create_sql, e := db.Prepare(`CREATE TABLE item (
		id TEXT NOT NULL,
		title TEXT NOT NULL,
		created_at DATETIME DEFAULT current_timestamp,
		updated_at DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
		url TEXT,
		memo TEXT,
		tag TEXT,
		PRIMARY KEY(ID(128))
	)`)
	if e != nil {
		log.Fatalln(e)
	}
	defer create_sql.Close()
	create_sql.Exec()

	return &FakeSQLService{
		&SQLService{
			db: db,
		},
	}
}

func (s *FakeSQLService) GetItemTop() (id model.ID, result model.Item, e error) {
	db := s.db

	det := result.Detail
	gettop_sql := db.QueryRow(`SELECT id, title, updated_at, memo, url, tag FROM item LIMIT 1`)
	e = gettop_sql.Scan(&id, &det.Title, &det.Updated_at, &det.Memo, &det.URL, &det.Tag)
	if e != nil {
		log.Println(e)
		return
	}

	result.Detail = det

	return
}

func (s *FakeSQLService) DeleteAll() {
	db := s.db
	drop_sql, e := db.Prepare("TRUNCATE TABLE item")
	if e != nil {
		log.Println(e)
		return
	}
	defer drop_sql.Close()
	drop_sql.Exec()
}
