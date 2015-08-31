package dijkstra

import (
	"errors"
	"fmt"
	"math"
)

func NewRooms() Rooms {
	return Rooms{
		unvisitedRooms: map[string]Room{},
		visitedRooms:   map[string]Room{},
	}
}

type Rooms struct {
	unvisitedRooms map[string]Room
	visitedRooms   map[string]Room
}

func (r *Rooms) UnvisitedLen() int {
	return len(r.unvisitedRooms)
}

func (r *Rooms) SaveRoom(room Room) {
	key := r.positionToKey(room.X, room.Y)
	r.unvisitedRooms[key] = room
}

func (r *Rooms) positionToKey(x, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}

func (r *Rooms) Get(x, y int) (Room, error) {
	key := r.positionToKey(x, y)
	room, ok := r.unvisitedRooms[key]
	if !ok {
		room, ok = r.visitedRooms[key]
		if !ok {
			return Room{}, errors.New("Room not found")
		}
	}

	return room, nil
}

func (r Rooms) MaxMinHPRoom() (maxMinHPRoom Room) {
	maxMinHP := math.MinInt32

	for _, room := range r.unvisitedRooms {
		if room.MinHP >= maxMinHP {
			maxMinHP = room.MinHP
			maxMinHPRoom = room
		}
	}

	return maxMinHPRoom
}

func (r *Rooms) Visit(x, y int) error {
	key := r.positionToKey(x, y)
	room, ok := r.unvisitedRooms[key]
	if !ok {
		return errors.New("Room not found")
	}
	delete(r.unvisitedRooms, key)
	r.visitedRooms[key] = room

	return nil
}
