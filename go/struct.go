package main

import "fmt"

type person struct{
    name string
    age int
}


func main() {
    fmt.Println(person{"tommy",30})
    fmt.Println(person{name:"json",age:32})
    fmt.Println(person{name:"space"})

    tommy := person{name : "tommy"}
    tommy.age = 15
    fmt.Println(tommy)

    fmt.Println("tommy's address" , &tommy)
    _tommy := &tommy
    __tommy := tommy
    _tommy.age = 16

    
    fmt.Println(tommy, _tommy,__tommy)

}