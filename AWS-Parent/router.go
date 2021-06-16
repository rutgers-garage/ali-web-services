package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AWSServiceRequestBody struct {
    ContainerTag string
    Desc string
    Src string
    Account string
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			fmt.Println("GET request")
			http.ServeFile(w, req, "home.html")
		case "POST":
			fmt.Println("POST request")

			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			decoder := json.NewDecoder(req.Body)
			c := AWSServiceRequestBody{}
			err := decoder.Decode(&c)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(c)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func startService(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			decoder := json.NewDecoder(req.Body)
			c := AWSServiceRequestBody{}
			err := decoder.Decode(&c)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(c)

			newAlias := generateAlias()

			fmt.Println(newAlias)
		default:
			fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}
