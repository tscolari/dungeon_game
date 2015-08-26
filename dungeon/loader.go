package dungeon

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(filePath string) (Dungeon, error) {
	dungeonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Dungeon{}, errors.New("Failed to open dungeon file: " + err.Error())
	}

	var dungeon Dungeon
	err = yaml.Unmarshal(dungeonData, &dungeon)
	if err != nil {
		return Dungeon{}, errors.New("Invalid dungeon file: Not a valid yaml")
	}

	if !validateIfRoomsHaveSameLenght(dungeon) {
		return Dungeon{}, errors.New("Invalid dungeon file: All rows must have the same size")
	}

	if !validateMinimumRooms(dungeon) {
		return Dungeon{}, errors.New("Invalid dungeon file: Not enough rooms (min 2)")
	}

	return dungeon, nil
}

func validateIfRoomsHaveSameLenght(dungeon Dungeon) bool {
	var width = len(dungeon.Rooms[0])

	for _, row := range dungeon.Rooms[1:] {
		if width != len(row) {
			return false
		}
	}

	return true
}

func validateMinimumRooms(dungeon Dungeon) bool {
	if len(dungeon.Rooms) == 0 {
		return false
	}

	if len(dungeon.Rooms) >= 2 {
		return true
	}

	if len(dungeon.Rooms[0]) >= 2 {
		return true
	}

	return false
}
