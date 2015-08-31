package solver

import "github.com/tscolari/dungeon_game/dungeon"

type Solver struct {
	dungeon  dungeon.Dungeon
	strategy Strategy
}

type Strategy interface {
	FindBestRoute(dungeon.Dungeon) (minDamage int, bestRoute []string, err error)
}

func New(dungeon dungeon.Dungeon, strategy Strategy) *Solver {
	return &Solver{
		dungeon:  dungeon,
		strategy: strategy,
	}
}

func (s *Solver) Solve() (int, []string, error) {
	minDamage, route, err := s.strategy.FindBestRoute(s.dungeon)
	if err != nil {
		return 0, []string{}, err
	}

	minHP := 1 - minDamage
	return minHP, route, nil
}
