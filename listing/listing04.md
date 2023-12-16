Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

Это происходит, потому что цикл for, ожидающий данные из канала, 
не завершается, после ожидания горутины.
Канал не закрывается, и цикл продолжает ожидание данных.
```
