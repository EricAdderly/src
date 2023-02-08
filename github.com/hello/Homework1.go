package main

import (
	"fmt"
)

const baseCapacity int = 8 // константа для массива из 8
const capacity16 int = 16  // константа для массива из 16

type Slice struct { //объявляем структуру нашего "слайса"
	Array    interface{}
	Capacity int
	Lenght   int
}

func (s *Slice) AddElement(value int) { //Функция добавления значений в "слайс"
	if s.Array == nil {
		s.init()
	}

	s.addToSlice(value)

}

func (s *Slice) init() { // функция присвоение значений полям из структуры
	s.Capacity = baseCapacity
	s.Array = [baseCapacity]*int{}
	s.initLenght()
}

func (s *Slice) initLenght() { // функция, для подсчёта длины массива
	count := 0
	switch s.Capacity {
	case baseCapacity:
		arr8 := s.Array.([baseCapacity]*int)
		for _, v := range arr8 {
			if v != nil {
				count++
			}
		}
	case capacity16:
		arr16 := s.Array.([capacity16]*int)
		for _, v := range arr16 {
			if v != nil {
				count++
			}
		}
	}
	s.Lenght = count
}
func (s *Slice) addToSlice(value int) { // запись переданных значений в массив
	switch s.Capacity {
	case baseCapacity:
		arr8 := s.Array.([baseCapacity]*int) // приведение к массиву

		if s.Lenght >= baseCapacity { // проверка что длина больше капасити
			newArr16 := [capacity16]*int{}
			for i := 0; i < baseCapacity; i++ { // присвоение значений из более раннего массива в новый
				newArr16[i] = arr8[i]
			}

			s.initLenght() // подсчёт длины
			fmt.Println(s.Lenght)
			newArr16[s.Lenght+1] = &value //запись в массив
			s.Array = newArr16
		}

		s.initLenght()            // подсчёт длины
		arr8[s.Lenght+1] = &value // запись в массив
		s.Array = arr8
	}

}

func main() {
	var slice1 Slice
	slice1.AddElement(1)
	slice1.AddElement(2)
	fmt.Println(slice1)
}
