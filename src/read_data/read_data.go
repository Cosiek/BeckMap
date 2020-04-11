package read_data

import (
    "log"
    "strconv"

	"github.com/ksev/ods"

    "structs"
)

func Read_data() (map[int]structs.Stop, []structs.Line){
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
    stops := make(map[int]structs.Stop)
    lines := make([]structs.Line, 0)
    for _, row := range table {
        if row[0] == "Stops" || row[0] == "Lines"{
            t = row[0]
        } else if row[0] == "Id"{
            continue
        } else {
            if t == "Stops"{
                id, _ := strconv.Atoi(row[0])
                x, _ := strconv.Atoi(row[2])
                y, _ := strconv.Atoi(row[3])
                stop := structs.Stop{id, row[1], x, y}
                stops[stop.Id] = stop
            } else {
                id, _ := strconv.Atoi(row[0])
                l := make([]*structs.Stop, 0)
                for i := 2; i < len(row); i++{
                    v, _ := strconv.Atoi(row[i])
                    stop := stops[v]
                    l = append(l, &stop)
                }

                line := structs.Line{id, row[1], l}
                lines = append(lines, line)
            }
        }
    }
    return stops, lines
}
