package main

import "fmt"
import "math"
import "time"

const name string = "tommy"
func main() {
    fmt.Println("go" + "lang")
    fmt.Println("1+2:", 1 + 2)
    fmt.Println("7.0/3.0" , 7.0 / 3.0)

    fmt.Println(true || false)
    fmt.Println(!false)

    var a string = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b + c)

    d := 1
    fmt.Println(d)
    fmt.Println(name)

    //name := "wanghui";

    // erro
    //const abc int = a + 

    const myname = "my name is " + name

    fmt.Println(myname)

    fmt.Println(math.Sin(1))
    fmt.Println(int64(math.Sin(1)))

    i := 0
    for i <= 2 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 2; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("abcdefg")
        break
    }
    var result bool
    isOdd(6, result)
    fmt.Println("7 is ", &result)

    switch b{
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        default:
            fmt.Println("not one not two")
    }

    fmt.Println(time.Now())
    fmt.Println(time.Now().Day())
    fmt.Println(time.Now().Month())
    fmt.Println(time.Now().Weekday())

    isInt := func (i interface{}){
        switch t := i.(type) {
            case int:
                fmt.Println(i, " is int")
            default:
                fmt.Println("unkown type" , t)
        }
    }
    isInt(7)
    isInt("abc")
    isInt(false)

    var arr[5] int
    fmt.Println("emp: ", arr)
    arr[1] = 100
    fmt.Println("emp: ", arr)
    fmt.Println("arr[1]", arr[1])

    arr1 := [5]int{1, 2, 3}
    fmt.Println("arr1: ", arr1)
    fmt.Println("arr1 len: ", len(arr1))

    s := make([]string ,3)

    fmt.Println(s)
    s[0] = "nice job"
    s[1] = "omg"
    s[2] = ""
    fmt.Println(s,len(s))
    s = append(s, "a")
    fmt.Println(s,len(s))

    cs := make([]string, len(s))
    copy(cs, s)
    fmt.Println(cs)
    
    tt := []string {"a", "b", "c"}
    fmt.Println(tt)

    ll := tt[:1]
    mm := tt[2:]
    fmt.Println(ll,mm)

    as := make([][] int, 3)
    as[0] = make([]int, 3)
    as[1] = make([]int, 3)
    as[0][0] = 1
    fmt.Println(as)

    at := [5] int {1}
    fmt.Println(at)

    mp := make(map[string]int)
    mp["abc"] = 124
    mp["efg"] = 567
    delete(mp, "abc")

    _, prs := mp["efg"]
    fmt.Println(mp)
    fmt.Println(prs)

    abcde, prs1 := mp["abcde"]
    fmt.Println(abcde, prs1)

    
    sd := map[string] int{"level" : 5, "age" :10}
    fmt.Println(sd)
    

    nums := []int{1, 2, 3}
    sums := 0

    for _, num := range nums{
        sums += num
    }
    fmt.Println(sums)

    fmt.Println("\n")


    kvs := map[string] string{"a" : "A", "b" : "B", "c" : "C"}
    for k, v := range kvs{
        fmt.Println(k, ":", v)
    }

    for k, v := range "go go go!"{
        fmt.Println(k, ":", v)
    }

    fmt.Println(plus(1, 2))

    s1, t1 := values()
    fmt.Println(s1, t1)

    m1, _ := values()
    fmt.Println(m1)
    _sums(1, 2, 3, 4, 5)

    _sum := _sums

    _sum(100, 200, 300)

    _func := intSeq()
    fmt.Println(_func())
    fmt.Println(_func())
    fmt.Println(_func())
    fmt.Println(_func())
}

func intSeq() func() int {
    i := 0
    return func() int{
        i ++
        return i
    }
}


func _sums(ns ...int){
    fmt.Println(ns, " ")
    total := 0
    for _, n := range ns{
        total += n
    }
    fmt.Println("total", total)
}

func values () (int, int){
    return 1, 2
}

func plus(a int, b int) int {
    return a + b
}

func isOdd(i int, result bool){
    if i % 2 == 0 {
        result = true 
    }else{
        result = false
    }
    
}

