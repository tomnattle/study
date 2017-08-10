package main

import "fmt"

type person struct{
    name string
    age int
}

func (p person) say(word string){
    fmt.Println(word, "my name is ", p.name, "i am ", p.age)
}

func (p person) test() {
    p.age = 15
}

func (p *person) test1() {
    p.age = 20
}

func main() {
    tommy := person{name:"tommy"}
    tommy.age = 10
    tommy.say("hi,")

    tommy.test()
    tommy.say("hi,")
    tommy.test1()
    tommy.say("hi,")
}