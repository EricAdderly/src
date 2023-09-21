package models

type Event struct {
	EventName        string              // обязательное поле
	EventTheme       string              // наверное тоже обязательное
	EventLocation    EventLocation       // тоже обязательное
	EventData        string              // обязательное
	EventBanner      map[int]EventBanner // если исходить из смысла что будет несколько фоток, то наверное удобно в хэшмапу засунуть всё. Необязательное
	EventDesctiption []string            // необязательное
	EventPublic      bool                // обязательное поле
	EventMembers     uint16              // необязательное
	EventChat        EventChat           //?
	EventAdmins      EventAdmins         // обязательное поле???
}

type EventLocation struct {
	Latitude  float32
	Longitude float32
}

type EventBanner struct {
	BannerId   int // если есть хэшмапа, то нужен ли BannerID
	BannerName string
	BannerFile []byte // не уверен что в таком формате должен быть файл
}

type EventChat struct { // непонятно как должно быть устроено

}

type EventAdmins struct {
	Admins map[int]User // не уверен что ключ должен быть интом
}

type User struct {
	FirstName    string
	SecondName   string
	BirthDate    string //или time.Time?
	Avatar       Avatar
	UserLocation UserLocation
}

type Avatar struct {
	AvatarId   int // если есть хэшмапа, то нужен ли BannerID
	AvatarName string
	AvatarFile []byte // не уверен что в таком формате должен быть файл
}

type UserLocation struct {
	Latitude  float32
	Longitude float32
}
