package main

import "fmt"
import "os"
import "time"
import "os/signal"
import "syscall"

func main() {
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    defer fmt.Println("!...........")
    
    go func(){
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    for{
        time.Sleep(2)
        fmt.Println("...")
        break
    }
    fmt.Println("exiting")
    
    os.Exit(3)
}