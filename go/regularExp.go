package main

import "fmt"
import "regexp"
import "bytes"

func main() {
    match, x := regexp.MatchString("p([a-z]+)ch" ,"peach")
    fmt.Println(match, x)

    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))
    fmt.Println(r.FindString("peach apple"))
    fmt.Println(r.FindString("peaeech apple"))
    fmt.Println(r.FindStringIndex("peach apple"))
    fmt.Println(r.FindStringIndex("peaeech apple"))
    fmt.Println(r.FindStringIndex("peaeech apple"))
    fmt.Println(r.FindStringIndex("pea-eech apple punch peach"))
    fmt.Println("FindStringIndex [start end]")
    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println(r.FindStringSubmatch("punch peach"))

    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    fmt.Println(r.FindAllString("peach punch pinch", -1))
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}