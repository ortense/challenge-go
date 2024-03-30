package card

import (
	"time"
)

type PaymentCard struct {
	Number     string
	Holder     string
	Expiration time.Time
	CVV        string
}

func New(number, holder, expiration, cvv string) PaymentCard {
	return PaymentCard{
		Number:     number[0:4],
		Holder:     holder,
		Expiration: getExpirationDate(expiration),
		CVV:        cvv,
	}
}

func (pc PaymentCard) IsExpired() bool {
	return time.Now().After(pc.Expiration)
}

func getExpirationDate(expiration string) time.Time {
	date, err := time.Parse("02/01/06", "01/"+expiration)

	if err != nil {
		return time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	date = date.AddDate(0, 1, 0)
	date = date.Add(-time.Minute)
	return date
}
