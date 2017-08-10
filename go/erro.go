package main 

import "errors"
import "fmt"


func func_name(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can not work with 42")
    }

    fmt.Println("lalala demaxiya")
    return arg + 3, nil
}

type argError struct {
    arg int
    prob string
}

func (e * argError) Error() string{
    return fmt.Sprintf("%d %s" , e.arg, e.prob)
}


func func_name1(args int) (int, error) {
    if args == 42 {
        return -1, &argError{args, "can not work with 42"}
    }
    fmt.Println("lalala demaxiya")
    return args + 3, nil
}

func main(){
    for _, i := range []int{7, 42}{
        if r, e := func_name(i); e != nil{
            fmt.Println("func_name faile", e)
        }else{
            fmt.Println("func_name work", r)
        }
    }

    for _, i := range []int{7, 42}{
        if r, e := func_name1(i); e != nil{
            fmt.Println("func_name faile", e)
        }else{
            fmt.Println("func_name work", r)
        }
    }

}