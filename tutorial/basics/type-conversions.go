package main

import (
    "fmt"
    "math"
)

/**
* In general the T(v) converts the value v to type T
*/
func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f) // Assignment of non-matching types in Go REQUIRES explicit conversions

    fmt.Println(x, y, z)
}

