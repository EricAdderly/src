// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	models "e-project/model"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	// Translation interface {
	// 	Translate(context.Context, entity.Translation) (entity.Translation, error)
	// 	History(context.Context) ([]entity.Translation, error)
	// }

	// TranslationRepo -.
	// Registration interface {
	// 	SaveRegistrationData(context.Context, models.RegistrationRequest) error
	// }

	// TranslationWebAPI -.
	// TranslationWebAPI interface {
	// 	Translate(entity.Translation) (entity.Translation, error)
	// }

	//Registration -.
	// RegistrationRepo interface {
	// 	SaveUserData(context.Context, models.User) error
	// }

	//AUTH -.
	ReceivingTransaction interface {
		ReceivingTransaction(context.Context, models.ReceivingTransactionRequest) error
	}

	ReceivingTransactionRepo interface {
		GetSession(context.Context, string) error
	}

	BalanceChecker interface {
		BalanceChecker(context.Context, models.ReceivingTransactionRequest) error
	}

	BalanceCheckerRepo interface {
		GetSession(context.Context, string) (*models.TransactionSession, error)
		GetWallet(context.Context, string) (*models.Wallet, error)
		IncreaseBalance(context.Context, string, float32) error
		ReduceBalance(context.Context, string, float32) error
		ChangeStatusTransactionSession(context.Context, string) error
	}
)
