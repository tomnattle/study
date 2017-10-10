package main
import ( 
    "net" 
    "time"
    "fmt"
    "os"
)

func main() {
    service := ":1203"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    //checkError(err)
    conn, err := net.ListenUDP("udp", udpAddr) 
    checkError(err)
    for { 
        handleClient(conn)
    }
}

func handleClient(conn *net.UDPConn) { 
    var buf [512]byte
    msg, addr, err := conn.ReadFromUDP(buf[0:]) 
    if err != nil {
        return
    }
    fmt.Println("msg:", msg)
    daytime := time.Now().String()
    conn.WriteToUDP([]byte(daytime), addr) 
}

func checkError(err error) { if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1) }
}