package main

import "time"
import "fmt"

func main() {
    timer1 := time.NewTimer(time.Second * 2)
    fmt.Println("create timmer1")
    <-timer1.C
    fmt.Println("timer1 expired")

    timer2 := time.NewTimer(time.Second)
    go func(){
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    //  定时器 应该在2秒中 写入 信道 可以在主进程中 这个定时器被关闭了 所以不会触发信道读取操作 
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}