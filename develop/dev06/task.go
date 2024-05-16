package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var fields string
	var delimiter string
	var separated bool

	flag.StringVar(&fields, "f", "", "Fields")
	flag.StringVar(&delimiter, "d", "\t", "Delimiter")
	flag.BoolVar(&separated, "s", false, "Separated")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: go run task.go <flags> <file>")
		return
	}
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if separated && !strings.Contains(line, delimiter) {
			continue
		}
		cols := strings.Split(line, delimiter)
		if len(fields) > 0 {
			selectedCols := selectColumns(cols, fields)
			fmt.Println(strings.Join(selectedCols, delimiter))
		} else {
			fmt.Println(line)
		}
	}
}

func selectColumns(cols []string, fields string) []string {
	fieldIndices := parseFieldIndices(fields)
	if len(fieldIndices) > len(cols) {
		fieldIndices = fieldIndices[:len(cols)]
	}

	selectedCols := make([]string, len(fieldIndices))
	for i, idx := range fieldIndices {
		selectedCols[i] = cols[idx-1]
	}
	return selectedCols
}

func parseFieldIndices(fieldStr string) []int {
	fields := strings.Split(fieldStr, ",")
	fieldIndices := make([]int, len(fields))
	for i, f := range fields {
		fieldIndices[i], _ = strconv.Atoi(f)
	}
	return fieldIndices
}
