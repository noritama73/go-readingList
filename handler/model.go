package handler

type ItemRepository interface {
	ListItems() (ItemList, error)
	GetItem(id ID) (Item, error)
	PutItemData(id ID, data []byte) error
	UpdateItemData(id ID, data []byte) error
}

type ID string

type Item struct {
	ID         string
	Title      string
	Updated_at string
	Data       ItemData
}

type ItemData struct {
	URL  string
	Memo string
	Tag  string
}

type ItemList []Item
