package models

type Currency struct {
	Persistable
	Name string `json:"name" gorm:"unique;not null"`
}

type Amount struct {
	Persistable
	Value      float64  `json:"value" gorm:"type:decimal;default:0.0"`
	CurrencyID string   `json:"-" gorm:"not null"`
	Currency   Currency `json:"currency"`
}

func GetCurrency(name string) (*Currency, error) {
	currency := &Currency{}
	result := db.First(&currency, "currency = ?", name)
	return currency, HandleError(result.Error)
}
