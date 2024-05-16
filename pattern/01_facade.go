package pattern

import (
	"errors"
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Плюсы: изолирование клиента от поведения сложной структуры поведений
// Минусы: фасад может стать супер-объектом, и все функции будут проходить через этот объект

type Product struct {
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

func (s Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Second / 10)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка может ли %s купить товар\n", user.Name)
	time.Sleep(time.Second / 10)

	for _, prod := range s.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}
		fmt.Printf("[Магазин] товар %s куплен\n", prod.Name)
	}
	return nil
}

type Bank struct {
	Name  string
	Cards []Card
}

func (b Bank) CheckBalance(cardNumber string) error {
	fmt.Println("[Банк] Получение остатка по карте")
	time.Sleep(time.Second / 10)

	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
	}
	fmt.Println("[Банк] Остаток положительный!")
	return nil
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для получения остатка")
	time.Sleep(time.Second / 10)
	return c.Bank.CheckBalance(c.Name)
}

type User struct {
	Name string
	Card *Card
}

func (u User) GetBalance() float64 {
	return u.Card.Balance
}

//var (
//	bank = pattern.Bank{
//		Name:  "Банк",
//		Cards: []pattern.Card{},
//	}
//	card1 = pattern.Card{
//		Name:    "CRD-1",
//		Balance: 200,
//		Bank:    &bank,
//	}
//	card2 = pattern.Card{
//		Name:    "CRD-2",
//		Balance: 5,
//		Bank:    &bank,
//	}
//	user1 = pattern.User{
//		Name: "Покупатель-1",
//		Card: &card1,
//	}
//	user2 = pattern.User{
//		Name: "Покупатель-2",
//		Card: &card2,
//	}
//	prod = pattern.Product{
//		Name:  "Сыр",
//		Price: 150,
//	}
//	shop = pattern.Shop{
//		Name: "",
//		Products: []pattern.Product{
//			prod,
//		},
//	}
//)
//
//func main() {
//	fmt.Println("[Банк] Выпуск карты")
//	bank.Cards = append(bank.Cards, card1, card2)
//	fmt.Printf("[%s]\n", user1.Name)
//	err := shop.Sell(user1, prod.Name)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = shop.Sell(user2, prod.Name)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
