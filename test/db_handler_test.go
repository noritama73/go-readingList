package handler

import (
	"log"
	"testing"

	"github.com/noritama73/go-readinglist/handler"
	"github.com/noritama73/go-readinglist/model"
	"github.com/stretchr/testify/assert"
)

func Test_Account(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	fakeDB := handler.NewFakeSQLService()
	defer fakeDB.DeleteAll()

	json_str := `{
		"title":"sample title",
		"memo":"sample memo",
		"url":"http://xxx.sample",
		"tag":"sample tag"
	}`
	assert.NoError(t, fakeDB.PutItemData([]byte(json_str)))
	res, err := fakeDB.GetItem(model.ID(1))
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
	assert.NoError(t, fakeDB.UpdateItemData(model.ID(1), []byte(up_json_str)))
	res2, err := fakeDB.GetItem(model.ID(1))
	assert.NoError(t, err)
	assert.Equal(t, "update title", res2.Detail.Title)
	assert.Equal(t, "update memo", res2.Detail.Memo)
	assert.Equal(t, "http://xxx.update", res2.Detail.URL)
	assert.Equal(t, "update tag", res2.Detail.Tag)

	itemList, e := fakeDB.ListItems()
	assert.NoError(t, e)
	assert.Len(t, itemList, 1)
	assert.Equal(t, model.ID(1), itemList[0].ID)
	assert.Equal(t, "update title", itemList[0].Title)
	assert.Equal(t, "update tag", itemList[0].Tag)

	assert.NoError(t, fakeDB.DeleteItemData(model.ID(1)))
}
