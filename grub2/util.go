package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func quoteString(str string) string {
	return strconv.Quote(str)
}

func unquoteString(str string) string {
	if (strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`)) ||
		(strings.HasPrefix(str, `'`) && strings.HasSuffix(str, `'`)) {
		return str[1 : len(str)-1]
	}
	return str
}

func execAndWait(timeout int, name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		logError(err.Error()) // TODO
	}

	// wait for process finish
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			logError(err.Error()) // TODO
		}
		<-done
		logInfo("process killed") // TODO
	case err := <-done:
		logInfo("process output: %s", stdout.String())
		if err != nil {
			logError("process done with error = %v", err) // TODO
			logInfo("process error output: %s", stderr.String())
		}
	}

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
