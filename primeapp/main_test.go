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

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout
	}()

	os.Stdout = w

	go readUserInput(bufio.NewReader(r), doneChan)

	w.WriteString("7\n")
	w.WriteString("q\n")

	<-doneChan
	close(doneChan)

	w.Close()

	var buf bytes.Buffer
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		output := scanner.Text()
		if output != "-> 7 is a prime number!" && output != "-> Goodbye." {
			t.Errorf("Test_readUserInput: expected either \"-> 7 is a prime number!\" or \"-> Goodbye.\" but got %s", output)
		}
	}
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