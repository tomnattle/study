package main 

import (
    "fmt"
    //"net"
    "os"
    "encoding/json"
    //"bytes"
    //"io"
)


type Person struct{
    Name string
    Age int
    Email []Email
}

type Email struct{
    Kind string
    Address string
}

func main() {
    pers := Person{
        Name : "tommy",
        Age : 30,
        Email : [] Email {Email{Kind:"home", Address:"xiaohe cun"}, Email{Kind:"company", Address:"shangbai"}},
    }

    result, err := json.Marshal(&pers)
    if err != err{
        panic(err.Error())
    }
    fmt.Println(result)
    fmt.Println(string(result))

    p := Person{}
    fmt.Println(json.Unmarshal([]byte(result), &p))
    fmt.Println(p)
    fmt.Println("--")
    enc := json.NewEncoder(os.Stdout)
    enc.Encode(p) 

}