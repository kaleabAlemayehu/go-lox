package main

import (
	"bufio"
	"fmt"
	"os"
)

type Lox struct {
	hadError bool
}

func main() {
	l := Lox{}
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	}
	if len(os.Args) == 2 {
		l.runFile(os.Args[1])
	} else {
		l.runPrompt()
	}
}

func (l *Lox) runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	l.run(string(bytes))
	if l.hadError {
		os.Exit(65)
	}
	return nil
}

func (l *Lox) runPrompt() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if ok := input.Scan(); !ok {
			break
		}
		line := input.Text()
		l.run(line)
		l.hadError = false
	}
}

func (l *Lox) run(source string) {
	scanner := Scanner{source}
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (l *Lox) reportError(line int, msg string) {
	l.report(line, "", msg)
}

func (l *Lox) report(line int, where, msg string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, msg)
	l.hadError = true
}
