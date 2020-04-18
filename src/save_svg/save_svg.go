package save_svg

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/ajstarks/svgo"

	"helpers"
	"structs"
)

var STOP_ICON_RANGE int = 5

func SaveMap(stops map[int]*structs.Stop, lines []structs.Line) {
	// start new canvas
	f, err := os.Create("export.svg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	canvas := svg.New(f)
	// calculate canvas width and height
	max_x := 0
	max_y := 0
	offset_x := math.MaxInt32
	offset_y := math.MaxInt32
	for _, stop := range stops {
		max_x = helpers.Max(max_x, stop.X)
		offset_x = helpers.Min(offset_x, stop.X)
		max_y = helpers.Max(max_y, stop.Y)
		offset_y = helpers.Min(offset_y, stop.Y)
	}

	offset_x -= 2 * STOP_ICON_RANGE
	offset_y -= 2 * STOP_ICON_RANGE
	max_x += 2 * STOP_ICON_RANGE
	max_y += 2 * STOP_ICON_RANGE
	// draw grid
	canvas.Start(max_x, max_y)
	canvas.Grid(0, 0, max_x, max_y, 2*STOP_ICON_RANGE, "stroke:black;opacity:0.05")
	// draw lines
	for _, line := range lines {
		x_es := make([]int, 0)
		y_s := make([]int, 0)
		for _, stop := range line.Stops {
			x_es = append(x_es, stop.X-offset_x)
			y_s = append(y_s, stop.Y-offset_y)
		}
		canvas.Polyline(x_es, y_s, "stroke:black;fill:none")
	}
	// draw stops
	for _, stop := range stops {
		canvas.Circle(stop.X-offset_x, stop.Y-offset_y, STOP_ICON_RANGE)
		canvas.Text(stop.X-offset_x+5, stop.Y-offset_y, strconv.Itoa(stop.Id), "font:10px serif;stroke:black;fill:none")
	}
	// finish
	canvas.End()
	fmt.Println("saved.")
}

func SaveGrid(stopsPtr *map[int]*structs.Stop, linesPtr *[]structs.Line) {
	stops := *stopsPtr
	lines := *linesPtr

	// normalize all stops not to have negative X, Y positions
	min_x, min_y := float64(math.MaxInt32), float64(math.MaxInt32)
	for _, stop := range stops {
		min_y = math.Min(min_y, float64(stop.GridY))
		min_x = math.Min(min_x, float64(stop.GridX))
	}
	min_y = 0 - min_y
	min_x = 0 - min_x
	for _, stop := range stops {
		stop.GridY += int(min_y)
		stop.GridX += int(min_x)
	}

	// calculate canvas width and height
	max_x := 0
	max_y := 0
	offset_x := math.MaxInt32
	offset_y := math.MaxInt32
	for _, stop := range stops {
		max_x = helpers.Max(max_x, stop.GridX*10)
		offset_x = helpers.Min(offset_x, stop.GridX*10)
		max_y = helpers.Max(max_y, stop.GridY*10)
		offset_y = helpers.Min(offset_y, stop.GridY*10)
	}

	offset_x -= 2 * STOP_ICON_RANGE
	offset_y -= 2 * STOP_ICON_RANGE
	max_x += 3 * STOP_ICON_RANGE
	max_y += 3 * STOP_ICON_RANGE

	// start new canvas
	f, err := os.Create("export_grid.svg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	canvas := svg.New(f)

	// draw grid
	canvas.Start(max_x, max_y)
	canvas.Grid(0, 0, max_x, max_y, 2*STOP_ICON_RANGE, "stroke:black;opacity:0.05")

	// draw lines
	for _, line := range lines {
		x_es := make([]int, 0)
		y_s := make([]int, 0)
		for _, stop := range line.Stops {
			x_es = append(x_es, stop.GridX*10-offset_x)
			y_s = append(y_s, stop.GridY*10-offset_y)
		}
		canvas.Polyline(x_es, y_s, "stroke:red;fill:none")
	}
	// draw stops
	for _, stop := range stops {
		fmt.Println(stop.Name, stop.GridY, stop.GridY)
		canvas.Circle(stop.GridX*10-offset_x, stop.GridY*10-offset_y, 3)
		canvas.Text(stop.GridX*10-offset_x, stop.GridY*10-offset_y, strconv.Itoa(stop.Id), "font:8px serif;stroke:green;fill:none")
	}
	// finish
	canvas.End()
	fmt.Println("saved.")

}
