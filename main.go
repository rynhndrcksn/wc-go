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
	countBytes := flag.Bool("c", false, "Count bytes in file")
	countLines := flag.Bool("l", false, "Count lines in file")
	countWords := flag.Bool("w", false, "Count words in file")
	flag.Parse()

	if !*countBytes && !*countLines && !*countWords {
		log.Fatalln("No valid flag options set")
	}

	// The first argument should be a filename.
	filename := flag.Arg(0)
	if filename == "" {
		log.Fatalln("Filename can't be empty")
	}

	// Cycle through any possible flags and print their results.
	if *countBytes {
		fmt.Printf("Byte count = %d\n", counter(filename, bufio.ScanBytes))
	}
	if *countLines {
		fmt.Printf("Line count = %d\n", counter(filename, bufio.ScanLines))
	}
	if *countWords {
		fmt.Printf("Word count = %d\n", counter(filename, bufio.ScanWords))
	}
}

// counter will return the count of however many bufio.SplitFunc there are in the file.
func counter(filename string, split bufio.SplitFunc) int {
	// Try to open the test file, if we can't, bail.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new scanner that will scan over the file.
	s := *bufio.NewScanner(file)
	s.Split(split)
	count := 0
	for s.Scan() {
		count++
	}
	return count
}
