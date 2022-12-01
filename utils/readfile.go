package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fileName string) []string {
	var lines []string

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf(`Error reading file %s.`, fileName)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
