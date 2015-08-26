package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tscolari/dungeon_game/dungeon"
)

func main() {
	var dungeonMapName = flag.String("dungeon", "", "Dungeon map file")
	flag.Parse()

	if *dungeonMapName == "" {
		flag.Usage()
		os.Exit(1)
	}

	_, err := dungeon.Load(*dungeonMapName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
