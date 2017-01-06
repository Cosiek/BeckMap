package structs


type Stop struct {
    Id int
    Name string
    X, Y int    // coordinates
}


type Line struct {
    Id int
    Name string
    Stops []*Stop
}
