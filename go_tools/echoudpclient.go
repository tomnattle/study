package main 

import (
    "fmt"
    "os"
    "net"
    "log"
    "os/signal"
    "syscall"
    "sync/atomic"
    //"time"
)


func main() {

    _signal := make(chan os.Signal,1)
    signal.Notify(_signal, syscall.SIGINT, syscall.SIGTERM)
    var connCount uint64 = 0
    var errCount uint64 = 0

    go func (){
        fmt.Println("wait for signal")
        sig := <- _signal
        fmt.Print("sigal catch")
        fmt.Println()
        fmt.Println(sig)
        log.Println("total rquest", atomic.LoadUint64(&connCount))
        log.Println("error rquest", atomic.LoadUint64(&errCount))
        os.Exit(3)
    }()

    service := ":1203"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)

    conn, err := net.DialUDP("udp", nil, udpAddr)
    checkError(err)

    _, err = conn.Write([]byte("hello tommy"))
    checkError(err)

    var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)

    fmt.Println(string(buf[0:n]))
    os.Exit(0)
    
}



func checkError(err error) { if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1) }
}