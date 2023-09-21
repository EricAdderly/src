package usecase

import (
	"context"
	models "e-project/model"
)

type ReceivingTransactionUseCase struct {
	repo ReceivingTransactionRepo
}

func NewReceiveTransaction(r ReceivingTransactionRepo) *ReceivingTransactionUseCase {
	return &ReceivingTransactionUseCase{
		repo: r,
	}
}

func (uc *ReceivingTransactionUseCase) ReceivingTransaction(ctx context.Context, request models.ReceivingTransactionRequest) error {
	// отправляем запрос в БД на существование сессии
	//проверяем
	//если всё ок, отправляет Nil
	return nil
}
