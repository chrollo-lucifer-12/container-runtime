package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func cg() {
	cgroupPath := "/sys/fs/cgroup/minicontainer"
	os.MkdirAll(cgroupPath, 0755)

	must(os.WriteFile(
		cgroupPath+"/pids.max",
		[]byte("100"),
		0700,
	))

	must(os.WriteFile(
		cgroupPath+"/cpu.max",
		[]byte("50000 100000"),
		0700,
	))

	must(os.WriteFile(cgroupPath+"/memory.max", []byte("100000000"), 0700))

	must(os.WriteFile(cgroupPath+"/cgroup.procs", []byte(fmt.Sprintf("%d", os.Getpid())), 0700))
}

func child() {

	cg()

	must(syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""))

	must(exec.Command("ip", "link", "set", "lo", "up").Run())
	must(syscall.Sethostname([]byte("container")))

	must(syscall.Chroot("./rootfs"))

	must(syscall.Chdir("/"))

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
