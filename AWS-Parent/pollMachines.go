package main

import (
	"encoding/json"
	"io/ioutil"
)

func pollMachines() {
	m := make(map[string]bool)
	readBytes, _ := ioutil.ReadFile("takenAliases.json")
	json.Unmarshal(readBytes, &m)

	// for _, kv := range m {
    //     k := kv[1]
    //     v := kv[2]
    // }
}

func pollMachine(ip string) {
	
}