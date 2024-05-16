package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func readFile(filename string) ([]string, error) {
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
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func sortLines(lines []string, column int, numeric bool, reverse bool, unique bool) []string {
	sort.Slice(lines, func(i, j int) bool {
		a := lines[i]
		b := lines[j]
		if column > 0 {
			parts := strings.Split(a, " ")
			if len(parts) > column-1 {
				a = parts[column-1]
			} else {
				a = parts[len(parts)-1]
			}

			parts = strings.Split(b, " ")
			if len(parts) > column-1 {
				b = parts[column-1]
			} else {
				b = parts[len(parts)-1]
			}
		}
		if numeric {
			numA, err := strconv.Atoi(a)
			if err != nil {
				log.Fatalf("Failed to convert %s to integer: %v", a, err)
			}
			numB, err := strconv.Atoi(b)
			if err != nil {
				log.Fatalf("Failed to convert %s to integer: %v", a, err)
			}
			if numA < numB {
				return !reverse
			}
			if numA > numB {
				return reverse
			}
		} else {
			if a < b {
				return !reverse
			}
			if a > b {
				return reverse
			}
		}
		return false
	})
	if unique {
		seen := make(map[string]bool)
		var uniqueLines []string
		for _, line := range lines {
			if !seen[line] {
				seen[line] = true
				uniqueLines = append(uniqueLines, line)
			}
		}
		return uniqueLines
	}
	return lines
}

func main() {
	var column int
	var numeric bool
	var reverse bool
	var unique bool
	flag.IntVar(&column, "k", 0, "columns")
	flag.BoolVar(&numeric, "n", false, "numeric")
	flag.BoolVar(&reverse, "r", false, "reverse")
	flag.BoolVar(&unique, "u", false, "unique")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("There are no files to sort")
		return
	}
	for _, file := range files {
		lines, err := readFile(file)
		if err != nil {
			log.Fatal(err)
		}
		sortedLines := sortLines(lines, column, numeric, reverse, unique)
		for _, val := range sortedLines {
			fmt.Println(val)
		}
	}
}
