package main

//import "fmt"

// import (
// 	"fmt"
// )

// Множественное присвоение

// a, b, c := 1, 2, 3

// a, _, c := 4, 5, 6

// if - else

// func main() {
// 	message, entered := enterTheClub(70) //message - переменная на сообщение, entered - переменная на булин (да/нет), в скобках передаём число
// 	fmt.Println(message) // печать сообщения
// 	fmt.Println(entered) // печать булина
// }

// func enterTheClub(age int) (string, bool) {  // вызываем функцию, которая принимает переменную возраст, число. Обратно отдаём строку, и буль
// 	if age >= 18 && age < 45 {  // условие 1 && - или или
// 		return "Входи", true  // возвращаем текст и буль на 1 условие
// 	} else if age >= 45 && age < 65 { // условие 2 && - или или
// 		return "Are you sure", true //возвращаем текст и буль на 2 условие
// 	} else if age >= 65 { 	// условие 3
// 		return "Wrong door", false //возвращаем текст и буль на 3 условие
// 	}
// 	return "Нет 18-ти", false // что происходит, если не выполнились условия выше
// }

// switch -case

// func main() {
// 	fmt.Println(prediction("Сб")) // вызываем печать , в скобках вызываем функцию, и передаём строку
// }

// func prediction(dayOfWeek string) (string, error) { //Объявляем функцию, принимаем строку (и объявляем её?), выводим строку и ошибку
// 	switch dayOfWeek { // Начало конструкции switch, и пишем что будет "свитчить"
// 	case "Пн": // условие 1
// 		return "Понедельник", nil //что возвращаем в условии 1
// 	case "Вт":
// 		return "Вторник", nil
// 	case "Ср":
// 		return "Среда", nil
// 	case "Чт":
// 		return "Четверг", nil
// 	case "Пт":
// 		return "Пятница", nil
// 	default: // что происходит если не попали в кейсы выше
// 		return "Это не будни", errors.New("It's weekend") //что возвращаем
// 	}
// }

// Неограниченное количество аргументов

// func main() {
// 	fmt.Println(findMin(1, 250, -15, 40, 13)) // вызываем печать , в скобках вызываем функцию, и передаём числа
// }

// func findMin(numbers ...int) int { // объявляем функцию, в которую передаём передаём/создаём переменную numbers, ... означает, что можем передать неограниченное кол-во целых чисел, возвращаем число
// 	if len(numbers) == 0 { // условия, в котором считаем кол-во элементов внутри массива numbers, и берём первое из него (0)
// 		return 0 // возвращаем 0
// 	}

// 	min := numbers[0] //задаём переменную, которая равняется первому значению из массива

// 	for _, i := range numbers { // непонятно почему for _, - безымянная переменная (возможно таким способом мы задаём значение, так как не можем вписать его сами), range - перебираем значения из массива?
// 		if i < min { // услвоия, если переменная i < меньше минимальной из значения из массива
// 			min = i // присваем i это значене
// 		}
// 	}

// 	return min // возвращаем самое переменную min
// }

// анонимные функции

// func main() {
// 	func() { //задаём анонимную функцию (без названия и значения)
// 		fmt.Println("Анонимная функция") // вызываем печать определённый тест
// 	}() //вызываем анонимную функцию
// }

// замыкания

// func main() {
// 	inc := increment() // присваиваем inc - функцию, так как increment возвращает функцию, то inc теперь переменная для вызова функции
// 	fmt.Println(inc()) // вызываем функцию increment, так как мы каждый раз возвращаем значение count в main, то значение count при каждом вызове будет увеличиваться
// 	fmt.Println(inc())
// 	fmt.Println(inc())
// 	fmt.Println(inc())
// 	fmt.Println(increment2()) // каждый раз возвращается 1, так как нет анонимной функции, которая запоминает своё состояние прошлого раза
// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// }

// func increment() func() int { //задаём функцию increment, которая возвращает анонимную функцию, которая возвращает целое число
// 	count := 0          // задаём переменную count, и даём ей значени 0
// 	return func() int { // возвращаем анонимную функцию, вызывая её
// 		count++      // увеличиваем count на 1
// 		return count // возвращаем count
// 	}
// }

// func increment2() int {
// 	count := 0
// 	count++
// 	return count
// }

// init

// var msg string

// func init() { // функция init срабатывает при запуске пакета (до main)
// 	msg = "Какой то текст"
// }
// func main() {
// 	fmt.Println(msg)
// }

// Указатели

// func main() {
// 	message := "Какой то текст"
// 	printMessage(message) // если вызывать ещё раз message, не передавая её в в сл. функцию, то у нас снова выдаст значение из main, так как message не был перезаписан
// }

// func printMessage(message string) {
// 	message += ", и ещё какой то текст" // += означает что мы добавляем ещё текст, к переменной, в которой уже был текст
// 	fmt.Println(message)
// }

// func main() {
// 	message := "Какой то текст"
// 	printMessage(&message) // передаём ссылку (область памяти) на переменную message
// 	fmt.Println(message)
// }

// func printMessage(message *string) { //задаём функцию, в нём получаем переменную message, и указатель на строку
// 	*message += ", и ещё какой то текст" // к области памяти мы прибавляем ещё текст+= означает что мы добавляем ещё текст, к переменной, в которой уже был текст
// }

// массивы и слайсы
// массив
// func main() {
// 	messages := [3]string{"Яблоко", "Дерево", "Дорога"} //в массиве [] - кол-во элементов, string - тип данных в массиве, {} - значения
// }

//слайс

// func main() {
// 	messages := []string{"1", "2", "3"} // задаём слайс, разница в том, что в слайсе не указано изначально кол-во элементов
// 	printMessage(messages)              //передаём в функцию слайс

// 	fmt.Println(messages) // Вызываем запись слайса, так как слайс перезаписывается, то получаем уже изменённые данные
// }

// func printMessage(messages []string) error { // объявляем функцию, которая принимает messages, с типом слайс, возвращаем ошибку
// 	if len(messages) == 0 { // условие - если слайс пустой
// 		return errors.New("empty array") // то возвращаем ошибку
// 	}

// 	messages[1] = "5" // присваиваем второму элементу значение "5"

// 	fmt.Println(messages) // печатаем messages

// 	return nil // возращаем остутствие ошибки
// }

// матрицы, цикл for

// func main() {
// 	matrix := make([][]int, 10) // инициализируем слайс-матрицу 10х10

// 	for x := 0; x < 10; x++ { // х стартует с 0, пока не станет < 10, шаг в +1, запускается сл цикл for
// 		for y := 0; y < 10; y++ { // y стартует с 0, пока не станет < 10, шаг в +1, y отвечает за движение по горизонтали
// 			matrix[y] = make([]int, 10)
// 			matrix[y][x] = x // y движется по горизонтали, подставляя значение отличное от 0, когда он совпадает с x
// 		}
// 		fmt.Println(matrix[x])
// 	}
// }

// defer
// func main() {
// 	defer printMessage() // defer показывает последний шаг для исполнения и выходит из него

// 	fmt.Println("main()")
// 	fmt.Println("main() 2")
// 	fmt.Println("main() 3")
// }

// func printMessage() {
// 	fmt.Println("printMessage()")
// }

// panic

// func main() {
// 	defer handlerPanic()  // defer показывает последний шаг для исполнения и выходит из него
// 	messages := []string{ //задаём слайс
// 		"String 1",
// 		"string 2",
// 		"string 3",
// 		"string 4",
// 	}
// 	messages[4] = "string 5" // пытаеимся поменять значение в слайсе, но так нельзя
// 	fmt.Println(messages)    // печатаем messages
// }

// func handlerPanic() { // функция, которая будет обрабатывать панику
// 	if r := recover(); r != nil { //если r = функции recover (которая не останавливает наше приложение), в которой r не равно отсутствии ошибки
// 		fmt.Println(r) // вызываем r , которая = функции recover
// 	}
// 	fmt.Println("handlerPanic Выполнилась успешно") // пишем сообщение
// }

// мапы

// func main() {
// 	users := map[string]int{ // объявляем маппу - мапа это неупорядочный массив с ключ - значение
// 		"Vasya":  15, // записываем ключ - значение, ключ считается уникальным значением, и не может повторяться внутри одной мапы
// 		"Petya":  23,
// 		"Kostya": 48,
// 	}
// 	age, exists := users["Kostya"] // задаём переменную age и exist, которая проверяет есть ли человек в мапе
// 	if exists {
// 		fmt.Println("Kostya", age)
// 	} else {
// 		fmt.Println("Кости нет в списке")
// 	}
// }

// func main() {
// 	users := map[string]int{ // объявляем маппу - мапа это неупорядочный массив с ключ - значение
// 		"Vasya":  15, // записываем ключ - значение, ключ считается уникальным значением, и не может повторяться внутри одной мапы
// 		"Petya":  23,
// 		"Kostya": 48,
// 	}

// 	var users1 map[string]int     // задаём мапу, но не инициализируем её, пока мы ее не синицилиазируем, мы не сможем в неё что то записывать
// 	users1 = make(map[string]int) // инициализируем мапу
// 	users1["Igor"] = 12           // добавляем значение в мапу
// 	fmt.Println(users1)           // печатаем

// 	users["Serega"] = 21   // добавляем ещё одно ключ - значение в мапу
// 	delete(users, "Vasya") // удаляем Васю из мапы

// 	for key, value := range users { // берём ключ значение из мапы users
// 		fmt.Println(key, value) // и печатаем их
// 	}
// }

// структуры

// type Age int // создаём тип Age, смысл в том, что для типа можем делать методы

// func (a Age) isAdult() bool { // создаём метод для Age, который будет возвращать бул, если a >= 18
// 	return a >= 18
// }

// type User struct { // задаём структуру, стуктуру нельзя менять после инициализации, но можно наполнять
// 	name   string
// 	age    Age // вместо типа int, указываем Age
// 	sex    string
// 	weight int
// 	height int
// }

// func (u *User) SetName(name string) { // делаем ресивер для изменения данных, *user, означает, что будем перезаписывать данные
// 	u.name = name
// }

// func (u User) getName() string { // делаем ресивер для запроса данных, user, означает, что будем только их получать
// 	return u.name
// }

// func NewUser(name string, age int, sex string, weight int, height int) User { // инициилизируем конструктор - функция, которая инициизирует объект опр. типа
// 	return User{
// 		name:   name,     // перечисляем из чего состоит
// 		age:    Age(age), // делаем приведение типа
// 		sex:    sex,
// 		weight: weight,
// 		height: height,
// 	}
// }
// func main() {
// 	// user := struct { // задаём структуру, стуктуру нельзя менять после инициализации, но можно наполнять (второй вариант)
// 	// 	name   string
// 	// 	age    int
// 	// 	sex    string
// 	// 	weight int
// 	// 	height int
// 	// }{"Vasya", 25, "male", 70, 185}
// 	user1 := NewUser("Vasya", 25, "male", 70, 185) //формат через конструктор
// 	user2 := User{"Igor", 12, "male", 50, 160}     //ручной формат

// 	fmt.Println(user1.age)
// 	fmt.Println(user1, user2)  // можно ли вынести всю структуру
// 	fmt.Printf("%+v\n", user1) // можно ли как то вынести сразу 2х юзеров? и более
// 	fmt.Printf("%+v\n", user2)

// 	fmt.Println(user1.getName()) // запрашиваем данные имени из getName
// 	fmt.Println(user1.name)      // убеждаемся, что ничего не имзенилось
// 	user1.SetName("Ivan")        // Меняем имя user1
// 	fmt.Println(user1.name)      // убеждаемся, что изменилось

// 	fmt.Println(user2.age.isAdult()) // взываем проверку на есть ли 18

// }

// func main() {
// 	x := [4]int{1, 2, 3, 4}
// 	y := [8]int{5, 6, 7, 8, 9, 10}
// 	for i := 0; i < 4; i++ {
// 		y[i] = x[i]
// 	}
// 	fmt.Println(y)
// 	i := len(y)
// 	fmt.Println(i)
// }

// Бинарный поиск

// func main() {
// 	nums := []int{-1, 0, 3, 5, 9, 12}
// 	target := 12
// 	fmt.Println(search(nums, target))
// }

// func search(nums []int, target int) int {
// 	var array []int
// 	array = nums
// 	localTarget := target
// 	low := 0
// 	high := len(array) - 1

// 	for low <= high {
// 		median := (low + high) / 2

// 		if array[median] < localTarget {
// 			low = median + 1
// 		} else if array[median] > localTarget {
// 			high = median - 1
// 		} else {
// 			return median
// 		}
// 	}
// 	return -1
// }

// алгоритм поиска значения
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	nums := []int{-1, 0, 3, 5, 9, 12}
// 	target := 12
// 	fmt.Println(search(nums, target))
// }

// func search(nums []int, target int) int {
// 	var array []int
// 	array = nums
// 	localTarget := target
// 	low := 0
// 	high := len(array) - 1

// 	for low <= high {
// 		median := (low + high) / 2

// 		if array[median] < localTarget {
// 			low = median + 1
// 		} else if array[median] > localTarget {
// 			high = median - 1
// 		} else {
// 			return median
// 		}
// 	}
// 	return -1
// }
