package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func decode(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if str[0] >= '0' && str[0] <= '9' {
		return "", errors.New("invalid string")
	}

	result := ""

	for i := 0; i < len(str); i++ {
		if i == len(str)-1 {
			result += string(str[i])
			//fmt.Printf("%c", str[i])
		} else {
			numStr := ""

			j := i + 1
			for (j != len(str)) && (str[j] >= '0' && str[j] <= '9') {
				numStr += string(str[j])
				j++
			}
			if numStr == "" {
				result += string(str[i])
				//fmt.Printf("%c", str[i])
			} else {
				times, err := strconv.Atoi(numStr)
				if err != nil {
					return "", err
				}
				for times > 0 {
					result += string(str[i])
					//fmt.Printf("%c", str[i])
					times--
				}
			}
			i = j - 1

		}
	}
	return result, nil
}

func main() {
	slices := []string{"", "a", "3a", "ab2c3"}
	for _, val := range slices {
		fmt.Println(decode(val))
	}
}
