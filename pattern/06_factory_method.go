package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Плюсы: Делает код более гибким и расширяемым, скрывает детали реализации от клиентского кода.
// Минусы: Может усложнить структуру кода, если используется слишком часто или в неподходящих случаях.

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

//func main() {
//	ak47, _ := getGun("ak47")
//	musket, _ := getGun("musket")
//
//	printDetails(ak47)
//	printDetails(musket)
//}
//
//func printDetails(g IGun) {
//	fmt.Printf("Gun: %s", g.getName())
//	fmt.Println()
//	fmt.Printf("Power: %d", g.getPower())
//	fmt.Println()
//}
