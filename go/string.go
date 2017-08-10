package main 

import (
    "fmt"
    s "strings"    
)
var p = fmt.Println

func main() {
    p("contains" , s.Contains("test that","that"))
    p("count", s.Count("hi joso", "o"))
    p("hasPrefix", s.HasPrefix("hi koa", "h"))
    p("index", s.Index("abc" , "b"))

    p("join", s.Join([]string{"hi", "tom"}, "-"))
    p("repeat", s.Repeat("a", 5))
    p("replace", s.Replace("abcde", "b", "-", 1))
    p("replace", s.Replace("abcdeb", "b", "-", -1))
    p("split", s.Split("abcdef", ""))
    p("tolower", s.ToLower("JWKJC"))
    p("toupper", s.ToUpper("jwkjc"))
    p()

    p("len", len("hello"))
    p("char", "hello" [4])
}