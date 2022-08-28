package main

import "fmt"

/**
* Uninitialized variables are given a 'zero value'. This value is 0 for numeric types, false for booleans, and the empty string for strings
*/
func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

