package dijkstra

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func prevDirection(roomSource, roomDestination Room) string {
	if roomSource.X < roomDestination.X {
		return "RIGHT"
	}

	return "DOWN"
}
