package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tscolari/dungeon_game/dungeon"
	"github.com/tscolari/dungeon_game/solver"
	"github.com/tscolari/dungeon_game/strategy"
)

func main() {
	var dungeonMapName = flag.String("dungeon", "", "Dungeon map file")
	flag.Parse()

	if *dungeonMapName == "" {
		flag.Usage()
		os.Exit(1)
	}

	dungeonMap, err := dungeon.Load(*dungeonMapName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	strategy := strategy.Recursive{}
	solver := solver.New(dungeonMap, &strategy)
	minHP, bestRoute, err := solver.Solve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("MIN HP: %d\n", minHP)
	fmt.Printf("BEST ROUTE: %s\n", strings.Join(bestRoute, " -> "))

}
