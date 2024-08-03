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
	got := counter("test.txt", bufio.ScanBytes)
	want, err := strconv.Atoi(strings.Split(string(output[:]), " ")[0])
	if err != nil {
		t.Fatal("Ran into an error parsing the numbers from `wc`'s output:", err.Error())
	}
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCounterOutputsCorrectNumberOfLines(t *testing.T) {
	cmd := exec.Command("wc", "-l", "test.txt")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal("Ran into an error running `wc -l test.txt`:", err.Error())
	}
	got := counter("test.txt", bufio.ScanLines)
	want, err := strconv.Atoi(strings.Split(string(output[:]), " ")[0])
	want++ // wc -l returns the amount of \n characters, not actual lines, so it's short by 1.
	if err != nil {
		t.Fatal("Ran into an error parsing the numbers from `wc`'s output:", err.Error())
	}
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCounterOutputsCorrectNumberOfWords(t *testing.T) {
	cmd := exec.Command("wc", "-w", "test.txt")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal("Ran into an error running `wc -w test.txt`:", err.Error())
	}
	got := counter("test.txt", bufio.ScanWords)
	want, err := strconv.Atoi(strings.Split(string(output[:]), " ")[0])
	if err != nil {
		t.Fatal("Ran into an error parsing the numbers from `wc`'s output:", err.Error())
	}
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCounterOutputsCorrectNumberOfChars(t *testing.T) {
	cmd := exec.Command("wc", "-m", "test.txt")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal("Ran into an error running `wc -m test.txt`:", err.Error())
	}
	got := counter("test.txt", bufio.ScanRunes)
	want, err := strconv.Atoi(strings.Split(string(output[:]), " ")[0])
	if err != nil {
		t.Fatal("Ran into an error parsing the numbers from `wc`'s output:", err.Error())
	}
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
