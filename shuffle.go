package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	infile := "in.txt"
	if len(os.Args) > 1 {
		infile = os.Args[1]
	}

	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines),
		func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })

	outfile, err := os.Create(infile + ".out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = outfile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count := 0
	total := 0
	for _, line := range lines {
		count, err = fmt.Fprintln(outfile, line)
		if err != nil {
			log.Fatal(err)
		}
		total += count
	}
	fmt.Printf("%d bytes written.", total)
}
