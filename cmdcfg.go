//go:build !windows
// +build !windows

package wkhtmltopdf

import "os/exec"

func cmdConfig(cmd *exec.Cmd) {}
