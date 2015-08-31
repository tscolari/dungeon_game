package strategy_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tscolari/dungeon_game/dungeon"
	"github.com/tscolari/dungeon_game/solver"
)

func testStrategy(strategy solver.Strategy) {
	var _ = Context(fmt.Sprintf("Behaves like a strategy (%T)", strategy), func() {

		type testSample struct {
			dungeon   dungeon.Dungeon
			minHP     int
			bestRoute []string
		}

		testSamples := []testSample{
			testSample{
				dungeon: dungeon.Dungeon{
					Rooms: [][]int{
						[]int{-2, -3, 3},
						[]int{-5, -10, 1},
						[]int{10, 30, -5},
					},
				},
				minHP:     -6,
				bestRoute: []string{"RIGHT", "RIGHT", "DOWN", "DOWN"},
			},
			testSample{
				dungeon: dungeon.Dungeon{
					Rooms: [][]int{
						[]int{1, 2, -5, -4},
						[]int{-5, -6, -7, 9},
						[]int{9, 10, 11, 12},
					},
				},
				minHP:     -3,
				bestRoute: []string{"RIGHT", "DOWN", "DOWN", "RIGHT", "RIGHT"},
			},
		}

		for i, testSample := range testSamples {
			sample := testSample

			Context(fmt.Sprintf("for test sample '%d'", i), func() {
				Describe("#FindBestRoute()", func() {
					It("returns the best route", func() {
						_, route, err := strategy.FindBestRoute(sample.dungeon)
						Expect(err).ToNot(HaveOccurred())
						Expect(route).To(Equal(sample.bestRoute))
					})

					It("returns the correct minimum HP needed for the best route", func() {
						hp, _, err := strategy.FindBestRoute(sample.dungeon)
						Expect(err).ToNot(HaveOccurred())
						Expect(hp).To(Equal(sample.minHP))
					})
				})
			})
		}
	})
}
