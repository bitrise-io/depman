package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/viktorbenei/depman/depman"
	"log"
	"os"
)

type Command struct {
	Name  string
	Usage string
	Run   func() error
}

var (
	availableCommands = []Command{
		Command{
			Name:  "update",
			Usage: "update - updates",
			Run:   doUpdateCommand,
		},
		Command{
			Name:  "init",
			Usage: "init - creates a deplist.json file in the current folder",
			Run:   doInitCommand,
		},
	}
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s command [FLAGS]", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("Available commands:")
	for _, cmd := range availableCommands {
		fmt.Println(" *", cmd.Name)
		fmt.Println("    ", cmd.Usage)
	}
}

func readDepListFile() (depman.DepList, error) {
	deplist, err := depman.ReadDepListFromFile("./deplist.json")
	if err != nil {
		return depman.DepList{}, errors.New(fmt.Sprintf("Failed to load deplist: %s", err))
	}
	return deplist, nil
}

func doUpdateCommand() error {
	deplist, err := readDepListFile()
	if err != nil {
		return err
	}

	log.Printf("Updating %d dependencies...\n", len(deplist.Deps))
	deplocks, err := depman.PerformUpdateOnDepList(deplist)
	if err != nil {
		return errors.New(fmt.Sprintf("Update failed: %s", err))
	}
	// write DepLocks
	log.Println("Writing deplock...")
	if err := depman.WriteDepLocksToFile("./deplock.json", deplocks); err != nil {
		return errors.New(fmt.Sprintf("Failed to write deplock.json: %s", err))
	}
	log.Println("Update finished!")
	return nil
}

func doInitCommand() error {
	deplist := depman.DepList{
		Deps: []depman.DepStruct{
			depman.DepStruct{
				URL:       "http://repo.url",
				StorePath: "relative/store/path",
			},
		},
	}
	err := depman.WriteDepListToFile("./deplist.json", deplist)
	if err != nil {
		return err
	}
	fmt.Println("deplist.json file saved")
	return nil
}

func main() {
	fmt.Println("DepMan")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	// fmt.Println(args)
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	theCommandName := args[0]

	for _, cmd := range availableCommands {
		if cmd.Name == theCommandName {
			// cmd.Flag.Usage = func() { cmd.UsageExit() }
			// cmd.Flag.Parse(args[1:])
			err := cmd.Run()
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
	}
}
