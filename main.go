package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("docker")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println(binary)
	args := []string{"exec", "axiom_tool", "/root/.axiom/interact/axiom-images", "get"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
