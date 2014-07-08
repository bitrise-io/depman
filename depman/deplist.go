package depman

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type DepLockStruct struct {
	URL      string `json:"url"`
	Revision string `json:"revision"`
}

type depLocksStruct struct {
	DepLocks []DepLockStruct `json:"dep_locks"`
}

type DepStruct struct {
	URL       string `json:"url"`
	StorePath string `json:"store_path"`
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

func generateFormattedJSONForDepLocks(deplocks []DepLockStruct) ([]byte, error) {
	jsonContBytes, err := json.MarshalIndent(depLocksStruct{DepLocks: deplocks}, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

func WriteDepLocksToFile(fpath string, deplocks []DepLockStruct) error {
	if fpath == "" {
		return errors.New("No path provided!")
	}

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonContBytes, err := generateFormattedJSONForDepLocks(deplocks)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonContBytes)
	if err != nil {
		return err
	}

	return nil
}
