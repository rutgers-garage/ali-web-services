// exec("docker run <name>") + docker stop + docker stat
// edit /etc/nginx/nginx.conf
// restart nginx

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func startDockerService(container string, port string, innerPort string) string {
	docker := "docker"
	run := "run"
	d := "-d"
	p := "-p"
	args := fmt.Sprintf("%s:%s", port, innerPort) // FIRST NUMBER SHOULD BE portAssigner.checkPort() SECOND NUMBER SHOULD BE USER PORT
	fmt.Println(fmt.Sprintf("%s %s %s %s %s %s", docker, run, d, p, args, container))
	cmd := exec.Command(docker, run, d, p, args, container) // add port flag + value
	stdout, err := cmd.Output()
	containerId := string(stdout)
	if err != nil {
		fmt.Println("Docker?????")
		log.Fatal(err)
	}

	fmt.Print(containerId)
	return containerId[:len(containerId)-2]

}
