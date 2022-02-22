package model

type ItemRepository interface {
	// ListItems() (result ItemList, e error)
	GetItem(id ID) (result Item, e error)
	PutItemData(data []byte) error
	// UpdateItemData(id ID, data []byte) error
	// DeleteItemData(id ID) error
}

type ID string

type Item struct {
	ID     ID
	Detail ItemDetail
}

type ItemDetail struct {
	Title      string
	Updated_at string
	Memo       string
	URL        string
	Tag        string
}

type ItemThumbnail struct {
	ID         ID
	Title      string
	Updated_at string
}

type ItemList []ItemThumbnail
