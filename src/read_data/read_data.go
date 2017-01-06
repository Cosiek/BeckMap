package read_data

import (
    "fmt"
    "log"
    "strconv"
    "strings"

    "third_party/ods"

    "structs"
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

    var t string
    stops := make([]structs.Stop, 0)
    for _, row := range table {
        if row[0] == "Stops" || row[0] == "Lines"{
            fmt.Println(row[0])
            t = row[0]
        } else if row[0] == "Id"{
            continue
        } else {
            if t == "Stops"{
                id, _ := strconv.Atoi(row[0])
                x, _ := strconv.Atoi(row[2])
                y, _ := strconv.Atoi(row[3])
                stop := structs.Stop{id, row[1], x, y}
                stops = append(stops, stop)
                fmt.Println(stop)
            } else {
                fmt.Println(strings.Join(row, "|"))
            }
        }
    }
    fmt.Println(stops)
}
