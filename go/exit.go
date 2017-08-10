package main 

import "os"
import "fmt"

func main() {

    defer fmt.Println("lalalal", "this will nerver be call")

    os.Exit(3)
}