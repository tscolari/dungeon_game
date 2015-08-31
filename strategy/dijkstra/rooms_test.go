package dijkstra_test

import (
	"github.com/tscolari/dungeon_game/strategy/dijkstra"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rooms", func() {
	var room1, room2, room3 dijkstra.Room
	var rooms dijkstra.Rooms

	BeforeEach(func() {
		rooms = dijkstra.NewRooms()

		room1 = dijkstra.Room{X: 1, Y: 2, MinDamage: -1000}
		room2 = dijkstra.Room{X: 2, Y: 1, MinDamage: 100}
		room3 = dijkstra.Room{X: 0, Y: 0, MinDamage: 0}
		rooms.SaveRoom(room1)
		rooms.SaveRoom(room2)
		rooms.SaveRoom(room3)
	})

	Describe("#MaxMinDamageRoom", func() {
		It("returns the index of the Room with the highest `MinDamage`", func() {
			maxMinDamageRoom := rooms.MaxMinDamageRoom()
			Expect(maxMinDamageRoom).To(Equal(room2))
		})

		Context("visited rooms", func() {
			It("ignores visited rooms", func() {
				room := dijkstra.Room{X: 10, Y: 10, MinDamage: 1000}
				rooms.SaveRoom(room)
				rooms.Visit(10, 10)
				maxMinDamageRoom := rooms.MaxMinDamageRoom()
				Expect(maxMinDamageRoom).ToNot(Equal(room))
			})
		})
	})

	Describe("#Get", func() {
		It("returns the correct room based on X,Y", func() {
			room, err := rooms.Get(2, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(room).To(Equal(room2))
		})

		Context("when it doesn't exist", func() {
			It("returns a not found error", func() {
				_, err := rooms.Get(2, 10)
				Expect(err).To(MatchError("Room not found"))
			})
		})
	})

	Describe("#Visit", func() {
		It("changes the UnvisitedLen", func() {
			Expect(rooms.UnvisitedLen()).To(Equal(3))
			err := rooms.Visit(1, 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(rooms.UnvisitedLen()).To(Equal(2))
		})

		Context("when it doesn't exist", func() {
			It("returns a not found error", func() {
				err := rooms.Visit(2, 10)
				Expect(err).To(MatchError("Room not found"))
			})
		})
	})

	Describe("#SaveRoom", func() {
		It("adds the given room", func() {
			newRoom := dijkstra.Room{X: 5, Y: 5, MinDamage: 10}
			_, err := rooms.Get(5, 5)
			Expect(err).To(MatchError("Room not found"))

			rooms.SaveRoom(newRoom)
			room, err := rooms.Get(5, 5)
			Expect(err).ToNot(HaveOccurred())
			Expect(room).To(Equal(newRoom))
		})

		Context("when a room with same X,Y already exists", func() {
			It("updates the room object with the new values", func() {
				oldRoom := dijkstra.Room{X: 5, Y: 5, MinDamage: 10}
				rooms.SaveRoom(oldRoom)

				newRoom := dijkstra.Room{X: 5, Y: 5, MinDamage: 100}
				rooms.SaveRoom(newRoom)

				room, err := rooms.Get(5, 5)
				Expect(err).ToNot(HaveOccurred())

				Expect(room.MinDamage).To(Equal(100))
			})
		})
	})
})
