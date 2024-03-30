package currency

import "errors"

var mapCurrency = map[string]Currency{
	BRL.code: &BRL,
	USD.code: &USD,
}

func FromCode(code string) (Currency, error) {
	c, ok := mapCurrency[code]

	if !ok {
		return nil, errors.New("unknown currency")
	}

	return c, nil
}
