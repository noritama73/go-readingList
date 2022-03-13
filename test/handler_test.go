package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/noritama73/go-readinglist/internal/handler"
	"github.com/noritama73/go-readinglist/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testEcho  = echo.New()
	testTitle = "test title"
	testMemo  = "test memo"
	testURL   = "http://xxx.test"
	testTag   = "test tag"
)

func testItemRequest() string {
	return fmt.Sprintf(`
		{
			"title":"%s",
			"memo":"%s",
			"url":"%s",
			"tag":"%s"
		}	
	`, testTitle, testMemo, testURL, testTag)
}

func testPutItem(e *echo.Echo, body string) (*httptest.ResponseRecorder, echo.Context) {
	f := make(url.Values)
	f.Set("data", body)
	req := httptest.NewRequest(echo.POST, "/item", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

func testItemList() (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(echo.GET, "/itemList", nil)
	rec := httptest.NewRecorder()
	c := testEcho.NewContext(req, rec)
	return rec, c
}

func Test_ItemHandler(t *testing.T) {
	fakeDB := handler.NewFakeSQLService()
	hdl := handler.NewItemHandler(fakeDB.SQLService)
	defer fakeDB.DeleteAll()

	t.Run("レコード作成", func(t *testing.T) {
		_, c := testPutItem(testEcho, testItemRequest())
		require.NoError(t, hdl.PutItemData(c))
	})
}

func Test_ItemList(t *testing.T) {
	fakeDB := handler.NewFakeSQLService()
	hdl := handler.NewItemHandler(fakeDB.SQLService)
	defer fakeDB.DeleteAll()

	t.Run("レコード一覧", func(t *testing.T) {
		defer fakeDB.DeleteAll()
		_, c := testPutItem(testEcho, testItemRequest())
		require.NoError(t, hdl.PutItemData(c))

		rec, c := testItemList()
		require.NoError(t, hdl.ListItems(c))
		assert.Equal(t, http.StatusOK, rec.Code)

		var res model.ItemList
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&res))
		require.Len(t, res, 1)
		assert.Equal(t, testTitle, res[0].Title)
		assert.Equal(t, testTag, res[0].Tag)
	})

	t.Run("レコード一覧#レコード0件", func(t *testing.T) {
		defer fakeDB.DeleteAll()

		rec, c := testItemList()
		require.NoError(t, hdl.ListItems(c))
		assert.Equal(t, http.StatusOK, rec.Code)

		var res model.ItemList
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&res))
		require.Equal(t, model.ItemList(nil), res)
	})
}
