package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"time"
)

type AWSServiceRequestBody struct {
	ContainerTag string
	Desc         string
	Src          string
	Account      string
	InnerPort    string
}

type StartAWSServiceParams struct {
	Alias string
	Req   AWSServiceRequestBody
}

type Child struct{}

func (c *Child) StartAWSService(args *StartAWSServiceParams, ret *string) error {
	n, err := json.Marshal(*args)
	if err != nil {
		fmt.Println("reeeeee")
	}

	fmt.Println(string(n))
	port := checkPort()
	containerId := startDockerService(args.Req.ContainerTag, port, args.Req.InnerPort)

	initNginx()
	addProxy(args.Alias, port)

	addPort(args.Alias, args.Req.ContainerTag, time.Now().String(), time.Now().String(), args.Req.Desc, args.Req.Src, args.Req.Account, args.Req.InnerPort, port, containerId)

	*ret = "200"

	return nil
}

func (c *Child) Up(req *string, ret *int) error {
	m := make(map[string]bool)
	readBytes, _ := ioutil.ReadFile("ports.json")
	json.Unmarshal(readBytes, &m)

	*ret = len(m)

	return nil
}

func main() {
	c := &Child{}
	rpc.Register(c)

	fmt.Println("CHILD UP ON 8000")
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(fmt.Printf("Unable to listen on given port: 8000", err))
	}

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
