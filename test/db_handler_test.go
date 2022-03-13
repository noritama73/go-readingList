package handler

import (
	"errors"
	"testing"

	"github.com/noritama73/go-readinglist/internal/handler"
	"github.com/noritama73/go-readinglist/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ErrSQLNoResult = errors.New("sql: no rows in result set")
)

func Test_Account(t *testing.T) {
	fakeDB := handler.NewFakeSQLService()
	defer fakeDB.DeleteAll()

	json_str := `{
		"title":"sample title",
		"memo":"sample memo",
		"url":"http://xxx.sample",
		"tag":"sample tag"
	}`
	assert.NoError(t, fakeDB.PutItemData([]byte(json_str)))
	id, res, err := fakeDB.GetItemTop()
	assert.NoError(t, err)
	assert.Equal(t, "sample title", res.Detail.Title)
	assert.Equal(t, "sample memo", res.Detail.Memo)
	assert.Equal(t, "http://xxx.sample", res.Detail.URL)
	assert.Equal(t, "sample tag", res.Detail.Tag)

	up_json_str := `{
		"title":"update title",
		"memo":"update memo",
		"url":"http://xxx.update",
		"tag":"update tag"
	}`
	assert.NoError(t, fakeDB.UpdateItemData(id, []byte(up_json_str)))
	_, res2, err := fakeDB.GetItemTop()
	assert.NoError(t, err)
	assert.Equal(t, "update title", res2.Detail.Title)
	assert.Equal(t, "update memo", res2.Detail.Memo)
	assert.Equal(t, "http://xxx.update", res2.Detail.URL)
	assert.Equal(t, "update tag", res2.Detail.Tag)

	itemList, e := fakeDB.ListItems()
	assert.NoError(t, e)
	require.Len(t, itemList, 1)
	assert.Equal(t, "update title", itemList[0].Title)
	assert.Equal(t, "update tag", itemList[0].Tag)

	assert.NoError(t, fakeDB.DeleteItemData(id))

	t.Run("存在しないレコードの更新", func(t *testing.T) {
		assert.Equal(t, ErrSQLNoResult, fakeDB.UpdateItemData(model.ID("hoge"), []byte(up_json_str)))
	})
}
