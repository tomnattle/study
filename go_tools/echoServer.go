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

    service := ":1202"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    
    
    checkError(err)

    for {
        conn, err := listener.Accept()
        atomic.AddUint64(&connCount, 1)
        // 如果链接出现问题丢提
        if err != nil {
            atomic.AddUint64(&errCount, 1)
            //log.Println("erro happend, ignore", err.Error())
            continue
        }
        //log.Println("new conn")
        log.Println("new conn", atomic.LoadUint64(&connCount), conn.RemoteAddr())
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn){
    //time.Sleep(time.Second * 1)
    conn.Write([]byte("hello tommy!"))
    defer conn.Close()
}

func checkError(err error) { if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1) }
}