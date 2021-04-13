// exec("docker run <name>") + docker stop + docker stat
// edit /etc/nginx/nginx.conf
// restart nginx

/*
key: port
value: struct containing
	alias,
	container tag,
	date added,
	date modified,
	Desc,
	src,
	Account associated
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type AWSAppEntry struct {
	Alias        string
	ContainerTag string
	DateAdded    string
	DateModified string
	Desc         string
	Src          string
	Account      string
	Line         int
}

func addPort(alias string, tag string, added string, mod string, desc string, src string, acc string, line int) int {
	// create struct
	test := AWSAppEntry{
		Alias:        alias,
		ContainerTag: tag,
		DateAdded:    added,
		DateModified: mod,
		Desc:         desc,
		Src:          src,
		Account:      acc,
		Line:         line,
	}

	// read in json
	m := make(map[string]AWSAppEntry)

	// unmarshal

	// create new port
	port := checkPort()

	// add new port and info
	m[fmt.Sprint(port)] = test

	// marshal
	file, _ := json.MarshalIndent(m, "", "\t")

	// overwrite file
	_ = ioutil.WriteFile("ports.json", file, 0644)

	return 0
}

func checkPort() int {
	prg1 := "lsof"
	arg1 := "-i"
	arg2 := "-P"
	arg3 := "-n"
	prg2 := "grep"
	arg4 := "LISTEN"
	// declare variable for port flag
	// declare variable for port fowarding

	lsof := exec.Command(prg1, arg1, arg2, arg3) // add port flag + value

	grep := exec.Command(prg2, arg4)

	outPipe, err := lsof.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	lsof.Start()
	grep.Stdin = outPipe
	out, err := grep.Output()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print(string(out))

	// parse ports

	portsArray := strings.Split(string(out), "\n")
	var modifiedPortsArray []string

	// remove the last element
	portsArray = portsArray[:len(portsArray)-1]

	for _, s := range portsArray {
		temp := strings.Split(string(s), ":")[1]
		tempPort := strings.Split(string(temp), " ")[0]
		modifiedPortsArray = append(modifiedPortsArray, tempPort)
		//fmt.Println(tempPort)
	}

	retPort := -1
	flag := 0
	// 2000 to 9000
	for i := 2000; i <= 9000; i++ {
		flag = 0
		for _, inUse := range modifiedPortsArray {
			temp, _ := strconv.Atoi(inUse)
			if i == temp {
				flag = 1
				break
			}
		}
		if flag == 0 {
			retPort = i
			break
		}
	}
	// fmt.Println(retPort)
	return retPort
}

func removePort(port int) int {

	return 0
}
