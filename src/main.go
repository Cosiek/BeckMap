package main

import (
	"fmt"

	//"helpers"
	"read_data"
	"save_svg"
	"transformations"
)

func main() {
	fmt.Println("Hello!")
	stops, lines := read_data.Read_data()
	grid.ApplyGrid(&stops, &lines)
	save_svg.SaveMap(stops, lines)
	save_svg.SaveGrid(&stops, &lines)
}
