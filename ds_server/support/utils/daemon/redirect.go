// +build !windows

package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var stderr *os.File

func getProgName() string {
	fullPath, _ := exec.LookPath(os.Args[0])
	fname := filepath.Base(fullPath)

	return fname
}

func redirectStderr() {
	fmt.Println("redirect stderr in unix mode")
	filename := fmt.Sprintf("./logex/stderr_%v.logex", getProgName())
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	stderr = file
	if err != nil {
		fmt.Printf("open redirect file failed, err = %v\n", err)
		return
	}
	if err = syscall.Dup2(int(stderr.Fd()), int(os.Stderr.Fd())); err != nil {
		fmt.Printf("redirect failed, cannot redirect panic info, err = %v\n", err)
		return
	}
}
