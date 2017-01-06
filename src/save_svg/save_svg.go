package save_svg

import (
    "fmt"
    "io/ioutil"

    "structs"
)

func Save(stops map[int]structs.Stop, lines []structs.Line){
    d1 := []byte("Hello World")
    err := ioutil.WriteFile("export.svg", d1, 0644)
    if err != nil{
        panic(err)
    } else {
        fmt.Println("saved.")
    }
}
