package main 

import "os"
import "fmt"

func main() {
    argWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    arg := os.Args[0]
    fmt.Println(argWithProg, argsWithoutProg, arg)
}