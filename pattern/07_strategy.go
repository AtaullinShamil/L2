package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Плюсы: Декомпозиция поведения, что позволяет изменять или добавлять новые стратегии без изменения класса, который их использует.
// Минусы: Может привести к большому количеству классов, если у вас много различных стратегий.

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

//func main() {
//	lfu := &Lfu{}
//	cache := initCache(lfu)
//
//	cache.add("a", "1")
//	cache.add("b", "2")
//
//	cache.add("c", "3")
//
//	lru := &Lru{}
//	cache.setEvictionAlgo(lru)
//
//	cache.add("d", "4")
//
//	fifo := &Fifo{}
//	cache.setEvictionAlgo(fifo)
//
//	cache.add("e", "5")
//
//}
