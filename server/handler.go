package main

import (
    "fmt"
)

type HelloWorldHandler struct {
}

func NewHelloWorldHandler() *HelloWorldHandler {
    return &HelloWorldHandler{}
}

func (p *HelloWorldHandler) Helloworld() (err error) {
    fmt.Println("Hello World")
    return nil
}
