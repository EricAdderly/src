package main

// import (
// 	"testing"

// 	"github.com/wallet_stats/wallet"
// )

// // тест функции создания кошелька
// func TestCreateWallet(t *testing.T) {
// 	// тестовые данные
// 	testTable := []struct {
// 		name string
// 	}{
// 		{
// 			name: "Ivanov Ivan Ivanovich",
// 		},
// 		{
// 			name: "",
// 		},
// 	}

// 	// Тест
// 	_, errCreateWallet := wallet.CreateWallet(testTable[0].name)
// 	if errCreateWallet != nil {
// 		t.Error("Incorrect result. | Correct_name | Expected: nil | got error")
// 	}
// 	_, errCreateWallet2 := wallet.CreateWallet(testTable[1].name)
// 	if errCreateWallet2 == nil {
// 		t.Error("Incorrect result. | Empty_name | Expected: error | got: nil")
// 	}
// }

// // тест функции получения кошелька
// func TestGetWallet(t *testing.T) {
// 	// тестовые данные
// 	id1, _ := wallet.CreateWallet("Ivanov Ivan Ivanovich")
// 	id2, _ := wallet.CreateWallet("Teodor Adderly Romanov")
// 	wallet.DeleteWallet(id2)
// 	id3 := ""
// 	id4 := "4"

// 	// Тест
// 	_, err1 := wallet.GetWalletById(id1)
// 	if err1 != nil {
// 		t.Error("Incorrect result. | Correct_id | Expected: nil | got error")
// 	}

// 	_, err2 := wallet.GetWalletById(id2)
// 	if err2 == nil {
// 		t.Error("Incorrect result. | Inactive_wallet | Expected: error | got: nil")
// 	}

// 	_, err3 := wallet.GetWalletById(id3)
// 	if err3 == nil {
// 		t.Error("Incorrect result. | Empty_id | Expected: error | got: nil")
// 	}

// 	_, err4 := wallet.GetWalletById(id4)
// 	if err4 == nil {
// 		t.Error("Incorrect result. | Non_existent_id | Expected: error | got: nil")
// 	}
// }

// // тест функции обновления кошелька
// func TestUpdateWallets(t *testing.T) {
// 	// тестовые данные
// 	// 1 кейс
// 	name1 := "Кузнецов Фёдор Алиевич"
// 	id1, _ := wallet.CreateWallet(name1)
// 	newName := "Смирнов Артём Семёнович"
// 	// 2 кейс
// 	name2 := "Teodor Adderly Romanov"
// 	id2, _ := wallet.CreateWallet(name2)
// 	wallet.DeleteWallet(id2)
// 	// остальные
// 	id3 := ""
// 	name3 := ""
// 	id4 := "4"
// 	name4 := "Никитин Леонид Леонидович"

// 	// тест
// 	// 1 кейс
// 	err1 := wallet.UpdateWallet(id1, newName)
// 	if err1 != nil {
// 		t.Error("Incorrect result. | Correct_name | Correct_id | Expected: nil | got error")
// 	}
// 	wallet1, _ := wallet.GetWalletById(id1)
// 	if wallet1.Name != newName {
// 		t.Error("Incorrect result. | Correct_name | Correct_id | Name did not changed")
// 	}

// 	// 2 кейс
// 	err2 := wallet.UpdateWallet(id2, newName)
// 	if err2 == nil {
// 		t.Error("Incorrect result. | Inactive_wallet | Expected: error | got nil")
// 	}

// 	// 3 кейс
// 	err3 := wallet.UpdateWallet(id3, newName)
// 	if err3 == nil {
// 		t.Error("Incorrect result. | Empty_id | Correct_name | Expected: error | got nil")
// 	}

// 	// 4 кейс
// 	err4 := wallet.UpdateWallet(id1, name3)
// 	if err4 == nil {
// 		t.Error("Incorrect result. | Correct_id | Empty_name | Expected: error | got nil")
// 	}

// 	// 5 кейс
// 	err5 := wallet.UpdateWallet(id4, name4)
// 	if err5 == nil {
// 		t.Error("Incorrect result. | Inactive_wallet | Empty_name | Expected: error | got nil")
// 	}

// 	// 6 кейс
// 	err6 := wallet.UpdateWallet(id3, name4)
// 	if err6 == nil {
// 		t.Error("Incorrect result. | Empty_id | Empty_name | Expected: error | got nil")
// 	}

// 	// 7 кейс
// 	err7 := wallet.UpdateWallet(id4, name4)
// 	if err7 == nil {
// 		t.Error("Incorrect result. | Non_existent_id | Correct_name | Expected: error | got nil")
// 	}
// }

// // тест функции удаления кошелька
// func TestDeleteWallets(t *testing.T) {
// 	// тестовые данные
// 	name1 := "Ivanov Ivan Ivanovich"
// 	id1, _ := wallet.CreateWallet(name1)
// 	id2 := ""
// 	id3 := "4"

// 	// тест
// 	err1 := wallet.DeleteWallet(id1)
// 	if err1 != nil {
// 		t.Error("Incorrect result. | Correct_id | Expected: nil | got error")
// 	}

// 	err2 := wallet.DeleteWallet(id1)
// 	if err2 == nil {
// 		t.Error("Incorrect result. | Inactive_wallet | Expected: error | got nil")
// 	}

// 	err3 := wallet.DeleteWallet(id2)
// 	if err3 == nil {
// 		t.Error("Incorrect result. | Empty_id | Expected: error | got nil")
// 	}

// 	err4 := wallet.DeleteWallet(id3)
// 	if err4 == nil {
// 		t.Error("Incorrect result. | Non_existent_id | Expected: error | got nil")
// 	}
// }

// // текст функции пополнения кошелька
// func TestDepositBallance(t *testing.T) {
// 	//тестовые данные
// 	name1 := "Ivanov Ivan Ivanovich"
// 	name2 := "Teodor Adderly Romanov"
// 	id1, _ := wallet.CreateWallet(name1)
// 	id2, _ := wallet.CreateWallet(name2)
// 	wallet.DeleteWallet(id2)
// 	id3 := ""
// 	id4 := "4"
// 	summ1 := 500.0
// 	summ2 := -500.0
// 	summ3 := 0.0

// 	// тесты
// 	err1 := wallet.DepositBallance(id1, summ1)
// 	if err1 != nil {
// 		t.Error("Incorrect result. | Correct_id | Correct_summ | Expected: nil | got error")
// 	}

// 	err2 := wallet.DepositBallance(id1, summ2)
// 	if err2 == nil {
// 		t.Error("Incorrect result. | Correct_id | Negative amount | Expected: error | got nil")
// 	}

// 	err3 := wallet.DepositBallance(id1, summ3)
// 	if err3 == nil {
// 		t.Error("Incorrect result. | Correct_id | Amount = 0 | Expected: error | got nil")
// 	}

// 	err4 := wallet.DepositBallance(id2, summ1)
// 	if err4 == nil {
// 		t.Error("Incorrect result. | Inactive_wallet | Correct_summ | Expected: error | got nil")
// 	}

// 	err5 := wallet.DepositBallance(id3, summ1)
// 	if err5 == nil {
// 		t.Error("Incorrect result. | Empty_id | Correct_summ | Expected: error | got nil")
// 	}

// 	err6 := wallet.DepositBallance(id4, summ1)
// 	if err6 == nil {
// 		t.Error("Incorrect result. | Non_existent_id | Correct_summ | Expected: error | got nil")
// 	}
// }
