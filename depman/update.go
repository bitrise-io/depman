package depman

import (
	"errors"
	// "io"
	"os"
	"os/exec"
	"path/filepath"
	// "strings"
)

func updateDependency(dep DepStruct) (DepLockStruct, error) {
	cleanStorePath := filepath.Clean(dep.StorePath)

	absStorePath, err := filepath.Abs(cleanStorePath)
	if err != nil {
		return DepLockStruct{}, err
	}
	absStorePath = filepath.Clean(absStorePath)

	if absStorePath == cleanStorePath {
		return DepLockStruct{}, errors.New("Only relative paths allowed for StorePath!")
	}

	c := exec.Command("git", []string{"clone", dep.URL, absStorePath}...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		return DepLockStruct{}, err
	}
	if err := c.Wait(); err != nil {
		return DepLockStruct{}, err
	}

	deplock := DepLockStruct{URL: dep.URL, Revision: "x"}
	return deplock, nil
}

func PerformUpdateOnDepList(deplist DepList) error {
	deplocks := make([]DepLockStruct, len(deplist.Deps), len(deplist.Deps))
	for idx, aDep := range deplist.Deps {
		if aDepLock, err := updateDependency(aDep); err != nil {
			return err
		} else {
			deplocks[idx] = aDepLock
		}
	}
	return nil
}
