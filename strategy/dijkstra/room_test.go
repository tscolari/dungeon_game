package dijkstra_test

import (
	"github.com/tscolari/dungeon_game/strategy/dijkstra"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	Describe("NewRoom", func() {
		It("initializes a Room object correctly", func() {
			room := dijkstra.NewRoom(100, 50)
			Expect(room.X).To(Equal(100))
			Expect(room.Y).To(Equal(50))
		})
	})
})
