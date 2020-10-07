package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	output, err := exec.Command("uwsgi", "--version").Output()
	if err != nil {
		t.Fatal(err)
	}

	s := strings.TrimRight(string(output), "\n")
	if s != "2.0.19.1" {
		t.Errorf("want 2.0.19.1, got %v", s)
	}
}
