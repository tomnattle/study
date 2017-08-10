package main

import "fmt"
import "time"

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go func(index int, jobs <-chan int, results chan<- int) {
            for j := range jobs{
                //fmt.Println("worker", index, "start job")
                time.Sleep(time.Second)
                fmt.Println("worker", index, "finished job")
                results <- j * 2
            }
        }(w, jobs, results)
    }   

    for j := 1; j <= 10; j ++ {
        jobs <- j
    }

    // close至关重要 关闭后 意味着 range结束  活着说 <-jobs 是读取不到内容的

    close(jobs)

    // 这里阻塞 5次 则推出程序 前面的多路执行 5次阻塞后 后面的5次阻塞执行时 程序依然推出了
    
    for a := 1; a <= 5; a++ {
        fmt.Println(<- results)
    }
    time.Sleep(time.Second * 3)
}