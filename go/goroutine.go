package main 
import "fmt"

func f(arg string) {
    for i:=0; i < 3; i++ {
        fmt.Println(arg, ":", i)
    }
}

func main() {
    f("nomal")

    go f("routine")

    go func(msg string) {
        fmt.Println(msg)
    }("go")

    var input string
    fmt.Scanln(&input)
    fmt.Println("down")

    messages := make(chan string)
    go func() {
        fmt.Println(1)
        messages <- "ping"
        fmt.Println(2)
    }()

    fmt.Println(3)
    msg := <- messages
    fmt.Println(4)
    fmt.Println(msg)

}