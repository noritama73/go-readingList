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
	res, err := fakeDB.GetItem(model.ID("1"))
	assert.NoError(t, err)
	assert.Equal(t, "sample title", res.Detail.Title)
	assert.Equal(t, "sample memo", res.Detail.Memo)
	assert.Equal(t, "http://xxx.sample", res.Detail.URL)
	assert.Equal(t, "sample tag", res.Detail.Tag)
}
