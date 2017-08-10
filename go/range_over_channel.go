package main 

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "hi"
    queue <- "hello"
    fmt.Println("end fill")

    close(queue)
    for ele := range queue{
        fmt.Println(ele)
    }

    fmt.Println("nice finish")
}