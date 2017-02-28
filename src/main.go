package main

import (
    "fmt"

    "helpers"
    "read_data"
    "save_svg"
    "transformations"
)

func main() {
    fmt.Println("Hello!")
    stops, lines := read_data.Read_data()
    grid := grid.Build_grid(&stops)
    //fmt.Println(grid)
    fmt.Println("- - - -")
    helpers.PrintGrid(&grid)
    fmt.Println("- - - -")
    save_svg.Save(stops, lines)
}
