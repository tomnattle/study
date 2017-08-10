package main 

import (
    "fmt"
    "sort"
)
func main() {
    strs := []string {"e", "b", "c"}
    sort.Strings(strs)
    fmt.Println("str", strs)

    ints  := []int {3, 2, 1}
    sort.Ints(ints)
    fmt.Println("int", ints)

    fmt.Println("int is sorts ", sort.IntsAreSorted(ints))
}