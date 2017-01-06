package main

import (
    "fmt"

    "read_data"
)

func main() {
    fmt.Println("Hello!")
    stops, lines := read_data.Read_data()
    fmt.Println(stops)
    fmt.Println(lines)
    fmt.Println("Done!")
}
