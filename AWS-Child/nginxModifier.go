package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var m map[string]string
var lines []string
var nginxConf *os.File

func initNginx() {
	nginxConf = openNginxConf()
	dump, _ := ioutil.ReadAll(nginxConf)
	lines = strings.Split(string(dump), "\n")
	m = deserializeNginx(nginxConf, lines)
}

func openNginxConf() *os.File {
	f, err := os.OpenFile("nginx.conf", os.O_RDWR, 0644)

	if err != nil {
		log.Println(err)
	}
	return f
}

func closeNginxConf() {
	err := nginxConf.Close()
	if err != nil {
		panic(err)
	}
}

func addProxy(alias string, port string) {
	m[alias] = port
	locations := serializeNginx()
	newFile := strings.Join(append(lines[:54], locations...), "\n")
	writeFile(nginxConf, newFile)
}

func removeProxy(alias string) {
	delete(m, alias)
	locations := serializeNginx()
	newFile := strings.Join(append(lines[:54], locations...), "\n")
	writeFile(nginxConf, newFile)
}

func deserializeNginx(f *os.File, lines []string) map[string]string {

	m = make(map[string]string)

	aliasRegex, _ := regexp.Compile(`/[a-z]*_[a-z]*/`)
	portRegex, _ := regexp.Compile(`\d\d\d\d`)

	for i := 54; i < len(lines)-3; i += 4 {
		alias := aliasRegex.FindString(lines[i])
		if len(alias) > 0 {
			alias = alias[1 : len(alias)-1]
		} else {
			fmt.Println("???", alias)
		}
		port := portRegex.FindString(lines[i+1])

		// fmt.Println(alias, port)

		m[alias] = port
	}

	return m

}

func serializeNginx() []string {
	serialized := []string{}

	for k, v := range m {

		serialized = append(serialized, fmt.Sprintf("\t\tlocation /%s/ {", k))
		serialized = append(serialized, fmt.Sprintf("\t\t\tproxy_pass http://localhost:%s/", v))
		serialized = append(serialized, "\t\t}\n")
	}

	serialized = append(serialized, "\t}")
	serialized = append(serialized, "}")

	return serialized
}

func writeFile(f *os.File, newFile string) {
	f.Truncate(0)
	f.Seek(0, 0)
	fmt.Fprintf(f, "%s", newFile)
}
