package depman

import (
	"strings"
	"testing"
)

func TestDepList(t *testing.T) {
	t.Log("Test DepList")

	_, err := ReadDepListFromReader(strings.NewReader(`{}`))
	if err != nil {
		t.Error("Empty test failed: ", err)
	}

	testDepListContent := `{
	"deps":[
		{
			"url": "test/url.url",
			"store_path": "a/relative/path"
		}
	]
}`
	testDepList, err := ReadDepListFromReader(strings.NewReader(testDepListContent))
	if err != nil {
		t.Error(err)
	}
	if len(testDepList.Deps) != 1 {
		t.Error("Failed to read 'Deps'")
	}
	testDepStruct := testDepList.Deps[0]
	if testDepStruct.URL != "test/url.url" {
		t.Error("Failed to read .URL")
	}
	if testDepStruct.StorePath != "a/relative/path" {
		t.Error("Failed to read .StorePath")
	}
}
