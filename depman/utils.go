package depman

import (
	"fmt"
)

// ReadDepListFile ...
func ReadDepListFile() (DepList, error) {
	deplist, err := ReadDepListFromFile("./deplist.json")
	if err != nil {
		return DepList{}, fmt.Errorf("Failed to load deplist: %s", err)
	}
	return deplist, nil
}
