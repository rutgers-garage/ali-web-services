package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func checkPort() string {
	prg1 := "lsof"
	arg1 := "-i"
	arg2 := "-P"
	arg3 := "-n"
	prg2 := "grep"
	arg4 := "LISTEN"

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

	retPort := ""
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
			retPort = strconv.Itoa(i)
			break
		}
	}
	fmt.Println(retPort)
	return retPort
}

func removePort(port int) int {

	return 0
}
