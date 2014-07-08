package depman

import (
	"encoding/json"
	"io"
	"os"
)

type DepStruct struct {
}

type DepList struct {
	Deps []DepStruct `json:"deps"`
}

func ReadDepListFromReader(reader io.Reader) (DepList, error) {
	var deplist DepList
	jsonParser := json.NewDecoder(reader)
	if err := jsonParser.Decode(&deplist); err != nil {
		return DepList{}, err
	}

	return deplist, nil
}

func ReadDepListFromFile(fpath string) (DepList, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return DepList{}, err
	}
	defer file.Close()

	return ReadDepListFromReader(file)
}
