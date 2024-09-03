package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		panic("CFS-kkZMmhqM9x Usage: go run main.go [run|child] <cmd> <args>")
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("CFS-Sfmm9NWZeA Invalid command. Available commands:\n" +
			"\t'run'    : Creates a new process in a containerized environment.\n" +
			"\t'child'  : Runs the specified command in the isolated environment.")
	}
}

func run() {
	fmt.Printf("Running %v as pid %d (run)\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Use platform-specific function to set SysProcAttr
	setSysProcAttr(cmd)

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Failed to start child process: %v", err))
	}
}

// Inside the namespace
func child() {
	fmt.Printf("Running %v as pid %d (child)\n", os.Args[2:], os.Getpid())

	// Use platform-specific function to set Hostname
	if err := setHostName(); err != nil {
		panic(fmt.Sprintf("Failed to set hostname: %v", err))
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Failed to run command in child process: %v", err))
	}
}
