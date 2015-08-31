package strategy_test

import (
	"github.com/tscolari/dungeon_game/strategy"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Dijkstra", func() {
	testStrategy(&strategy.Dijkstra{})
})
