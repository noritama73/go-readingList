package handler

import (
	"testing"

	"github.com/noritama73/go-readinglist/handler"
	"github.com/stretchr/testify/assert"
)

func Test_PutItemData(t *testing.T) {
	fakeDB := handler.NewFakeSQLService()
	defer fakeDB.DeleteAll()

	json_str := `{
		"title":"sample title",
		"memo":"sample memo",
		"url":"http://xxx.sample",
		"tag":"sample tag"
	}`
	assert.NoError(t, fakeDB.PutItemData([]byte(json_str)))
}
