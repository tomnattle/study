package main 

import "fmt"


func main(){
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <- jobs
            if more{
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                // 这里如果不返回 则有问题
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }

    for j := 1; j <= 33; j+=10 {
        jobs <- j
        fmt.Println("sent job", j)
    }

    close(jobs)

    fmt.Println("send all ============")
    fmt.Println(<-done)
}