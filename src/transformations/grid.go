package grid

import (
    "fmt"
    "math"

    "helpers"
    "structs"
)

func Build_grid(stops map[int]structs.Stop)structs.Grid{
    // find minimum distance between any two stops
    min_distance := math.MaxInt32
    d := 0
    for idx1, stop1 := range stops{
        for idx2, stop2 := range stops{
            if idx1 != idx2{
                d = simple_math.Distance(stop1.X, stop2.X, stop1.Y, stop2.Y)
                min_distance = simple_math.Min(min_distance, d)
            }
        }
    }
    fmt.Println(min_distance)
    return structs.Grid{min_distance}
}
