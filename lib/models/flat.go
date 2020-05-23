package models

//Телефон (+7 по базе должно высвечиваться автоматически)
//Район
//Адрес
//дом №
//квартира №
//Цена объекта  от ………………….до …………..
//Кол -во комнат
//тип комнат
//Этаж / этажность
//Серия
//площадь М2
//Сан узел (Раздельный ,Во дворе,  Нет ,   Совмещенный)
//Кол -во сан узлов
//Год постройки
//Угловая или нет как у тебя стоит с галочкой так и оставим
//Описание

type Area string

const (
	Abayskiy      Area = "Абайский"
	AlFarabiskiy  Area = "Аль-Фарабийский"
	Enbekshinskiy Area = "Енбекшинский"
	Karatauskiy   Area = "Каратауский"
)

type state string

const (
	EuroState     state = "евро"
	VeryGoodState state = "класс"
	GoodState     state = "хороший"
	MiddleState   state = "средний"
	CosmeticState state = "косметический"
	BlackState    state = "черновой"
	NoneState     state = "без ремонта"
)

type toilet string

const (
	SeparatedToilet toilet = "Раздельный"
	OutSideToilet   toilet = "Во дворе"
	NoneToilet      toilet = "нет"
	JoinedToilet    toilet = "Совмещенный"
)

type Flat struct {
	ID           int64
	Area         Area
	LandMark     string
	Address      string
	HomeNumber   int
	FlatNumber   int
	PriceMin     int
	PriceMax     int
	RoomsCount   int
	RoomsType    string //TODO: addType
	Floor        int
	FloorsCount  int
	Square       int
	FlatType     string //TODO: addType
	State        state
	Toilet       toilet
	ToiletCount  int
	BuildYear    int
	IsCorner     bool
	Description  string
	PictureURLs  []string
	PhoneNumbers map[string]string
}

func NewFlat(Area Area,
	LandMark string,
	Address string,
	HomeNumber int,
	FlatNumber int,
	PriceMin int,
	PriceMax int,
	RoomsCount int,
	RoomsType string,
	Floor int,
	FloorsCount int,
	Square int,
	FlatType string,
	State state,
	Toilet toilet,
	ToiletCount int,
	BuildYear int,
	IsCorner bool,
	Description string,
	PictureURLs []string,
	PhoneNumbers map[string]string) Flat {

	return Flat{
		Area:         Area,
		LandMark:     LandMark,
		Address:      Address,
		HomeNumber:   HomeNumber,
		FlatNumber:   FlatNumber,
		PriceMin:     PriceMin,
		PriceMax:     PriceMax,
		RoomsCount:   RoomsCount,
		RoomsType:    RoomsType,
		Floor:        Floor,
		FloorsCount:  FloorsCount,
		Square:       Square,
		FlatType:     FlatType,
		State:        State,
		Toilet:       Toilet,
		ToiletCount:  ToiletCount,
		BuildYear:    BuildYear,
		IsCorner:     IsCorner,
		Description:  Description,
		PictureURLs:  PictureURLs,
		PhoneNumbers: PhoneNumbers,
	}
}
