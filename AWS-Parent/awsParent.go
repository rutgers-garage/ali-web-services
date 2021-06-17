package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type AWSServiceRequestBody struct {
    ContainerTag 	string
    Desc 			string
    Src 			string
    Account 		string
	InnerPort 		string
}

type StartAWSServiceParams struct {
	Alias 	string
	Req 	AWSServiceRequestBody
}

var port int = 8000
var ip string = "192.168.1.242"
var client *rpc.Client

func createClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	fmt.Printf("Client connected on port %v\n", port)
	return client
}

func startAWSService(body AWSServiceRequestBody) {
	newAWSService := StartAWSServiceParams{Alias: generateAlias(), Req: body}

	AWSRetString := ""
	AWSServiceErr := client.Call("Child.StartAWSService", &newAWSService, &AWSRetString)
	if AWSServiceErr != nil {
		log.Fatal(AWSServiceErr)
	}

	fmt.Printf("startAWSService response is %v\n", AWSRetString)
}

func up() {
	upReqString := ""
	upRetInt := 0
	upErr := client.Call("Child.Up", &upReqString, &upRetInt)
	if upErr != nil {
		log.Fatal(upErr)
	}

	fmt.Printf("up response is %v\n", upRetInt)
}

func main() {
	client = createClient()

	http.HandleFunc("/start-service", startService)
	http.ListenAndServe(":8080", nil)
}
