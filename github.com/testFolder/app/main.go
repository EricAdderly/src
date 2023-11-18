package main

import (
	"net/http"

	"github.com/testFolder/app/internal"
)

// type Data struct {
// 	sync.RWMutex
// 	Text string
// }

func main() {
	internal.Tpl, _ = internal.Tpl.ParseGlob("app/templates/*.html")
	http.HandleFunc("/town", internal.AnswerHandleFunc)
	http.HandleFunc("/choosetown", internal.ChoseTownHandleFunc)
	http.HandleFunc("/weather", internal.GetWetherHandleFunc)
	http.ListenAndServe(":8080", nil)

	// 	Data := CreateData()

	// 	var wg sync.WaitGroup

	// 	for i := 0; i < 3; i++ {
	// 		wg.Add(1) // добавляем в общую группу
	// 		go reader(Data, i, &wg)
	// 	}

	// 	for i := 0; i < 2; i++ {
	// 		wg.Add(1) // добавляем в общую группу
	// 		go wtiter(Data, i, &wg)
	// 	}

	// 	wg.Wait() // нужно чтобы дождаться завершения всех горутин
	// }

	// func CreateData() *Data {
	// 	return &Data{}
	// }

	// func reader(Data *Data, id int, wg *sync.WaitGroup) {
	// 	defer wg.Done()

	// 	for i := 0; i < 5; i++ {
	// 		Data.read(id)
	// 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	// 	}

	// }

	// func (d *Data) read(id int) {
	// 	d.RLock() // позволяет читать в несколько потоков
	// 	defer d.RUnlock()

	// 	fmt.Println("Читатель читает текст: %s\n", id, d.Text)
	// 	time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	// 	fmt.Printf("Читатель %d закончил читать\n", id)
	// }

	// func wtiter(Data *Data, id int, wg *sync.WaitGroup) {
	// 	defer wg.Done()

	// 	for i := 0; i < 3; i++ {
	// 		Data.wtite(id)
	// 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	// 	}

	// }

	// func (d *Data) wtite(id int) {
	// 	d.Lock() // не позволяет ни читать ни писать, пока не разблочим
	// 	defer d.Unlock()

	// 	newString := randString(10)
	// 	fmt.Printf("Писатель %d пишет текст: %s\n", id, newString)
	// 	d.Text = newString
	// 	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	// 	fmt.Printf("Писатель %d закончил писать\n", id)

	// }

	// func randString(length int) string {
	// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 	result := make([]byte, length)
	// 	for i := range result {
	// 		result[i] = charset[rand.Intn(len(charset))]
	// 	}
	// 	return string(result)
}
