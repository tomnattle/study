package main
import "fmt"

func main(){
    msg := make(chan string, 2)

    msg <- "first"
    msg <- "second"

    fmt.Println(<-msg)
    fmt.Println(<-msg)
}