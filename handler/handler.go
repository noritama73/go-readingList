package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemRepository ItemRepository
}

func NewItemHandler(sqlService *SQLService) *ItemHandler {
	return &ItemHandler{
		itemRepository: sqlService,
	}
}

func (h *ItemHandler) ListItems(c echo.Context) error {
	return nil
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	var param GetItem

	if e := c.Bind(&param); e != nil {
		return e
	}

	item, e := h.itemRepository.GetItem(param.ID)
	if e != nil {
		return e
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) PutItemData(c echo.Context) error {
	var param PutItemData

	if e := c.Bind(&param); e != nil {
		return e
	}

	if e := h.itemRepository.PutItemData([]byte(param.Data)); e != nil {
		return e
	}
	return c.String(http.StatusOK, "Successfully put data!")
}
