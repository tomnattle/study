package main
import "fmt"
import "time"
    
func main(){

    done := make (chan bool, 1)

    go func (chan bool){
        fmt.Println("working ...")
        time.Sleep(time.Second * 2)
        fmt.Println("down")
        done <- false
        fmt.Println("----")
        done <- true
    }(done)

    fmt.Println("wait for value")
    for {
        var msg bool = <-done
        fmt.Println(msg) 
        time.Sleep(time.Second * 2)
        if msg {
            break
        }
    }
    fmt.Println("finish") 
}