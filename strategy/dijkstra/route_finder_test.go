package dijkstra_test

import (
	"github.com/tscolari/dungeon_game/dungeon"
	"github.com/tscolari/dungeon_game/strategy/dijkstra"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RouteFinder", func() {
	var dungeonMap dungeon.Dungeon
	var finder *dijkstra.RouteFinder

	BeforeEach(func() {
		dungeonMap = dungeon.Dungeon{
			Rooms: [][]int{
				[]int{1, 2, -5, -4},
				[]int{-5, -6, -7, 9},
				[]int{9, 10, 11, 12},
			},
		}

		finder = dijkstra.NewRouteFinder(dungeonMap)
	})

	Describe("#FindBestRoute", func() {
		Context("from 0,0 to 3,2", func() {
			It("returns the optimal path", func() {
				_, path, err := finder.FindRoute(0, 0, 3, 2)
				Expect(err).ToNot(HaveOccurred())
				Expect(path).To(Equal([]string{"RIGHT", "DOWN", "DOWN", "RIGHT", "RIGHT"}))
			})

			It("returns the correct minDamange for the path", func() {
				minDamange, _, err := finder.FindRoute(0, 0, 3, 2)
				Expect(err).ToNot(HaveOccurred())
				Expect(minDamange).To(Equal(-3))
			})
		})

		Context("from 1,1 to 3,2", func() {
			It("returns the optimal path", func() {
				_, path, err := finder.FindRoute(1, 1, 3, 2)
				Expect(err).ToNot(HaveOccurred())
				Expect(path).To(Equal([]string{"DOWN", "RIGHT", "RIGHT"}))
			})

			It("returns the correct minDamange for the path", func() {
				minDamange, _, err := finder.FindRoute(1, 1, 3, 2)
				Expect(err).ToNot(HaveOccurred())
				Expect(minDamange).To(Equal(-6))
			})
		})

		Context("from 0,0 to 3,0", func() {
			It("returns the optimal path", func() {
				_, path, err := finder.FindRoute(0, 0, 3, 0)
				Expect(err).ToNot(HaveOccurred())
				Expect(path).To(Equal([]string{"RIGHT", "RIGHT", "RIGHT"}))
			})

			It("returns the correct minDamange for the path", func() {
				minDamange, _, err := finder.FindRoute(0, 0, 3, 0)
				Expect(err).ToNot(HaveOccurred())
				Expect(minDamange).To(Equal(-6))
			})
		})

		Context("from 0,0 to 0,2", func() {
			It("returns the optimal path", func() {
				_, path, err := finder.FindRoute(0, 0, 0, 2)
				Expect(err).ToNot(HaveOccurred())
				Expect(path).To(Equal([]string{"DOWN", "DOWN"}))
			})

			It("returns the correct minDamange for the path", func() {
				minDamange, _, err := finder.FindRoute(0, 0, 0, 2)
				Expect(err).ToNot(HaveOccurred())
				Expect(minDamange).To(Equal(-4))
			})
		})
	})

	Context("when there's a problem", func() {
		Context("when one of the coordinates is out of range", func() {
			It("returns an error", func() {
				_, _, err := finder.FindRoute(0, 0, 0, 3)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("Could not find a path"))
			})
		})
	})
})
