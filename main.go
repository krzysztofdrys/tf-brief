package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/krzysztofdrys/tf-brief/brief"
)

func processManyTimes(lines []string) []string {
	for {
		newLines := brief.Plan(lines)
		if len(newLines) == len(lines) {
			return newLines
		}
		lines = newLines
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return
	}

	result := processManyTimes(lines)
	for _, line := range result {
		fmt.Println(line)
	}
}
