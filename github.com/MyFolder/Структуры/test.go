package main

import (
	"fmt"
)

type Collor struct {
	Number  int
	Name    string
	Element string
}

// Семантика метода – это описание того, что данный метод делает.
func (c *Collor) AddCollor(number int, name string) { // (c *Collor) - ресивер(указатель на будущий объект). Семантика метода: объявление функции + ресивер + функци с принимаемыми значениями (функция структуры Collor)
	c.Name = name
	c.Number = number
	return
}

type House struct {
	Material     string // material - поле
	SquerMetrics int32
	Flats        int32
	Collors      []Collor
}

func (h *House) GetMaterial() string {
	return h.Material
}

const A string = "sdfsdfs"

func main() {

	var house1 House // объект/экземпляр
	fmt.Println(house1.Material)

	wood := A
	var collorBlack Collor
	collorBlack.AddCollor(10, "black")
	var collorWhite Collor
	collorWhite.AddCollor(10, "white")
	var collors []Collor
	lenght := len(collors)
	fmt.Println(lenght)
	collors = append(collors, collorWhite, collorBlack)
	//collor2 = AddCollor(10, "black") // вызываем функцию, и возвращаем (копируя из result) значение в переменную collor2
	house2 := House{
		Material:     wood,
		SquerMetrics: 100,
		Flats:        2,
		//Collors:      []Collor{Collor{Number: 10, Name: "black", Element: "door"}},
		Collors: collors,
	}
	fmt.Println(house2)
}

// Создать новую структуру Slice (Должны быть поля вместимость Capacity, поле кол-во объектов (lenght), поле указатель на массив) // для упрощения сначала без указателя
// Создать метод для записи в массив (Append)
// Создать метод для получение вместимости
// Создать метод для получение кол-ва объектов
// Создать метод для вывода слайса в консоль
// *Создать метод для сортировки пузырьком
// ***Создать метод для удаления элемента из слайса
// ****использовать interface как пустой тип данных. Почитать про upcast и downcast - приведение данных
