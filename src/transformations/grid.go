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

func GetCommonStopIdx(visitedStopsPtr *map[*structs.Stop]bool, linePtr *structs.Line) int {
	visitedStops := *visitedStopsPtr
	for idx, listItemPtr := range linePtr.Stops {
		_, ok := visitedStops[listItemPtr]
		if ok {
			return idx
		}
	}
	return -1
}

func getDirectionDeltas(prev *structs.Stop, curr *structs.Stop) (int, int) {
	/* Returns grid X and Y deltas based on stops relative positions */
	angle := math.Atan2(float64(curr.Y-prev.Y), float64(curr.X-prev.X))
	var x, y int

	if angle >= 0 {
		if angle < math.Pi/8 {
			y, x = 0, 1
		} else if angle < math.Pi*3/8 {
			y, x = 1, 1
		} else if angle < math.Pi*5/8 {
			y, x = 1, 0
		} else if angle < math.Pi*7/8 {
			y, x = 1, -1
		} else {
			y, x = 0, -1
		}
	} else {
		if angle > -math.Pi/8 {
			y, x = 0, 1
		} else if angle > -math.Pi*3/8 {
			y, x = -1, 1
		} else if angle > -math.Pi*5/8 {
			y, x = -1, 0
		} else if angle > -math.Pi*7/8 {
			y, x = -1, -1
		} else {
			y, x = 0, -1
		}
	}
	return y, x
}

func adjustStopsPositions(visitedStops *map[*structs.Stop]bool, currY, currX, dY, dX int) {
	for stop, _ := range *visitedStops {
		if stop.GridY >= currY {
			stop.GridY += dY
		}
		if stop.GridX >= currX {
			stop.GridX += dX
		}
	}
}

func markStops(linePtr *structs.Line, visitedStopsPtr *map[*structs.Stop]bool, idx int, step int) {
	line := *linePtr
	visitedStops := *visitedStopsPtr
	lastStop := line.Stops[idx]

	for i := idx + step; i >= 0 && i < len(line.Stops); i += step {
		currentStop := line.Stops[i]
		dY, dX := getDirectionDeltas(lastStop, currentStop)
		// if stop is already in grid
		if visitedStops[currentStop] {
			// check if positions need to be adjusted on Y axis
			if currentStop.GridY-lastStop.GridY*dY < 0 {
				dY = int(math.Abs(float64(currentStop.GridY - lastStop.GridY + dY)))
			} else {
				dY = 0
			}

			// check if positions need to be adjusted on X axis
			if currentStop.GridX-lastStop.GridX*dX < 0 {
				dX = int(math.Abs(float64(currentStop.GridX - lastStop.GridX + dX)))
			} else {
				dX = 0
			}

			adjustStopsPositions(visitedStopsPtr, currentStop.GridY, currentStop.GridX, dY, dX)
		} else {
			// set stops grid positions with regard to previous one
			currentStop.GridX = lastStop.GridX + dX
			currentStop.GridY = lastStop.GridY + dY
			visitedStops[currentStop] = true
		}
		lastStop = currentStop
	}
}

func BuildGrid(stopsPtr *map[int]*structs.Stop, lines_ *[]structs.Line) {
	lines := *lines_
	var line structs.Line
	// prepare mapings
	var visitedStops = make(map[*structs.Stop]bool)
	var markedLines = make(map[int]bool)
	// set first stop at 0, 0 grid position
	stop := lines[0].Stops[0]
	stop.GridX = 0
	stop.GridY = 0
	visitedStops[stop] = true

	for {
		// select a line that has a stop that is already in the grid
		idx := -1
		for _, l := range lines {
			if markedLines[l.Id] {
				continue
			}
			// mark this as a selected one and continue...
			line = l
			// to check if this one has any stops that have been already marked
			idx = GetCommonStopIdx(&visitedStops, &l)
			if idx > -1 {
				break
			}
		}
		// if no such line was found, then its time to leave
		// NOTE: this will fail if any of the lines is detached from the rest
		if idx == -1 {
			break
		}
		// iterate up starting from that common stop
		markStops(&line, &visitedStops, idx, 1)
		// return to common stop and iterate the other way
		markStops(&line, &visitedStops, idx, -1)

		markedLines[line.Id] = true
	}
}
