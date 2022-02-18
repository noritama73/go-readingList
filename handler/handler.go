package handler

type ItemHandler struct {
	itemRepository ItemRepository
}

func NewItemHandler(sqlSerive *SQLService) *ItemHandler {
	return &ItemHandler{
		itemRepository: sqlService,
	}
}
