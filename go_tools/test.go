package main 


import "time"
import "os/exec"
import "fmt"
import "strings"
import "bytes"
import "sync/atomic"

func main(){
    var connCount uint64 = 0
    for i := 0; i < 1; i = i+ 1{
        go func(){
            for {
                //time.Sleep(time.Millisecond * 1000)
                atomic.AddUint64(&connCount, 1)
                cmd := exec.Command("telnet", "localhost", "1202")
                cmd.Stdin = strings.NewReader("some input")
                var out bytes.Buffer
                var stderr bytes.Buffer
                cmd.Stdout = &out
                cmd.Stderr = &stderr
                cmd.Run()

                //if err != nil {
                //    fmt.Println("err:", err.Error(), stderr.String(), out.String())
                //}
                time.Sleep(1 * time.Second)
                fmt.Println("run", atomic.LoadUint64(&connCount), out.String())
                
            }
        }()
    }
    
    time.Sleep(time.Second * 60)
}