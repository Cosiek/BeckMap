package save_svg

import (
	"fmt"
	"math"
	"os"

	"github.com/ajstarks/svgo"

	"helpers"
	"structs"
)

var STOP_ICON_RANGE int = 5

func Save(stops map[int]structs.Stop, lines []structs.Line) {
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
	canvas.Start(max_x, max_y)
	canvas.Grid(0, 0, max_x, max_y, 2*STOP_ICON_RANGE, "stroke:black;opacity:0.05")
	for _, line := range lines {
		x_es := make([]int, 0)
		y_s := make([]int, 0)
		for _, stop := range line.Stops {
			x_es = append(x_es, stop.X-offset_x)
			y_s = append(y_s, stop.Y-offset_y)
		}
		canvas.Polyline(x_es, y_s, "stroke:black;fill:none")
	}

	for _, stop := range stops {
		canvas.Circle(stop.X-offset_x, stop.Y-offset_y, STOP_ICON_RANGE)
	}
	canvas.End()

	fmt.Println("saved.")
}
