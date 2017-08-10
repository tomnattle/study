package main
import "fmt"

func main(){

    a := make(chan string, 1)
    b := make(chan string, 1)
    go func (pings chan<- string, msg string){
        a <- msg
    }(a, "hello")

    go func(c1 <-chan string, c2 chan<- string){
        msg := <-c1
        c2 <- msg
    }(a, b)

    fmt.Println(<-b)
}