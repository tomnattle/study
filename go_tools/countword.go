package main 
//go run countword.go  test.txt


import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sync"
)

// 定义一个数据结构
type words struct{
    sync.Mutex
    found map[string] int
}

// 数据结构定义方法
func (w *words) add(word string, n int){
    w.Lock()
    defer w.Unlock()
    count, ok := w.found[word]
    if !ok{
        w.found[word] = n
        return
    }
    w.found[word] = count + n
}

// 数据结构实例化
func newWords() *words{
    return &words{ found: map[string] int {}}
}

func main() {
    var wg sync.WaitGroup
    w := newWords()
    for _, f := range os.Args[1:] {
        wg.Add(1)
        go func(filename string){
             print(filename)
            if err := tallyWords(filename, w); err != nil{
                fmt.Println(err.Error())
            }
            wg.Done()
        }(f)
    }
    wg.Wait()

    for k, v := range w.found{
        fmt.Println(k, v)
    }
}

func tallyWords(fileName string, dict *words) error{
    file, err := os.Open(fileName)
    if err != nil {
        return err
    }
    print(11)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        //fmt.Println(scanner.Text())
        word := strings.ToLower(scanner.Text())
        dict.add(word, 1)
    }
    return scanner.Err()
}