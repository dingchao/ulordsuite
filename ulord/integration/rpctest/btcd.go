// Copyright (c) 2017 The ulordsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpctest

import (
	"fmt"
	"go/build"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	// compileMtx guards access to the executable path so that the project is
	// only compiled once.
	compileMtx sync.Mutex

	// executablePath is the path to the compiled executable. This is the empty
	// string until ulord is compiled. This should not be accessed directly;
	// instead use the function ulordExecutablePath().
	executablePath string
)

// ulordExecutablePath returns a path to the ulord executable to be used by
// rpctests. To ensure the code tests against the most up-to-date version of
// ulord, this method compiles ulord the first time it is called. After that, the
// generated binary is used for subsequent test harnesses. The executable file
// is not cleaned up, but since it lives at a static path in a temp directory,
// it is not a big deal.
func ulordExecutablePath() (string, error) {
	compileMtx.Lock()
	defer compileMtx.Unlock()

	// If ulord has already been compiled, just use that.
	if len(executablePath) != 0 {
		return executablePath, nil
	}

	testDir, err := baseDir()
	if err != nil {
		return "", err
	}

	// Determine import path of this package. Not necessarily ulordsuite/ulord if
	// this is a forked repo.
	_, rpctestDir, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("Cannot get path to ulord source code")
	}
	ulordPkgPath := filepath.Join(rpctestDir, "..", "..", "..")
	ulordPkg, err := build.ImportDir(ulordPkgPath, build.FindOnly)
	if err != nil {
		return "", fmt.Errorf("Failed to build ulord: %v", err)
	}

	// Build ulord and output an executable in a static temp path.
	outputPath := filepath.Join(testDir, "ulord")
	if runtime.GOOS == "windows" {
		outputPath += ".exe"
	}
	cmd := exec.Command("go", "build", "-o", outputPath, ulordPkg.ImportPath)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Failed to build ulord: %v", err)
	}

	// Save executable path so future calls do not recompile.
	executablePath = outputPath
	return executablePath, nil
}
