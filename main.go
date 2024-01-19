package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("axiom-images")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println(binary)
	args := []string{"get"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
