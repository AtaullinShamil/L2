Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Это происходит, потому что интерфейсы реализуются в виде двух элементов, T типа и V значения.
V - это конкретное значение, такое как int. Значение V также известно как динамическое значение интерфейса.
Значение интерфейса равно nil только в том случае, если оба параметра V и T не заданы (T=nil, V не задано).
```
