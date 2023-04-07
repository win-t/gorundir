package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	originalDir, err := os.Getwd()
	check(err)

	targetDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		targetDir = os.Args[1]
	}

	targetSum256 := sha256.Sum256([]byte(targetDir))
	targetName := hex.EncodeToString(targetSum256[:])

	targetFullPath := filepath.Join(os.TempDir(), "gorundir", targetName)

	err = os.Chdir(targetDir)
	check(err)

	goBuild := exec.Command("go", "build", "-o", targetFullPath, ".")
	goBuild.Stdin, goBuild.Stdout, goBuild.Stderr = nil, os.Stdout, os.Stderr
	err = goBuild.Run()
	check(err)

	err = os.Chdir(originalDir)
	check(err)

	err = syscall.Exec(targetFullPath, os.Args[1:], os.Environ())
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
