package strategy

import (
	"math"

	"github.com/tscolari/dungeon_game/dungeon"
)

type Recursive struct{}

func (s *Recursive) FindBestRoute(dungeon dungeon.Dungeon) (int, []string, error) {
	_, minDamange, routes := s.findRouteFrom(0, 0, dungeon, 0, 0)

	return minDamange, routes, nil
}

func (s *Recursive) findRouteFrom(posX, posY int, dungeon dungeon.Dungeon, currentHP, currentMinDamage int) (int, int, []string) {
	maxX := len(dungeon.Rooms[0]) - 1
	maxY := len(dungeon.Rooms) - 1

	if posX > maxX || posY > maxY {
		return math.MinInt32, math.MinInt32, []string{}
	}

	totalHP := currentHP + dungeon.Rooms[posY][posX]
	minDamange := s.minDamangeBetween(currentMinDamage, totalHP)

	if posX == maxX && posY == maxY {
		return totalHP, minDamange, []string{}
	}

	rightTotalHP, rightMinDamage, rightRoutes := s.findRouteFrom(posX+1, posY, dungeon, totalHP, minDamange)
	downTotalHP, downMinDamage, downRoutes := s.findRouteFrom(posX, posY+1, dungeon, totalHP, minDamange)

	if rightMinDamage > downMinDamage {

		return rightTotalHP, rightMinDamage, append([]string{"RIGHT"}, rightRoutes...)
	}

	return downTotalHP, downMinDamage, append([]string{"DOWN"}, downRoutes...)
}

func (s *Recursive) minDamangeBetween(hp1, hp2 int) int {
	if hp1 < hp2 {
		return hp1
	}

	return hp2
}
