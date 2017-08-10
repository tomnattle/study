package main 

import "fmt"
import "time"

func main() {
    message := make(chan string)
    signal := make(chan bool)

    select{
        case msg := <-message:
            fmt.Println("received message", msg)
        default:
            fmt.Println("no message received")
    }

    msg := "hi"

    select{
        case message <- msg:
            fmt.Println("send message", msg)
        default:
            fmt.Println("no message sent")
    }

    select {
        case msg := <-message:
            fmt.Println("received message", msg)
        case sig := <-signal:
            fmt.Println("received signal", sig)
        default:
            fmt.Println("no activity")
    }

    _message := make(chan string)

    go func(){
        time.Sleep(time.Second * 3)
        _message <- "hello jo"
    }()

    _m := <-_message
    fmt.Println(_m)
    fmt.Println("finish")

    time.Sleep(time.Second * 3)

}
