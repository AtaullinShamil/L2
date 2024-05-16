package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagramGroups(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		added := false

		word = strings.ToLower(word)
		sorted := sortLetters(word)
		for key, _ := range groups {
			sortedKey := sortLetters(key)
			if sortedKey == sorted {
				groups[key] = append(groups[key], word)
				added = true
				break
			}
		}
		if added == true {
			continue
		} else {
			groups[word] = append(groups[word], word)
		}
	}

	for key := range groups {
		sort.Strings(groups[key])
		groups[key] = removeDuplicates(groups[key])
		if len(groups[key]) == 1 {
			delete(groups, key)
		}
	}

	return groups
}

func sortLetters(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func removeDuplicates(words []string) []string {
	seen := make(map[string]struct{}, len(words))
	j := 0
	for _, v := range words {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		words[j] = v
		j++
	}
	return words[:j]
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	result := findAnagramGroups(words)
	fmt.Println(result)
}
