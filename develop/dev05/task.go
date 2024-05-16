package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	A := flag.Int("A", 0, "print N lines after match")
	B := flag.Int("B", 0, "print N lines before match")
	C := flag.Int("C", 0, "print N lines around match")
	c := flag.Bool("c", false, "count matching lines")
	i := flag.Bool("i", false, "ignore case")
	v := flag.Bool("v", false, "invert match")
	F := flag.String("F", "", "exact string match")
	n := flag.Bool("n", false, "print line numbers")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: go run task.go <flags(if you use)> <pattern> <file>")
		return
	}

	pattern := os.Args[len(os.Args)-2]
	filename := os.Args[len(os.Args)-1]

	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if *A != 0 {
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], pattern) {
				fmt.Println(lines[i])
				afterNumber := 1
				for k := i + 1; k < len(lines) && afterNumber <= *A; k++ {
					fmt.Println(lines[k])
					afterNumber++
				}
			}
		}
	} else if *B != 0 {
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], pattern) {
				beforeNumber := 1
				for j := i - *B; j >= 0 && beforeNumber <= *B; j++ {
					fmt.Println(lines[j])
					beforeNumber++
				}
				fmt.Println(lines[i])
			}
		}

	} else if *C != 0 {

		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], pattern) {
				beforeNumber := 1
				for j := i - *C; j >= 0 && beforeNumber <= *C; j++ {
					fmt.Println(lines[j])
					beforeNumber++
				}
				fmt.Println(lines[i])
				afterNumber := 1
				for k := i + 1; k < len(lines) && afterNumber <= *C; k++ {
					fmt.Println(lines[k])
					afterNumber++
				}
			}
		}

	} else if *c != false {
		count := 0
		for _, line := range lines {
			if strings.Contains(line, pattern) {
				count++
			}
		}
		fmt.Println(count)
	} else if *i != false {
		lowerPattern := strings.ToLower(pattern)
		for _, line := range lines {
			if strings.Contains(strings.ToLower(line), lowerPattern) {
				fmt.Println(line)
			}
		}
	} else if *v != false {
		for _, line := range lines {
			if !strings.Contains(line, pattern) {
				fmt.Println(line)
			}
		}
	} else if *F != "" {
		for _, line := range lines {
			if line == *F {
				fmt.Println(line)
			}
		}
	} else if *n != false {
		for i, line := range lines {
			if strings.Contains(line, pattern) {
				fmt.Printf("%d: %s\n", i+1, line)
			}
		}
	} else {
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], pattern) {
				fmt.Println(lines[i])
			}
		}
	}

}
