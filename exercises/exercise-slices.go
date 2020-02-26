package main

import "golang.org/x/tour/pic"

/**
* Returns a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers
*/
func Pic(dx, dy int) [][]uint8 {
    slice :=  make([][]uint8, dy)
    
    for i, v := range slice {
        v := make([]uint8, dx)
        for j, w := range v {
            w := uint(i^j)
        }
    }
    
    return slice
}

func main() {
    pic.Show(Pic)
}
