package model

type GetItem struct {
	ID ID `form:"id"`
}

type PutItemData struct {
	Data string `form:"data"`
}

type UpdateItemData struct {
	ID   ID     `form:"id"`
	Data string `form:"data"`
}

type DeleteItem struct {
	ID ID `form:"id"`
}

type PutDetailData struct {
	Title string `json:"title"`
	Memo  string `json:"memo"`
	URL   string `json:"url"`
	Tag   string `json:"tag"`
}
