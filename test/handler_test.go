package handler

import (
	"encoding/json"
	"errors"
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
	testEcho       = echo.New()
	testTitle      = "test title"
	testMemo       = "test memo"
	testURL        = "http://xxx.test"
	testTag        = "test tag"
	updateTitle    = "update title"
	updateMemo     = "update memo"
	updateURL      = "http://xxx.update"
	updateTag      = "update tag"
	ErrSQLNoResult = errors.New("sql: no rows in result set")
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

func testUpdateItemRequest() string {
	return fmt.Sprintf(`
		{
			"title":"%s",
			"memo":"%s",
			"url":"%s",
			"tag":"%s"
		}
	`, updateTitle, updateMemo, updateURL, updateTag)
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

func testUpdateItem(e *echo.Echo, body string, id model.ID) (*httptest.ResponseRecorder, echo.Context) {
	f := make(url.Values)
	f.Set("id", fmt.Sprint(id))
	f.Set("data", body)
	req := httptest.NewRequest(echo.PUT, "/item", strings.NewReader(f.Encode()))
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

	_, c := testPutItem(testEcho, testItemRequest())
	require.NoError(t, hdl.PutItemData(c))

	t.Run("レコード更新", func(t *testing.T) {
		_, c := testUpdateItem(testEcho, testUpdateItemRequest(), model.ID(1))
		require.NoError(t, hdl.UpdateItemData(c))
	})

	t.Run("レコード更新#存在しないIDの更新", func(t *testing.T) {
		rec, c := testUpdateItem(testEcho, testUpdateItemRequest(), model.ID(9))
		assert.Equal(t, ErrSQLNoResult, hdl.UpdateItemData(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
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
