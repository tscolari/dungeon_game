package dijkstra

import (
	"errors"
	"math"

	"github.com/tscolari/dungeon_game/dungeon"
)

func NewRouteFinder(dungeon dungeon.Dungeon) *RouteFinder {
	return &RouteFinder{
		dungeon: dungeon,
	}
}

type RouteFinder struct {
	dungeon dungeon.Dungeon
}

func (f *RouteFinder) FindRoute(startX, startY, endX, endY int) (minHP int, routes []string, err error) {
	room := NewRoom(startX, startY)
	room.MinHP = f.dungeon.Rooms[startY][startX]
	room.TotalHP = f.dungeon.Rooms[startY][startX]
	room.Prev = nil
	roomsSet := f.initializeRoomsSetFor(room)

	for roomsSet.UnvisitedLen() > 0 {
		room := roomsSet.MaxMinHPRoom()
		roomsSet.Visit(room.X, room.Y)

		if room.X == endX && room.Y == endY {
			return f.findShortestPathFrom(room.X, room.Y, roomsSet)
		}

		for _, neighbor := range f.neighborsOf(room) {
			totalHP := room.TotalHP + f.dungeon.Rooms[neighbor[1]][neighbor[0]]
			minHP := min(room.MinHP, totalHP)

			neighborRoom, _ := roomsSet.Get(neighbor[0], neighbor[1])
			if (minHP > neighborRoom.MinHP) || (minHP == neighborRoom.MinHP && totalHP > neighborRoom.TotalHP) {

				neighborRoom.MinHP = minHP
				neighborRoom.TotalHP = totalHP
				neighborRoom.Prev = []int{room.X, room.Y}

				roomsSet.SaveRoom(neighborRoom)
			}
		}
	}
	return math.MinInt32, []string{}, errors.New("Could not find a path")
}

func (f *RouteFinder) initializeRoomsSetFor(room Room) Rooms {
	roomsSet := NewRooms()
	roomsSet.SaveRoom(room)

	for y := 0; y < len(f.dungeon.Rooms); y++ {
		for x := 0; x < len(f.dungeon.Rooms[0]); x++ {
			if x == room.X && y == room.Y {
				continue
			}
			room := NewRoom(x, y)
			room.MinHP = math.MinInt32
			room.TotalHP = 0
			room.Prev = nil
			roomsSet.SaveRoom(room)
		}
	}
	return roomsSet
}

func (f *RouteFinder) neighborsOf(room Room) [][]int {
	maxX := len(f.dungeon.Rooms[0]) - 1
	maxY := len(f.dungeon.Rooms) - 1

	neighbors := [][]int{}

	if room.X < maxX {
		neighbors = append(neighbors, []int{room.X + 1, room.Y})
	}

	if room.Y < maxY {
		neighbors = append(neighbors, []int{room.X, room.Y + 1})
	}

	return neighbors
}

func (f *RouteFinder) findShortestPathFrom(x, y int, roomsSet Rooms) (int, []string, error) {
	room, err := roomsSet.Get(x, y)
	if err != nil {
		return math.MinInt32, []string{}, err
	}

	routes := []string{}
	previous := room
	target, err := roomsSet.Get(room.Prev[0], room.Prev[1])
	if err != nil {
		return math.MinInt32, []string{}, err
	}

	for ; target.Prev != nil; target, _ = roomsSet.Get(target.Prev[0], target.Prev[1]) {
		direction := prevDirection(target, previous)
		routes = append([]string{direction}, routes...)
		previous = target
	}

	direction := prevDirection(target, previous)
	routes = append([]string{direction}, routes...)

	return room.MinHP, routes, nil
}