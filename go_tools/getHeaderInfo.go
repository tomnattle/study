package main 

import (
    "net"
    "os"
    "fmt"
    "io/ioutil"
)

func main(){
    service := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)

    _, err = conn.Write([] byte("HEAD / HTTP/1.0\r\n\r\n"))

    result, err := ioutil.ReadAll(conn)

    fmt.Println(string(result))

    checkError(err)
    os.Exit(0)
}

func checkError(err error) { if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1) }
}