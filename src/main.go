package main

import (
    "fmt"

    "read_data"
    "save_svg"
    "transformations"
)

func main() {
    fmt.Println("Hello!")
    stops, lines := read_data.Read_data()
    grid := grid.Build_grid(stops)
    fmt.Println(grid)
    save_svg.Save(stops, lines)
}
