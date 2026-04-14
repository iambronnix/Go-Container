package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//docker           run image <cmd> <params>
// go run main.go run        <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}
func run() { //container
	fmt.Printf("Running %v as %v\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...) //exec.Command("/proc/self/exe", "child", os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	syscall.Sethostname([]byte("container"))
	cmd.Run()

}
func child() { //container
	fmt.Printf("Running %v as %v", os.Args[2:], os.Getpid())
	syscall.Sethostname([]byte("container")) //syscall.Mount("proc", "proc", "proc", 0, "")
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	syscall.Chroot("/helmsman/ubuntu-fs")
	syscall.Chdir("/")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	syscall.Unmount("/proc", 0)
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
