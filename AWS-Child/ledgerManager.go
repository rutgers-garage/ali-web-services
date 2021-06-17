package main

import (
	"encoding/json"
	"io/ioutil"
)

type AWSAppEntry struct {
	Alias        string
	ContainerTag string
	DateAdded    string
	DateModified string
	Desc         string
	Src          string
	Account      string
	InnerPort    string
	Port: 		 string
	ContainerId  string
}

func addPort(alias string, tag string, added string, mod string, desc string, src string, acc string, innerPort string, port string, containerId string) {
	// create struct
	entry := AWSAppEntry{
		Alias:        alias,
		ContainerTag: tag,
		DateAdded:    added,
		DateModified: mod,
		Desc:         desc,
		Src:          src,
		Account:      acc,
		InnerPort:    innerPort,
		Port:         port,
		ContainerId:  containerId,
	}

	// read in json
	m := make(map[string]AWSAppEntry)

	// unmarshal
	readBytes, _ := ioutil.ReadFile("ports.json")
	json.Unmarshal(readBytes, &m)

	// add new port and info
	m[entry.Alias] = entry

	// marshal
	file, _ := json.MarshalIndent(m, "", "\t")

	// overwrite file
	_ = ioutil.WriteFile("ports.json", file, 0644)

}
