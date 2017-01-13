package simple_math

import "math"


func Max(a int, b int) int{
    if a > b { return a }
    return b
}

func Min(a int, b int) int{
    if a < b { return a }
    return b
}


func Distance(x1, x2, y1, y2 int) int{
    return math.Sqrt(math.Pow(x2 - x1) + math.Pow(y2 - y2))
}
