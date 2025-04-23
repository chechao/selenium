//go:build !windows
// +build !windows

package selenium

import "os/exec"

func setHideWindow(cmd *exec.Cmd) {
	// 非Windows平台无操作
}
