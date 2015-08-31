package dijkstra

type Room struct {
	X       int
	Y       int
	MinHP   int
	TotalHP int
	Prev    []int
}

func NewRoom(x, y int) Room {
	return Room{
		X: x,
		Y: y,
	}
}
