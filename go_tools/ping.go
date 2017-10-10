package main

import (
    "bytes"
    "fmt"
    "io"
    "net"
    "os"
)


func main(){
    if len(os.Args) != 2 {
        fmt.Println("usage: ", os.Args[0], "host")
        os.Exit(1)
    }

    addr, err := net.ResolveIPAddr("ip", os.Args[1])    
    if err != nil{
        fmt.Println("resolve ip err", err.Error())
    }


    conn, err := net.DialIP("ip4:icmp", addr, addr)
    fmt.Println(1)
    checkError(err)
    fmt.Println(2)
    var msg[512] byte
    msg[0] = 8
    msg[1] = 0 //echo
    msg[2] = 0 //code
    msg[3] = 0 //checksum fix later
    msg[4] = 0 //checksum fix later
    msg[4] = 0 // identifier[0]
    msg[5] = 13 //identifier[1]
    msg[6] = 0 //sequence[1]
    msg[7] = 37 //sequence[1]
    len := 8

    check := checkSum(msg[0:len])
    msg[2] = byte(check >> 8)
    msg[3] = byte(check & 255)

    _, err = conn.Write(msg[0:len])
    checkError(err)

    _, err = conn.Read(msg[0:])
    checkError(err)

    fmt.Println("got response")

    if msg[5] == 13 {
        fmt.Println("identifier matches")
    }
    if msg[7] == 37 {
        fmt.Println("Sequence matches")
    }
    os.Exit(0)
}

func checkSum(msg []byte) uint16 { 
    sum := 0
    for n := 1; n < len(msg)-1; n += 2 { 
        sum += int(msg[n])*256 + int(msg[n+1])
    }
    sum = (sum >> 16) + (sum & 0xffff) 
    sum += (sum >> 16)
    var answer uint16 = uint16(^sum) 
    return answer
}

func readFully(conn net.Conn) ([]byte, error) {
    defer conn.Close()
    result := bytes.NewBuffer(nil) 
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:]) 
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF { 
                break
            }
            return nil, err 
        }
    }
    return result.Bytes(), nil 
}


func checkError(err error) { 
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s \n", err.Error())
        os.Exit(1) 
    }
}