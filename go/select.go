package main

import "fmt"
import "time"

func main(){
    c1 := make(chan string)
    c2 := make(chan string)

    go func(_time int){
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }(1)

    go func(_time int){
        time.Sleep(time.Second * 3)
        c2 <- "two"
    }(2)

    for i := 0; i < 2; i++ {
        select {
            case msg := <-c2:
                fmt.Println("received", msg, " from c2")
            case msg := <-c1:
                fmt.Println("received", msg, " from c1")
        }
    }

    fmt.Println("finish")
}