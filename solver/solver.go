package solver

import "github.com/tscolari/dungeon_game/dungeon"

type Solver struct {
	dungeon  dungeon.Dungeon
	strategy Strategy
}

type Strategy interface {
	FindBestRoute(dungeon.Dungeon) (minHP int, bestRoute []string, err error)
}

func New(dungeon dungeon.Dungeon, strategy Strategy) *Solver {

	return &Solver{
		dungeon:  dungeon,
		strategy: strategy,
	}
}

func (s *Solver) Solve() (int, []string, error) {
	minHP, route, err := s.strategy.FindBestRoute(s.dungeon)
	if err != nil {
		return 0, []string{}, err
	}

	minHP = 1 - minHP
	return minHP, route, nil
}
