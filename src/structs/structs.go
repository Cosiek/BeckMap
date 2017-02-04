package structs

import (
    "math"
)


type Stop struct {
    Id int
    Name string
    X, Y int    // coordinates
}


type Line struct {
    Id int
    Name string
    Stops []*Stop
}


type Grid struct {
    Cell_size int

    X_size, Y_size int

    Min_X, Max_X int
    Min_Y, Max_Y int
}

func NewGrid() Grid{
    grid := Grid{}

    grid.Cell_size = math.MaxInt32
    grid.Min_X = math.MaxInt32
    grid.Max_X = -1 * math.MaxInt32
    grid.Min_Y = math.MaxInt32
    grid.Max_Y = -1 * math.MaxInt32
    return grid
}
