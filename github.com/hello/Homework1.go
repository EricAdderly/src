package main

import (
	"fmt"
)

const baseCapacity int = 8
const capacity16 int = 16

type Slice struct {
	Array    interface{} // как то нужно сделать поле, которое будет создавать массив (скорее всего через функцию)
	Capacity int
	Lenght   int
}

func (s *Slice) AddElement(value int) {
	if s.Array == nil {
		s.init()
	}

	s.addToSlice(value)

}

func (s *Slice) init() {
	s.Capacity = baseCapacity
	s.Array = [baseCapacity]*int{} // теперь len() функция должна работать правильно. Вернуть правильное кол-во элементов. Проверь
	s.initLenght()
}

func (s *Slice) initLenght() { //[5, 6, nil, nil, nil, nil]
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
func (s *Slice) addToSlice(value int) {
	switch s.Capacity {
	case baseCapacity:
		arr8 := s.Array.([baseCapacity]*int) // приведение тип

		if s.Lenght >= baseCapacity {
			newArr16 := [capacity16]*int{} // переписать этот цикл через for range из inirLeght
			for i := 0; i < baseCapacity; i++ {
				newArr16[i] = arr8[i]
			}

			s.initLenght()
			fmt.Println(s.Lenght)
			newArr16[s.Lenght+1] = &value
		}

		fmt.Println("arr8: ", arr8)
		if s.Lenght == 0 {
			arr8[s.Lenght] = &value
		}

		arr8[s.Lenght+1] = &value
		s.initLenght()
		s.Array = arr8
	}

}

func main() {
	var slice1 Slice
	slice1.AddElement(1)
	slice1.AddElement(2)
	fmt.Println(slice1)
}
