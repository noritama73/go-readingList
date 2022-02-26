package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/noritama73/go-readinglist/model"
)

type ItemHandler struct {
	itemRepository model.ItemRepository
}

func NewItemHandler(sqlService *SQLService) *ItemHandler {
	return &ItemHandler{
		itemRepository: sqlService,
	}
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	var param model.GetItem

	if e := c.Bind(&param); e != nil {
		return e
	}

	item, e := h.itemRepository.GetItem(param.ID)
	if e != nil {
		return e
	}

	return apiResponseOK(c, item)
}

func (h *ItemHandler) ListItems(c echo.Context) error {
	itemList, e := h.itemRepository.ListItems()
	if e != nil {
		return e
	}
	return apiResponseOK(c, itemList)
}

func (h *ItemHandler) PutItemData(c echo.Context) error {
	var param model.PutItemData

	if e := c.Bind(&param); e != nil {
		return e
	}

	if e := h.itemRepository.PutItemData([]byte(param.Data)); e != nil {
		return e
	}
	return c.String(http.StatusOK, "Successfully put data!")
}

func (h *ItemHandler) UpdateItemData(c echo.Context) error {
	var param model.UpdateItemData

	if e := c.Bind(&param); e != nil {
		return e
	}

	if e := h.itemRepository.UpdateItemData(param.ID, []byte(param.Data)); e != nil {
		return e
	}

	return c.String(http.StatusOK, "Successfully updata data!")
}

func (h *ItemHandler) DeleteItemData(c echo.Context) error {
	var param model.DeleteItem

	if e := c.Bind(&param); e != nil {
		return e
	}

	if e := h.itemRepository.DeleteItemData(param.ID); e != nil {
		return e
	}

	return c.String(http.StatusOK, "Successfully delete data!")
}
