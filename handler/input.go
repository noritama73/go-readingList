package handler

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
