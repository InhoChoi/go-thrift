package main

import (
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/inhochoi/go-thrift/thrift/helloworld"
)

func handleClient(client *helloworld.HelloWorldClient) (err error) {
	data := &helloworld.Hello{}
	data.Name = "InhoChoi"
	data.Number = 32
	result, _ := client.Helloworld(data)

	fmt.Println("Result :", result)
	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	return handleClient(helloworld.NewHelloWorldClientFactory(transport, protocolFactory))
}
