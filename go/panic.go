package main

import "os"
//import "fmt"

func main() {
    //panic("a problem")

    _, err := os.Create("file")
    if err != nil {
        panic(err)
    }
   
    //fmt.Println(a)
}