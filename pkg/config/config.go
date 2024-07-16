package config

import (
	"log"
	"os"
	"strings"
)

func LoadEnvArgsFromFile(name string) error {
	b, err := os.ReadFile(name)
	if err != nil {
		log.Printf("Error load env args from file, error: %v\n", err)
		return err
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		loadEnvArg(line)
	}

	return nil
}

func loadEnvArg(line string) {
	fields := strings.Split(line, "=")
	if len(fields) != 2 || len(fields[0]) == 0 || len(fields[1]) == 0 {
		return
	}

	os.Setenv(fields[0], fields[1])
}
