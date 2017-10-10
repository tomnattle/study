package main 

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/shutdown", shutdown)
    http.HandleFunc("/", homePage)
    fmt.Println("listen 8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil{
        panic(err)
    }
}

func shutdown(res http.ResponseWriter, req *http.Request){
    
    //fmt.Fprint(res, "server shutdown")
    defer os.Exit(0)
}

func homePage(res http.ResponseWriter, req *http.Request){
    if req.URL.Path != "/" {
        http.NotFound(res, req)
        return
    }
    fmt.Fprint(res, "this is home page.")
}