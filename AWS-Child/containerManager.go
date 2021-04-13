// exec("docker run <name>") + docker stop + docker stat
// edit /etc/nginx/nginx.conf
// restart nginx

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func startService(container string) {
	docker := "docker"
	run := "run"
	p := "-p"
	args := fmt.Sprintf("%d:%d", 3000, 3000) // FIRST NUMBER SHOULD BE portAssigner.checkPort() SECOND NUMBER SHOULD BE USER PORT

	cmd := exec.Command(docker, run, p, args, container) // add port flag + value
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(stdout))

}
