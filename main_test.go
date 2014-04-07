package main

import (
	"bytes"
	"os/exec"
	"testing"
)

type (
	CommandResult struct {
		stdout string
		stderr string
	}
)

func runCommandAndCaptureStdoutAndStderr(cmd string, args ...string) (CommandResult, error) {
	command := exec.Command(cmd, args...)
	stderr, err := command.StderrPipe()
	if err != nil {
		return CommandResult{}, err
	}
	stdout, err := command.StdoutPipe()
	if err != nil {
		return CommandResult{}, err
	}
	err = command.Start()
	if err != nil {
		return CommandResult{}, err
	}
	errBuf := new(bytes.Buffer)
	errBuf.ReadFrom(stderr)
	outBuf := new(bytes.Buffer)
	outBuf.ReadFrom(stdout)
	command.Wait()
	return CommandResult{outBuf.String(), errBuf.String()}, nil
}

func TestHelpMessageWithHelpArgument(t *testing.T) {
	result, err := runCommandAndCaptureStdoutAndStderr("./gocat", "-h")
	if err != nil {
		t.Fatalf("Execution error: %s", err)
	}
	expected := "usage: ./gocat [file]\n"
	if result.stderr != expected {
		t.Errorf("'%s' != '%s'", result.stderr, expected)
	}
	if result.stdout != "" {
		t.Error("Expected STDOUT to be empty")
	}
}

func TestReadFileAndOutputStdout(t *testing.T) {
	result, err := runCommandAndCaptureStdoutAndStderr("./gocat", "tests/testfile.txt")
	if err != nil {
		t.Fatalf("Execution error: %s", err)
	}
	expected := "Hello world!\n"
	if result.stdout != expected {
		t.Errorf("'%s' != '%s'", result.stdout, expected)
	}
	if result.stderr != "" {
		t.Error("Expected STDERR to be empty")
	}
}

func TestReadStdingAndOutputStdout(t *testing.T) {
	result, err := runCommandAndCaptureStdoutAndStderr("sh", "-c", "cat tests/testfile.txt | ./gocat")
	if err != nil {
		t.Fatalf("Execution error: %s", err)
	}
	expected := "Hello world!\n"
	if result.stdout != expected {
		t.Errorf("'%s' != '%s'", result.stdout, expected)
	}
	if result.stderr != "" {
		t.Error("Expected STDERR to be empty")
	}
}
