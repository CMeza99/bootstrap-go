package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

import (
	"bytes"
_	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func CmdRun(c string, args[]string) (error) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("ls", "-lah")

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	defer stdoutIn.Close()
	defer stderrIn.Close()

	var errStdout, errStderr error
	// TODO: Should change to write stderr and stdout if --debug and always to log.
	//       Right now, the *Buf vars are not actually being used.
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		return err
		//log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		return err
		//log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		return errStderr
		//log.Fatal("failed to capture stdout or stderr\n")
	}
	//outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	//fmt.Sprintf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return errStderr
}

func main() {
	args := []string{"-lah"}
	CmdRun("ls", args )

}
