package main

import (
	"os/exec"
)

func main() {
	c := exec.Command("plink", "-l", "root", "-P", "22", "172.16.6.26", "-pw", "xinge")
	c.StdoutPipe()
}
