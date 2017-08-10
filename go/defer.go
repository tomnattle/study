package main

import "fmt"
import "os"

func main() {
    f := createFile("/tmp/defer")
    defer closeFile(f)
    writeFile(f)
}

func createFile(fname string) *os.File {
    fmt.Println("create")
    f, err := os.Create(fname)
    if err != nil{
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("write")
    fmt.Fprintln(f, "nice day")
}

func closeFile(f *os.File) {
    fmt.Println("close")
    f.Close()
}