package config_test

import (
	"AnonimousChat/pkg/config"
	"fmt"
	"os"
	"testing"
)

type EnvData struct {
	Key   string
	Value string
}

func (d EnvData) String() string {
	return fmt.Sprintf("%s=%s\n", d.Key, d.Value)
}

func TestErrorNotExistsFile(t *testing.T) {
	err := config.LoadEnvArgsFromFile("")
	if err == nil {
		t.Fatal("Expected error: file not exists, got nil\n")
	}
}

func TestLoadEnvArgsForDontCorrectLines(t *testing.T) {
	file := newTempFile(t)

	testCases := []EnvData{
		{"", ""},
		{"key", ""},
		{"", "value"},
		{"key=", "value"},
		{"=key", "value"},
		{"key", "value="},
		{"key=", "=value"},
		{"=key=", "=value="},
	}

	for _, test := range testCases {
		file.WriteString(test.String())
	}

	err := config.LoadEnvArgsFromFile(file.Name())
	if err != nil {
		t.Fatalf("Expected nil, got error: %v", err)
	}

	for _, test := range testCases {
		if _, exists := os.LookupEnv(test.Key); exists {
			t.Errorf("Key %s exists for enviroment, key and value => dont correct line\n", test.Key)
		}
	}
}

func newTempFile(t *testing.T) *os.File {
	dir := os.TempDir()
	file, err := os.CreateTemp(dir, "*")
	if err != nil {
		t.Fatalf("Error creating temp fail, error: %v", err)
	}

	return file
}

func TestLoadEnvArgsForCorrectLines(t *testing.T) {
	file := newTempFile(t)

	testCases := []EnvData{
		{"a", "a"},
		{"abc", "qwerty"},
		{"qwerty", "dbc"},
		{"key", "value"},
		{"test_key", "test_value"},
		{"PG_USER", "TEST_USER"},
		{"ARG", "S"},
		{"q", "..."},
		{"/q", "yes"},
	}

	for _, test := range testCases {
		file.WriteString(test.String())
	}

	err := config.LoadEnvArgsFromFile(file.Name())
	if err != nil {
		t.Fatalf("Expected nil, got error: %v", err)
	}

	for _, test := range testCases {
		val, exists := os.LookupEnv(test.Key)
		if exists && test.Value != val {
			t.Errorf("Key %s, expected value = %s, got value = %s\n", test.Key, test.Value, val)
		}

		if !exists {
			t.Errorf("Key %s, expected exists value\n", test.Key)
		}
	}
}
