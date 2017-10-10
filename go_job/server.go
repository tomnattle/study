package server
import (
    "fmt"
    "log"
)

type struct server{
    config map[string]string
}

func (s *server) start(){
    log.Println("server start")
}

