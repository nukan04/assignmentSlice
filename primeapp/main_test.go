package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}


func Test_checkNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		done     bool
	}{
		{
			name:     "quit",
			input:    "q",
			expected: "",
			done:     true,
		},
		{
			name:     "invalid input",
			input:    "abc",
			expected: "Please enter a whole number!",
			done:     false,
		},
		{
			name:     "valid input",
			input:    "5",
			expected: "5 is a prime number!",
			done:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tc.input))
			result, done := checkNumbers(scanner)
			if result != tc.expected {
				t.Errorf("expected result: %s, got: %s", tc.expected, result)
			}
			if done != tc.done {
				t.Errorf("expected done: %v, got: %v", tc.done, done)
			}
		})
	}
}



func Test_intro(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	output := captureOutput(w, r)

	expected := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "
	if output != expected {
		t.Errorf("intro: expected %s but got %s", expected, output)
	}
}

func Test_prompt(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	output := captureOutput(w, r)

	expected := "-> "
	if output != expected {
		t.Errorf("prompt: expected %s but got %s", expected, output)
	}
}

func captureOutput(w *os.File, r *os.File) string {
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
