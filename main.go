package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Set up available CLI flags.
	countByte := flag.Bool("c", false, "Count bytes in file")
	flag.Parse()

	if !*countByte {
		log.Fatalln("No valid flag options set")
	}

	// The first argument should be a filename.
	filename := flag.Arg(0)
	if filename == "" {
		log.Fatalln("Filename can't be empty")
	}

	byteCount := counter(bufio.ScanBytes)
	fmt.Println("Byte count =", byteCount)
}

// counter will return the count of however many bufio.SplitFunc there are in the file.
func counter(split bufio.SplitFunc) int {
	// Try to open the test file, if we can't, bail.
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new scanner that will scan over the file.
	s := bufio.NewScanner(file)
	s.Split(split)
	count := 0
	for s.Scan() {
		count++
	}
	return count
}
