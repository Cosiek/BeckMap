package helpers

import "fmt"
import "structs"

func PrintGrid(grid *structs.Grid){
    for _, row := range grid.Grid{
        fmt.Println(row)
    }
}
