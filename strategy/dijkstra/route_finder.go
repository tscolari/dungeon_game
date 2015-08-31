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

/*
 * Dijstra algorithm for finding the best route between (startX, startY) and
 * (endX, endY). The concept of best route is the one with the higher minDamange
 * possible, and higher TotalHP (in case of 2 routes with the same minDamange).
 */
func (f *RouteFinder) FindRoute(startX, startY, endX, endY int) (minDamange int, routes []string, err error) {
	room := NewRoom(startX, startY)
	room.MinDamage = f.dungeon.Rooms[startY][startX]
	room.TotalHP = f.dungeon.Rooms[startY][startX]
	room.Prev = nil
	roomsSet := f.initializeRoomsSetFor(room)

	for roomsSet.UnvisitedLen() > 0 {
		room := roomsSet.MaxMinDamageRoom()
		roomsSet.Visit(room.X, room.Y)

		if room.X == endX && room.Y == endY {
			return f.findShortestPathFrom(room.X, room.Y, roomsSet)
		}

		for _, neighbor := range f.neighborsOf(room) {
			totalHP := room.TotalHP + f.dungeon.Rooms[neighbor[1]][neighbor[0]]
			minDamange := min(room.MinDamage, totalHP)

			neighborRoom, _ := roomsSet.Get(neighbor[0], neighbor[1])
			if (minDamange > neighborRoom.MinDamage) || (minDamange == neighborRoom.MinDamage && totalHP > neighborRoom.TotalHP) {

				neighborRoom.MinDamage = minDamange
				neighborRoom.TotalHP = totalHP
				neighborRoom.Prev = []int{room.X, room.Y}

				roomsSet.SaveRoom(neighborRoom)
			}
		}
	}
	return math.MinInt32, []string{}, errors.New("Could not find a path")
}

/*
 * Creates and return a new Rooms object, where the given room is already
 * correctly calculated and added to it.
 * All other rooms from the dungeon will be added with the slowest minDamange
 * possible as a mark of "unitialized"
 */
func (f *RouteFinder) initializeRoomsSetFor(room Room) Rooms {
	roomsSet := NewRooms()
	roomsSet.SaveRoom(room)

	for y := 0; y < len(f.dungeon.Rooms); y++ {
		for x := 0; x < len(f.dungeon.Rooms[0]); x++ {
			if x == room.X && y == room.Y {
				continue
			}
			room := NewRoom(x, y)
			room.MinDamage = math.MinInt32
			room.TotalHP = 0
			room.Prev = nil
			roomsSet.SaveRoom(room)
		}
	}
	return roomsSet
}

/*
 * Return the neighbor coordinates of the given room
 * It's aware of the dungeon borders and will not return an invalid point
 */
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

/*
 * This method suposes that roomSet is already computed and will return
 * the calculated route of the point (x,y).
 *
 * It will build the route from all the rooms that are linked by their
 * `Prev` field.
 *
 */
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

	return room.MinDamage, routes, nil
}
