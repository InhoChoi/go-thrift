package main

import (
    "fmt"
    "github.com/inhochoi/go-thrift/thrift/helloworld"
)

type HelloWorldHandler struct {
}

func NewHelloWorldHandler() *HelloWorldHandler {
    return &HelloWorldHandler{}
}

func (p *HelloWorldHandler) Helloworld(hello *helloworld.Hello) (int32, error) {
    fmt.Println("Hello World")
    return hello.Number, nil
}
