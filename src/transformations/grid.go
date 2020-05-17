package grid

import (
	//"fmt"
	"math"
	"sort"

	"structs"
)

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

func MoveLeft(byX []*structs.Stop) {
	// try pushing stops to the left, if they are in vertical relation to all
	// stops in previous column
	for idx, stopPtr := range byX {
		// first element can't be moved left
		if idx == 0 {
			continue
		}
		// get gridX of previous element
		refGridX := byX[idx-1].GridX
		canMoveLeft := true
		// iterate stops in previous column
		for idx2 := idx - 1; idx2 >= 0; idx2-- {
			// breake if reaching more then one column to the left
			if byX[idx2].GridX != refGridX {
				break
			}
			// check stops relative position
			_, dx := getDirectionDeltas(byX[idx], byX[idx2])
			if dx != 0 {
				canMoveLeft = false
				break
			}
		}
		if canMoveLeft {
			stopPtr.GridX = refGridX
		} else {
			stopPtr.GridX = refGridX + 1
		}
	}
}

func MoveUp(byY []*structs.Stop) {
	// try pushing stops up, if they are in horizontal relation to all
	// stops in previous row
	for idx, stopPtr := range byY {
		// first element can't be moved up
		if idx == 0 {
			continue
		}
		// get gridX of previous element
		refGridY := byY[idx-1].GridY
		canMoveUp := true
		// iterate stops in previous row
		for idx2 := idx - 1; idx2 >= 0; idx2-- {
			// breake if reaching more then one column up
			if byY[idx2].GridY != refGridY {
				break
			}
			// check stops relative position
			dy, _ := getDirectionDeltas(byY[idx], byY[idx2])
			if dy != 0 {
				canMoveUp = false
				break
			}
		}
		if canMoveUp {
			stopPtr.GridY = refGridY
		} else {
			stopPtr.GridY = refGridY + 1
		}
	}
}

func ApplyGrid(stopsPtr *map[int]*structs.Stop, lines_ *[]structs.Line) {
	stops := *stopsPtr
	// create slices of stops that will be ordered
	byX := make([]*structs.Stop, len(stops))
	byY := make([]*structs.Stop, len(stops))
	idx := 0
	for _, val := range stops {
		byX[idx] = val
		byY[idx] = val
		idx += 1
	}
	// apply sorting
	sort.Slice(byX, func(p, q int) bool { return byX[p].X < byX[q].X })
	sort.Slice(byY, func(p, q int) bool { return byY[p].Y < byY[q].Y })
	// update stops GridX and GridY attrs, basing on their position in table
	for idx, stopPtr := range byX {
		stopPtr.GridX = idx
	}
	for idx, stopPtr := range byY {
		stopPtr.GridY = idx
	}
	// condense grid by moving stops to rows/columns with others
	MoveLeft(byX)
	MoveUp(byY)
}
