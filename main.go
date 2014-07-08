package main

import (
	"fmt"
	"github.com/viktorbenei/depman/depman"
	"log"
)

func doUpdateWithDepList(deplist depman.DepList) {
	log.Printf("Updating %d dependencies...", len(deplist.Deps))
	if err := depman.PerformUpdateOnDepList(deplist); err != nil {
		log.Fatal("Update failed: ", err)
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
