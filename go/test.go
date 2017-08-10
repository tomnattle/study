package main 

import "fmt"
import "time"
func main(){

    message := make(chan int) //1
    go func(){
        time.Sleep(2) //3
        fmt.Println(<-message) //4
        message<- 2 //5
    }()
    message<- 1 //2

    fmt.Println(<-message) //6
}