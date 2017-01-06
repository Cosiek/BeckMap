package main

import (
    "fmt"

    "read_data"
    "save_svg"
)

func main() {
    fmt.Println("Hello!")
    stops, lines := read_data.Read_data()
    save_svg.Save(stops, lines)
}
