package main

import "fmt"

// Абстрактная фабрика
type RestaurantFactory interface {
	CreateFood() Food
	CreateDrink() Drink
}

// Абстрактный интерфейс для еды
type Food interface {
	Eat()
}

// Абстрактный интерфейс для напитков
type Drink interface {
	Drink()
}

// Конкретная фабрика для итальянского ресторана
type ItalianRestaurantFactory struct{}

func (irf ItalianRestaurantFactory) CreateFood() Food {
	return ItalianFood{}
}

func (irf ItalianRestaurantFactory) CreateDrink() Drink {
	return ItalianDrink{}
}

// Конкретная фабрика для японского ресторана
type JapaneseRestaurantFactory struct{}

func (jrf JapaneseRestaurantFactory) CreateFood() Food {
	return JapaneseFood{}
}

func (jrf JapaneseRestaurantFactory) CreateDrink() Drink {
	return JapaneseDrink{}
}

// Конкретные продукты для итальянского ресторана

type ItalianFood struct{}

func (ifd ItalianFood) Eat() {
	fmt.Println("Итальянская еда")
}

type ItalianDrink struct{}

func (idr ItalianDrink) Drink() {
	fmt.Println("Итальянский напиток")
}

// Конкретные продукты для японского ресторана

type JapaneseFood struct{}

func (jfd JapaneseFood) Eat() {
	fmt.Println("Японская еда")
}

type JapaneseDrink struct{}

func (jdr JapaneseDrink) Drink() {
	fmt.Println("Японский напиток")
}

func main() {
	// Создаем фабрику для итальянского ресторана
	italianFactory := ItalianRestaurantFactory{}
	italianFood := italianFactory.CreateFood()
	italianDrink := italianFactory.CreateDrink()

	fmt.Println("Итальянский ресторан:")
	italianFood.Eat()
	italianDrink.Drink()

	// Создаем фабрику для японского ресторана
	japaneseFactory := JapaneseRestaurantFactory{}
	japaneseFood := japaneseFactory.CreateFood()
	japaneseDrink := japaneseFactory.CreateDrink()

	fmt.Println("\nЯпонский ресторан:")
	japaneseFood.Eat()
	japaneseDrink.Drink()
}
