package grid

import (
    "fmt"
    "math"

    "helpers"
    "structs"
)

func Build_grid(stops *map[int]structs.Stop)structs.Grid{
    grid := structs.NewGrid()

    d := 0
    for idx1, stop1 := range *stops{
        // find minimum distance between any two stops
        for idx2, stop2 := range *stops{
            if idx1 != idx2{
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


func InitGridGrid(grid *structs.Grid){
    for y:=0; y < grid.Y_size; y++{
        grid.Grid = append(grid.Grid, make([]*structs.Stop, grid.X_size, grid.X_size))
    }
}


func AssignStopsToGrid(stops *map[int]structs.Stop, grid *structs.Grid){
    var stop_grid_x, stop_grid_y int
    for _, stop := range *stops{
        stop_grid_x = (stop.X - grid.Min_X) / grid.Cell_size
        stop_grid_y = (stop.Y - grid.Min_Y) / grid.Cell_size
        grid.Grid[stop_grid_y][stop_grid_x] = &stop
    }
}
