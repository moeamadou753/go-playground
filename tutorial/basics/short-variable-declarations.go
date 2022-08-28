package main

import "fmt"

func main() {
    var i, j int = 1, 2
    k := 3 // Used in place of the var declaration, uses implicit typing
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
} // := construct is not available outside of a function since top-level expressions are not allowed without a keyword (var, func, etc.)

