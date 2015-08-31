package strategy

import (
	"github.com/tscolari/dungeon_game/dungeon"
	"github.com/tscolari/dungeon_game/strategy/dijkstra"
)

type Dijkstra struct {
}

func (d *Dijkstra) FindBestRoute(dungeon dungeon.Dungeon) (int, []string, error) {
	finder := dijkstra.NewRouteFinder(dungeon)
	maxX := len(dungeon.Rooms[0]) - 1
	maxY := len(dungeon.Rooms) - 1
	return finder.FindRoute(0, 0, maxX, maxY)
}
