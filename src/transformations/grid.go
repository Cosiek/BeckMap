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

func MoveRight(byX []*structs.Stop) {
	// get right-most GridX
	idx := len(byX) - 1
	refGridX := byX[idx].GridX
	// get index of last stop from second to last column
	for byX[idx].GridX == refGridX {
		idx--
	}
	// iterate stops from right to left
	for idx := len(byX) - 1; idx >= 0; idx-- {
		// stops in last column can't be moved
		if byX[idx].GridX == refGridX {
			continue
		}
		// if there is another stop in place where stop would be moved, then
		// there is nothing that can be done.
		for idx2 := idx + 1; idx2 < len(byX); idx2++ {
			if byX[idx2].GridX == byX[idx].GridX+1 && byX[idx2].GridY == byX[idx].GridY {
				break // ŹLE - trzeba się wybić z zewnętrznej petli
			} else if byX[idx2].GridX > byX[idx].GridX+1 {
				// no need to check further then one column to the right
				break
			}
		}
		// calculate avrage X of this column (except current stop) and
		// avrage X of next colunm
		var thisColSum, thisColCount float64 = 0, 0
		var nextColSum, nextColCount float64 = 0, 0
		for idx2 := 0; idx2 < len(byX); idx2++ {
			// skip stops from previous columns
			if byX[idx].GridX > byX[idx2].GridX {
				continue
			}
			// skip current stop
			if idx == idx2 {
				continue
			}
			if byX[idx].GridX == byX[idx2].GridX {
				// update sum and count for this column
				thisColSum += float64(byX[idx2].X)
				thisColCount++
			} else if byX[idx2].GridX == byX[idx].GridX+1 {
				// update sum and count for next column
				nextColSum += float64(byX[idx2].X)
				nextColCount++
			} else {
				break
			}
		}
		// compare these avrages with current stops X
		dToThisAvg := float64(byX[idx].X) - thisColSum/thisColCount
		dFromNextAvg := nextColSum/nextColCount - float64(byX[idx].X)
		if dToThisAvg > 0 && dFromNextAvg < dToThisAvg {
			byX[idx].GridX++
		}
	}
}

func MoveDown(byY []*structs.Stop) {
	// get down-most GridY
	idx := len(byY) - 1
	refGridY := byY[idx].GridY
	// get index of last stop from second to last row
	for byY[idx].GridY == refGridY {
		idx--
	}
	// iterate stops from bottom to top
	for idx := len(byY) - 1; idx >= 0; idx-- {
		// stops in last row can't be moved
		if byY[idx].GridY == refGridY {
			continue
		}
		// if there is another stop in place where stop would be moved, then
		// there is nothing that can be done.
		for idx2 := idx + 1; idx2 < len(byY); idx2++ {
			if byY[idx2].GridY == byY[idx].GridY+1 && byY[idx2].GridX == byY[idx].GridX {
				break // TODO - trzeba się wybić z zewnętrznej petli
			} else if byY[idx2].GridY > byY[idx].GridY+1 {
				// no need to check further then one row up
				break
			}
		}
		// calculate avrage Y of this row (except current stop) and
		// avrage Y of next row
		var thisRowSum, thisRowCount float64 = 0, 0
		var nextRowSum, nextRowCount float64 = 0, 0
		for idx2 := 0; idx2 < len(byY); idx2++ {
			// skip stops from previous rows
			if byY[idx].GridY > byY[idx2].GridY {
				continue
			}
			// skip current stop
			if idx == idx2 {
				continue
			}
			if byY[idx].GridY == byY[idx2].GridY {
				// update sum and count for this row
				thisRowSum += float64(byY[idx2].Y)
				thisRowCount++
			} else if byY[idx2].GridY == byY[idx].GridY+1 {
				// update sum and count for next row
				nextRowSum += float64(byY[idx2].Y)
				nextRowCount++
			} else {
				break
			}
		}
		// compare these avrages with current stops Y
		dToThisAvg := float64(byY[idx].Y) - thisRowSum/thisRowCount
		dFromNextAvg := nextRowSum/nextRowCount - float64(byY[idx].Y)
		if dToThisAvg > 0 && dFromNextAvg < dToThisAvg {
			byY[idx].GridY++
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
	MoveRight(byX)
	MoveDown(byY)
	// TODO: Get rid of empty rows/columns
}
