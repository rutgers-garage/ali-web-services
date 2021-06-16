package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type StartAWSServiceParams struct {
	Alias string
	req AWSServiceRequestBody
}

var port int = 8000
var ip string = "192.168.1.242"

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Printf("Client connected on port %v\n", port)
	return client
}

func main() {
	client := CreateClient()

	newAWSServiceRequestBody := AWSServiceRequestBody{
		ContainerTag: "alfonso/test",
		Desc: "This does something",
		Src: "github.com/alfonso",
		Account: "Alfonso",
	}

	newAWSService := StartAWSServiceParams{Alias: generateAlias(), req: newAWSServiceRequestBody}

	AWSRetString := ""
	AWSServiceErr := client.Call("Child.StartAWSService", newAWSService, &AWSRetString)
	if AWSServiceErr != nil {
		log.Fatal(AWSServiceErr)
	}

	fmt.Println(AWSRetString)
	
	upReqString := ""
	upRetInt := 0
	upErr := client.Call("Child.Up", upReqString, &upRetInt)
	if upErr != nil {
		log.Fatal(upErr)
	}

	fmt.Println(upRetInt)

	// http.HandleFunc("/", handler)
	// http.HandleFunc("/start-service", startService)
	// fmt.Println("Server running on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
