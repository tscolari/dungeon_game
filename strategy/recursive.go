package strategy

import (
	"math"

	"github.com/tscolari/dungeon_game/dungeon"
)

type Recursive struct{}

func (s *Recursive) FindBestRoute(dungeon dungeon.Dungeon) (int, []string, error) {
	_, minHP, routes := s.findRouteFrom(0, 0, dungeon, 0, 0)

	return minHP, routes, nil
}

func (s *Recursive) findRouteFrom(posX, posY int, dungeon dungeon.Dungeon, currentHP, currentMinHP int) (int, int, []string) {
	maxX := len(dungeon.Rooms[0]) - 1
	maxY := len(dungeon.Rooms) - 1

	if posX > maxX || posY > maxY {
		return math.MinInt32, math.MinInt32, []string{}
	}

	totalHP := currentHP + dungeon.Rooms[posY][posX]
	minHP := s.minHPBetween(currentMinHP, totalHP)

	if posX == maxX && posY == maxY {
		return totalHP, minHP, []string{}
	}

	rightTotalHP, rightMinHP, rightRoutes := s.findRouteFrom(posX+1, posY, dungeon, totalHP, minHP)
	downTotalHP, downMinHP, downRoutes := s.findRouteFrom(posX, posY+1, dungeon, totalHP, minHP)

	if rightMinHP > downMinHP {

		return rightTotalHP, rightMinHP, append([]string{"RIGHT"}, rightRoutes...)
	}

	return downTotalHP, downMinHP, append([]string{"DOWN"}, downRoutes...)
}

func (s *Recursive) minHPBetween(hp1, hp2 int) int {
	if hp1 < hp2 {
		return hp1
	}

	return hp2
}
