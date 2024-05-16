package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Плюсы: изоляция сложности,гибкость, расширяемость.
// Минусы: большое количество объектов, трудности в поддержке, отсутствие единой точки входа.

const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollectorType(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (c *Computer) Print() {
	fmt.Printf("%s : Core:[%d], Memory:[%d], GraphicCard:[%d], Monitor:[%d]\n", c.Brand, c.Core, c.Memory, c.GraphicCard, c.Monitor)
}

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (a *AsusCollector) SetCore() {
	a.Core = 4
}

func (a *AsusCollector) SetBrand() {
	a.Brand = "Asus"
}

func (a *AsusCollector) SetMemory() {
	a.Memory = 8
}

func (a *AsusCollector) SetMonitor() {
	a.Monitor = 1
}

func (a *AsusCollector) SetGraphicCard() {
	a.GraphicCard = 1
}

func (a *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        a.Core,
		Brand:       a.Brand,
		Memory:      a.Memory,
		Monitor:     a.Monitor,
		GraphicCard: a.GraphicCard,
	}
}

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (h *HpCollector) SetCore() {
	h.Core = 8
}

func (h *HpCollector) SetBrand() {
	h.Brand = "HP"
}

func (h *HpCollector) SetMemory() {
	h.Memory = 16
}

func (h *HpCollector) SetMonitor() {
	h.Monitor = 1
}

func (h *HpCollector) SetGraphicCard() {
	h.GraphicCard = 1
}

func (h *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        h.Core,
		Brand:       h.Brand,
		Memory:      h.Memory,
		Monitor:     h.Monitor,
		GraphicCard: h.GraphicCard,
	}
}

type Factory struct {
	Collector Collector
}

func NewFactory(Collector Collector) *Factory {
	return &Factory{Collector: Collector}
}

func (f *Factory) SetCollector(collector Collector) {
	f.Collector = collector
}

func (f *Factory) CreateComputer() Computer {
	f.Collector.SetCore()
	f.Collector.SetMemory()
	f.Collector.SetBrand()
	f.Collector.SetGraphicCard()
	f.Collector.SetMonitor()
	return f.Collector.GetComputer()
}

//func main() {
//	asus := pattern.GetCollectorType("asus")
//	hp := pattern.GetCollectorType("hp")
//
//	factory := pattern.NewFactory(asus)
//
//	asusComputer := factory.CreateComputer()
//	asusComputer.Print()
//
//	factory.SetCollector(hp)
//
//	hpComputer := factory.CreateComputer()
//	hpComputer.Print()
//}
