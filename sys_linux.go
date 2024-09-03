// +build linux

package main

import (
    "os/exec"
    "syscall"
)

func setSysProcAttr(cmd *exec.Cmd) {
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
}

func setHostName() error {
    return syscall.Sethostname([]byte("container"))
}