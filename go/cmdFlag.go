package main 

import "flag"
import "fmt"

func main() {
    //go run  cmdFlag.go numb 1 fork true
    
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var sv string
    flag.StringVar(&sv, "sv", "bar", "a string bar")
    flag.Parse()
    
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", sv)
    fmt.Println("tail:", flag.Args())
    args := flag.Args()

    fmt.Println(args[4])
}