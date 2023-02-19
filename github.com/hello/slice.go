package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const baseCapacity int = 8 // константа для массива из 8
const capacity16 int = 16  // константа для массива из 16

type Slice struct { //объявляем структуру нашего "слайса"
	Array    interface{}
	Capacity int
	Lenght   int
}

func (s *Slice) AddElement(value int) error { //Функция добавления значений в "слайс"
	if s.Array == nil {
		s.init8()
	}

	err := s.addToSlice(value)
	if err != nil {
		return err
	}

	return nil
}

func (s *Slice) init8() { // функция присвоение значений полям из структуры
	s.Capacity = baseCapacity
	s.Array = [baseCapacity]*int{}
	s.Lenght = 0
}

func (s *Slice) init16() { // функция присвоение значений полям из структуры
	arr8 := s.Array.([baseCapacity]*int)
	s.Capacity = capacity16
	s.Lenght = baseCapacity
	newArr16 := [capacity16]*int{}

	for i := 0; i < baseCapacity; i++ { // присвоение значений из более раннего массива в новый
		newArr16[i] = arr8[i]
	}

	s.Array = newArr16
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

func (s *Slice) addToSlice(value int) error { // запись переданных значений в массив
	if s.Lenght >= capacity16 {
		return errors.New("size limit of slice")
	}

	if s.Lenght < baseCapacity {
		s.addInto8Array(value)
	} else if s.Lenght >= baseCapacity {
		s.addInto16Array(value)
	}

	return nil
}

func (s *Slice) addInto8Array(value int) {
	arr8 := s.Array.([baseCapacity]*int)
	if arr8[0] == nil {
		arr8[0] = &value
		s.Lenght = s.Lenght + 1
	} else {
		arr8[s.Lenght] = &value
		s.Lenght = s.Lenght + 1
	}
	s.Array = arr8 // запись в массив
}

func (s *Slice) addInto16Array(value int) {
	if s.Lenght == baseCapacity {
		s.init16()
	}

	arr16 := s.Array.([capacity16]*int)
	arr16[s.Lenght] = &value
	s.Lenght = s.Lenght + 1
	s.Array = arr16
}

func (s *Slice) PrintSlice() {
	if s.Lenght <= baseCapacity {
		for _, v := range s.Array.([baseCapacity]*int) {
			if v != nil {
				fmt.Printf("%v ", *v)
			}
		}
	} else {
		for _, v := range s.Array.([capacity16]*int) {
			if v != nil {
				fmt.Printf("%v ", *v)
			}
		}
	}
	fmt.Println()
}

func (s *Slice) DeleteSlice(index int) {
	if s.Lenght <= baseCapacity {
		s.deleteSlice8(index)
	} else if s.Lenght > baseCapacity {
		s.deleteSlice16(index)
	}
}
func (s *Slice) deleteSlice8(index int) {
	arr8 := s.Array.([baseCapacity]*int)
	arr8[index] = nil
	s.Array = arr8
	s.moveIndex8()
}

func (s *Slice) deleteSlice16(index int) {
	arr16 := s.Array.([capacity16]*int)
	arr16[index] = nil
	s.Array = arr16
	s.moveIndex16()
}

func (s *Slice) moveIndex8() {
	arr8 := s.Array.([baseCapacity]*int)
	for index, v := range arr8 {
		if v == nil && index < 7 {
			for i := index; i < 7; i++ {
				if i == 6 {
					arr8[i] = arr8[i+1]
					arr8[i+1] = nil
				} else {
					arr8[i] = arr8[i+1]
				}
			}
		}
	}
	s.Array = arr8
	s.initLenght()
}

func (s *Slice) moveIndex16() {
	arr16 := s.Array.([capacity16]*int)
	for index, v := range arr16 {
		if v == nil && index < 15 {
			for i := index; i < 15; i++ {
				if i == 14 {
					arr16[i] = arr16[i+1]
					arr16[i+1] = nil
				} else {
					arr16[i] = arr16[i+1]
				}
			}
		}
	}
	s.Array = arr16
	s.initLenght()
}

func (s *Slice) BubbleSort() {
	if s.Lenght <= baseCapacity {
		s.bubbleSort8()
	} else if s.Lenght > baseCapacity {
		s.BubbleSort16()
	}
}

func (s *Slice) bubbleSort8() {
	arr8 := s.Array.([baseCapacity]*int)
	newSlice := make([]int, s.Lenght)
	for i, v := range arr8 {
		if v != nil {
			newSlice[i] = *v
		}
	}
	sort.Ints(newSlice)
	for i := 0; i < s.Lenght; i++ {
		arr8[i] = &newSlice[i]
	}
	s.Array = arr8
}

func (s *Slice) BubbleSort16() {
	arr16 := s.Array.([capacity16]*int)
	newSlice := make([]int, s.Lenght)
	for i, v := range arr16 {
		if v != nil {
			newSlice[i] = *v
		}
	}
	sort.Ints(newSlice)
	for i := 0; i < s.Lenght; i++ {
		arr16[i] = &newSlice[i]
	}
	s.Array = arr16
}

func main() {
	// slice
	var slice1 Slice
	rand.Seed(time.Now().UnixNano())
	x := 11
	n := [11]int{5, 10, 13, 1, 4, 22, 15, 40, 99, 100, 16}
	for i := 0; i < x; i++ {
		y := n[i]
		err := slice1.AddElement(y)
		if err != nil {
			fmt.Println(err)
		}
	}

	slice1.PrintSlice()
	slice1.BubbleSort()
	slice1.PrintSlice()
	// 	fmt.Println(slice1)
	// 	slice1.DeleteSlice(1)
	// 	slice1.PrintSlice()
	// 	fmt.Println(slice1)
}

// функция Delete(index int) по индексу
// функция Update(index int) по индексу
// функция BubbleSort()
