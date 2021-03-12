// exec("docker run <name>") + docker stop + docker stat
// edit /etc/nginx/nginx.conf
// restart nginx

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func startService(container string) {
	prg := "docker"
	arg1 := "--version"
	// declare variable for port flag
	// declare variable for port fowarding

	cmd := exec.Command(prg, arg1) // add port flag + value
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(stdout))

}

func modifyNginx(port int) {
	f, err := os.OpenFile("nginx.conf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	if _, err := f.WriteString("penis penis penis\n"); err != nil {
		log.Println(err)
	}
}
