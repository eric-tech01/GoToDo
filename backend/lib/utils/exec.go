package utils

import (
	"os/exec"

	log "github.com/eric-tech01/simple-log"
)

func ExecCommand(cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)
	output, err := c.CombinedOutput()
	log.Debugf("exec command %v: output: %v, err: %v", cmd, string(output), err)
	return string(output), err
}

func ExecShell(cmd string) (string, error) {
	output, err := ExecCommand("sh", "-c", cmd)
	log.Debugf("exec shell %v: output: %v, err: %v", cmd, output, err)
	return output, err
}
