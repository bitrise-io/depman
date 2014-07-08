package main

import (
	"fmt"
	"github.com/viktorbenei/depman/depman"
	"log"
)

func doUpdateWithDepList(deplist depman.DepList) {
	log.Printf("Updating %d dependencies...\n", len(deplist.Deps))
	deplocks, err := depman.PerformUpdateOnDepList(deplist)
	if err != nil {
		log.Fatal("Update failed: ", err)
	}
	// write DepLocks
	log.Println("Writing deplock...")
	if err := depman.WriteDepLocksToFile("./deplock.json", deplocks); err != nil {
		log.Fatal("Failed to write deplock.json: ", err)
	}
	log.Println("Update finished!")
}

func main() {
	fmt.Println("DepMan")

	deplist, err := depman.ReadDepListFromFile("./deplist.json")
	if err != nil {
		log.Fatal("Failed to load deplist: ", err)
	}

	doUpdateWithDepList(deplist)
}
