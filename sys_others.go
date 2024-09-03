// +build !linux

package main

import (
    "errors"
    "os/exec"
)

func setSysProcAttr(cmd *exec.Cmd) {
    // No-op on non-Linux systems
}

func setHostName() error {
    return errors.New("setting hostname is not supported on this platform")
}
