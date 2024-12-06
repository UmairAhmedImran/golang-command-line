package main

import (
	"bufio"
	"fmt"
	"flag"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, bytesSize bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines && !bytesSize {
		scanner.Split(bufio.ScanWords)
	}
  	if bytesSize {
    		scanner.Split(bufio.ScanBytes)
  	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
