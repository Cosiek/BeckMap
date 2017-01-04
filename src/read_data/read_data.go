package read_data

import (
    "fmt"
    "log"
    "strings"

    "third_party/ods"
)

func Read_data() {
    reader, err := ods.OpenReader("file.ods")
    if err != nil {
        log.Fatal(err)
    }
    defer reader.Close()

    table, err := ods.Decode(reader)
    if err != nil {
        log.Fatal(err)
    }

    for _, row := range table {
        fmt.Println(strings.Join(row, "|"))
    }
}
