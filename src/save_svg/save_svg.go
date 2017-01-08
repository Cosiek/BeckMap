package save_svg

import (
    "fmt"
    "os"

    "third_party/svgo"

    "helpers"
    "structs"
)

var STOP_ICON_RANGE int = 5

func Save(stops map[int]structs.Stop, lines []structs.Line){
    // start new canvas
    f, err := os.Create("export.svg")
    if err != nil { panic(err) }
    defer f.Close()
    canvas := svg.New(f)
    // calculate canvas width and height
    max_x := 0
    max_y := 0
    for _, stop := range stops{
        max_x = simple_math.Max(max_x, stop.X)
        max_y = simple_math.Max(max_y, stop.Y)
    }

    canvas.Start(max_x, max_y)
    canvas.Grid(0, 0, max_x, max_y, 10, "stroke:black;opacity:0.05")
    for _, line := range lines{
        x_es := make([]int, 0)
        y_s := make([]int, 0)
        for _, stop := range line.Stops{
            x_es = append(x_es, stop.X)
            y_s = append(y_s, stop.Y)
        }
        canvas.Polyline(x_es, y_s, "stroke:black")
    }

    for _, stop := range stops{
        canvas.Circle(stop.X, stop.Y, STOP_ICON_RANGE)
    }
    canvas.End()

    fmt.Println("saved.")
}
