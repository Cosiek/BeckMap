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
    return int(math.Sqrt(math.Pow(float64(x2 - x1), 2) + math.Pow(float64(y2 - y1), 2)))
}
