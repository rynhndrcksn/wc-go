package main

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestCounterOutputsCorrectNumberOfBytes(t *testing.T) {
	cmd := exec.Command("wc", "-c", "test.txt")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal("Ran into an error running `wc -c test.txt`:", err.Error())
	}
	got := counter(bufio.ScanBytes)
	want, err := strconv.Atoi(strings.Split(string(output[:]), " ")[0])
	if err != nil {
		t.Fatal("Ran into an error parsing the numbers from `wc`'s output:", err.Error())
	}
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
