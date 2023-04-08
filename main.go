package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no directory is specified")
		os.Exit(1)
	}
	targetDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		targetDir = os.Args[1]
	}

	nameParts := strings.Split(targetDir, string(os.PathSeparator))
	for i := 0; i < len(nameParts)-1; i++ {
		if len(nameParts[i]) > 6 {
			nameParts[i] = nameParts[i][:4] + ".."
		}
	}
	if len(nameParts) > 0 && nameParts[0] == "" {
		nameParts = nameParts[1:]
	}

	targetFullPath := filepath.Join(os.TempDir(), "gorundir", strings.Join(nameParts, "+"))

	goBuild := exec.Command("go", "build", "-o", targetFullPath, "-C", targetDir, ".")
	goBuild.Stdin, goBuild.Stdout, goBuild.Stderr = nil, os.Stdout, os.Stderr
	err = goBuild.Run()
	check(err)

	err = syscall.Exec(targetFullPath, os.Args[1:], os.Environ())
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
