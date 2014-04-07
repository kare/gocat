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
	defer stderr.Close()
	if err != nil {
		return CommandResult{}, err
	}
	stdout, err := command.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		return CommandResult{}, err
	}
	err = command.Start()
	if err != nil {
		return CommandResult{}, err
	}
	errBuf := new(bytes.Buffer)
	_, err = errBuf.ReadFrom(stderr)
	if err != nil {
		panic(err)
	}
	outBuf := new(bytes.Buffer)
	_, err = outBuf.ReadFrom(stdout)
	if err != nil {
		panic(err)
	}
	if err := command.Wait(); err != nil {
		panic(err)
	}
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
		t.Errorf("Expected STDERR to be empty: '%s'", result.stderr)
	}
}

var avoid_compiler_optimization CommandResult

func runBenchmarkSimpleCatToStdout(b *testing.B, input string) {
	for n := 0; n < b.N; n++ {
		result, err := runCommandAndCaptureStdoutAndStderr("./gocat", input)
		if err != nil {
			b.Fatalf("Execution error: %s", err)
		}
		avoid_compiler_optimization = result
	}
}

func BenchmarkSimpleCatToStdoutTestfile(b *testing.B) {
	runBenchmarkSimpleCatToStdout(b, "tests/testfile.txt")
}

func BenchmarkSimpleCatToStdoutEtcPasswd(b *testing.B) {
	runBenchmarkSimpleCatToStdout(b, "/etc/passwd")
}

func BenchmarkSimpleCatToStdoutLargeFile(b *testing.B) {
	runBenchmarkSimpleCatToStdout(b, "tests/rand.txt")
}
