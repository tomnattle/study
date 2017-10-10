package main 

import (
    "net"
    "os"
    "fmt"
)

func main() {
    service := ":1201"
    tcpServer, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpServer)
    checkError(err)

    for {
        conn, err := listener.Accept()
        fmt.Println("loop")
        if err != nil{
            continue
        }
        handelClient(conn)
        conn.Close()
    }
}

func handelClient(conn net.Conn) {
    var buf[512] byte
    for {
        n, err := conn.Read(buf[:0])
        if err != nil{
            return
        }

        fmt.Println(string(buf[0:]))
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil{
            return
        }
    }
}


func checkError(err error) { 
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1) 
    }
}