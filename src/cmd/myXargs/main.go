package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myXargs команда [аргументы]")
		os.Exit(1)
	}

	cmdParts := os.Args[1:] // Команда и её аргументы, переданные в программу

	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			executeCommand(cmdParts, line)
		}(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения стандартного ввода: %s\n", err)
		os.Exit(1)
	}

	wg.Wait() // Ожидаем завершения всех горутин
}

func executeCommand(cmdParts []string, argsFromInput string) {
	additArgs := strings.Fields(argsFromInput)
	cmd := exec.Command(cmdParts[0], append(cmdParts[1:], additArgs...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка выполнения команды: %s\n", err)
	}
}
