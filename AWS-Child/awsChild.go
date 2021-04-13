// exec("docker run <name>") + docker stop + docker stat
// edit /etc/nginx/nginx.conf
// restart nginx

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	startService("https://hemangandhi.github.io")
}

func startService(container string) {
	prg := "open"
	//arg1 := "pns"
	// declare variable for port flag
	// declare variable for port fowarding

	cmd := exec.Command(prg, container) // add port flag + value
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(stdout))

}
