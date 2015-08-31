package strategy_test

import (
	"github.com/tscolari/dungeon_game/strategy"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Recursive", func() {
	testStrategy(&strategy.Recursive{})
})
