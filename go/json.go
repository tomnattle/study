package main

import "encoding/json"
import "fmt"
import "os"

type Response struct{
    Page int
    Fruits []string
}

type Response2 struct{
    Page int    `json : "page"`
    Fruits []string `json : "fruits"`
}

func main() {
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(1.234)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("abc")
    fmt.Println(string(strB))

    slcB, _ := json.Marshal([]string{"a", "b", "c"})
    fmt.Println(string(slcB))

    mapB, _ := json.Marshal(map[string]int {"a":1,"b":2})
    fmt.Println(string(mapB))

    res := &Response{
        Page : 1,
        Fruits : []string {"apple", "peach", "pear"},
    }
    strct, _ := json.Marshal(res)
    fmt.Println(string(strct))

    res2 := &Response2{
        Page : 1,
        Fruits : []string {"apple", "peach", "pear"},
    }
    strct2, _ := json.Marshal(res2)
    fmt.Println(string(strct2))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)


    var dat map[string] interface{}

    if err := json.Unmarshal(byt, &dat); err != nil{
        panic(err)
    }

    fmt.Println(dat)

    num := dat["num"].(float64)
    fmt.Println(num)

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    _res := &Response2{}

    json.Unmarshal([]byte(str), &_res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int {"apple" : 5, "lettuce" : 7}
    enc.Encode(d)
}












