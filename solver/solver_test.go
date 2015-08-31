package solver_test

import (
	"errors"

	"github.com/tscolari/dungeon_game/dungeon"
	"github.com/tscolari/dungeon_game/solver"
	"github.com/tscolari/dungeon_game/solver/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Solver", func() {
	var strategy *fakes.FakeStrategy
	var dungeonMap dungeon.Dungeon
	var subject *solver.Solver

	BeforeEach(func() {
		dungeonMap = dungeon.Dungeon{}
		strategy = new(fakes.FakeStrategy)
		subject = solver.New(dungeonMap, strategy)
	})

	Describe("#Solve", func() {
		It("converts the strategy value to actually minHP points", func() {
			strategy.FindBestRouteReturns(-10, []string{}, nil)

			minHP, _, _ := subject.Solve()
			Expect(minHP).To(Equal(11))
		})

		It("forwards the error if the strategy returns one", func() {
			expectedError := errors.New("something bad happened")
			strategy.FindBestRouteReturns(0, []string{}, expectedError)

			_, _, err := subject.Solve()
			Expect(err).To(MatchError("something bad happened"))
		})

		It("returns the same route returned from the strategy", func() {
			strategy.FindBestRouteReturns(0, []string{"RIGHT", "DOWN"}, nil)

			_, route, _ := subject.Solve()
			Expect(route).To(Equal([]string{"RIGHT", "DOWN"}))
		})
	})
})
