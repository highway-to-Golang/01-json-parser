package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := flag.String("file", "", "")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Usage: json-parser -file <filename>")
		os.Exit(1)
	}

	data, _ := os.ReadFile(*filename)
	var jsonData map[string]interface{}
	json.Unmarshal(data, &jsonData)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter keys (one per line, empty line to exit):")

	for {
		fmt.Print("> ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSpace(key)

		if key == "" {
			break
		}

		if value, exists := jsonData[key]; exists {
			fmt.Printf("%s: %v\n", key, value)
		} else {
			fmt.Printf("Key '%s' not found\n", key)
		}
	}
}
