package grid

import (
    "fmt"
    "math"

    "helpers"
    "structs"
)

func Build_grid(stops map[int]structs.Stop)structs.Grid{
    grid := structs.NewGrid()

    d := 0
    for idx1, stop1 := range stops{
        // find minimum distance between any two stops
        for idx2, stop2 := range stops{
            if idx1 != idx2{
                d = simple_math.Distance(stop1.X, stop2.X, stop1.Y, stop2.Y)
                grid.Cell_size = simple_math.Min(grid.Cell_size, d)
            }
        }

        // find utmost coordinates
        grid.Min_X = simple_math.Min(grid.Min_X, stop1.X)
        grid.Max_X = simple_math.Max(grid.Max_X, stop1.X)
        grid.Min_Y = simple_math.Min(grid.Min_Y, stop1.Y)
        grid.Max_Y = simple_math.Max(grid.Max_Y, stop1.Y)
    }

    // calculate grid size (cell count)
    delta_x := grid.Max_X - grid.Min_X
    delta_y := grid.Max_Y - grid.Min_Y
    fmt.Println(grid.Max_X, grid.Min_X, grid.Max_Y, grid.Min_Y)
    fmt.Println(delta_x, delta_y, grid.Cell_size)
    grid.X_size = int(math.Ceil(float64(delta_x) / float64(grid.Cell_size)))
    grid.Y_size = int(math.Ceil(float64(delta_y) / float64(grid.Cell_size)))
    fmt.Println(grid.X_size, grid.Y_size)

    return grid
}
