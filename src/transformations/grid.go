package grid

import (
	"fmt"
	"math"

	"helpers"
	"structs"
)

func Build_grid(stops *map[int]structs.Stop) structs.Grid {
	grid := structs.NewGrid()

	d := 0
	for idx1, stop1 := range *stops {
		// find minimum distance between any two stops
		for idx2, stop2 := range *stops {
			if idx1 != idx2 {
				d = helpers.Distance(stop1.X, stop2.X, stop1.Y, stop2.Y)
				grid.Cell_size = helpers.Min(grid.Cell_size, d)
			}
		}

		// find utmost coordinates
		grid.Min_X = helpers.Min(grid.Min_X, stop1.X)
		grid.Max_X = helpers.Max(grid.Max_X, stop1.X)
		grid.Min_Y = helpers.Min(grid.Min_Y, stop1.Y)
		grid.Max_Y = helpers.Max(grid.Max_Y, stop1.Y)
	}

	// calculate grid size (cell count)
	delta_x := grid.Max_X - grid.Min_X
	delta_y := grid.Max_Y - grid.Min_Y
	grid.X_size = int(math.Ceil(float64(delta_x) / float64(grid.Cell_size)))
	grid.Y_size = int(math.Ceil(float64(delta_y) / float64(grid.Cell_size)))

	// assign stops to grid cells
	//grid.Grid = [grid.Y_size][grid.X_size]*structs.Stop
	InitGridGrid(&grid)
	fmt.Println("jest")
	AssignStopsToGrid(stops, &grid)

	return grid
}

func InitGridGrid(grid *structs.Grid) {
	for y := 0; y < grid.Y_size; y++ {
		grid.Grid = append(grid.Grid, make([]*structs.Stop, grid.X_size, grid.X_size))
	}
}

func AssignStopsToGrid(stops *map[int]structs.Stop, grid *structs.Grid) {
	var stop_grid_x, stop_grid_y int
	for _, stop := range *stops {
		stop_grid_x = (stop.X - grid.Min_X) / grid.Cell_size
		stop_grid_y = (stop.Y - grid.Min_Y) / grid.Cell_size
		grid.Grid[stop_grid_y][stop_grid_x] = &stop
	}
}

func GetIndex(stop *structs.Stop, stops *[]*structs.Stop) int {
	for idx, listItem := range *stops {
		if *stop == *listItem {
			return idx
		}
	}
	return -1
}

func getDirectionDeltas(prev *structs.Stop, curr *structs.Stop) (int, int) {
	/* Returns grid X and Y deltas based on stops relative positions */
	return 1, 0
}

func markStops(linePtr *structs.Line, idx int, step int) {
	line := *linePtr
	lastStop := line.Stops[idx]

	for i := idx + step; i > 0 && i < len(line.Stops); i += step {
		currentStop := line.Stops[i]
		// if stop is already in grid
		if currentStop.GridX > -1 {
			// check if positions need to be adjusted
		} else {
			// set stops grid positions with regard to previous one
			dX, dY := getDirectionDeltas(lastStop, currentStop)
			currentStop.GridX = lastStop.GridX + dX
			currentStop.GridY = lastStop.GridY + dY
		}
		lastStop = currentStop
		fmt.Println(lastStop.Name, lastStop.GridX, lastStop.GridY)
	}
}

func BuildGrid(stops_ *map[int]structs.Stop, lines_ *[]structs.Line) {
	//stops := *stops_
	lines := *lines_
	var line structs.Line
	// set first stop at 0, 0 grid position
	stop := lines[0].Stops[0]
	stop.GridX = 0
	stop.GridY = 0
	for {
		// select a line that has a stop that is already in the grid
		idx := 0
		for _, l := range lines {
			idx = GetIndex(stop, &l.Stops)
			if idx > -1 {
				line = l
				break
			}
		}
		fmt.Println(line)
		fmt.Println(idx)
		// iterate up starting from that common stop
		markStops(&line, idx, 1)
		// return to common stop and iterate the other way
		markStops(&line, idx, -1)
		break
	}
}
