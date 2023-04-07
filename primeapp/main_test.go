package main

import (
	"bufio"
	// "bytes"
	// "io"
	// "os"
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

	go readUserInput(strings.NewReader("7\nq\n"), doneChan)

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
		{"quit", "q", "", true},
		{"invalid input", "xyz", "Please enter a whole number!", false},
		{"valid input", "7", "7 is a prime number!", false},
	}

	for _, tc := range testCases {
		result, done := checkNumbers(bufio.NewScanner(strings.NewReader(tc.input)))
		if result != tc.expected {
			t.Errorf("%s: expected result: %s but got: %s", tc.name, tc.expected, result)
		}
		if done != tc.done {
			t.Errorf("%s: expected done: %v but got: %v", tc.name, tc.done, done)
		}
	}
}


