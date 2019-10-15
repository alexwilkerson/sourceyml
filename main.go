package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("wrong number of arguments\nusage: sourceyml [src]")
	}
	src := os.Args[1]

	if !fileExists(src) {
		log.Fatal("file does not exist")
	}

	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#' {
			continue
		}
		keyValue := strings.SplitN(strings.SplitN(line, "#", 2)[0], ": ", 2)
		if len(keyValue) != 2 {
			continue
		}
		key := keyValue[0]
		value := keyValue[1]
		fmt.Printf("export %s=\"%s\"\n", key, value)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
