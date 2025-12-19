package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log/slog"
	"os"
	"strings"
)

func main() {
	filename := flag.String("file", "", "")
	flag.Parse()

	if *filename == "" {
		slog.Info("Usage: json-parser -file <filename>")
		os.Exit(1)
	}

	data, err := os.ReadFile(*filename)
	if err != nil {
		slog.Error("Error reading file", "error", err)
		os.Exit(1)
	}

	var jsonData = make(map[string]any)

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		slog.Error("Error parsing JSON", "error", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	slog.Info("Enter keys (one per line, empty line to exit):")

	for {
		key, err := reader.ReadString('\n')
		if err != nil {
			slog.Error("Error reading input", "error", err)
			continue
		}
		key = strings.TrimSpace(key)

		if key == "" {
			break
		}

		if value, ok := jsonData[key]; ok {
			slog.Info("Key found", "key", key, "value", value)
		} else {
			slog.Warn("Key not found", "key", key)
		}
	}
}
