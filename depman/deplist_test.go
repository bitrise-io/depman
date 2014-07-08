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
		{}
	]
}`
	testDepList, err := ReadDepListFromReader(strings.NewReader(testDepListContent))
	if err != nil {
		t.Error(err)
	}
	if len(testDepList.Deps) != 1 {
		t.Error("Failed to read 'Deps'")
	}
}
