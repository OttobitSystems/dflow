package main

import (
	"testing"
)

// Test that the cli has the enter command
func TestEnterCommandExists(t *testing.T) {
	// Find the "enter" command in the root command
	command, _, err := rootCmd.Find([]string{"enter"})
	if err != nil {
		t.Fatalf("Failed to find 'enter' command: %v", err)
	}
	if command == nil {
		t.Fatal("Expected 'enter' command to exist, but it was not found")
	}
	if command.Use != "enter" {
		t.Fatalf("Expected command use to be 'enter', got '%s'", command.Use)
	}
}
