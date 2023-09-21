package models

type RequiredEventFields struct {
	EventName        string // обязательное поле
	EventTheme       string // наверное тоже обязательное
	EventLocation    EventLocation // тоже обязательное
	EventData        string // обязательное 
	EventPublic      bool // обязательное поле 
}
