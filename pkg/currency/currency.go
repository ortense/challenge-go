package currency

import (
	"strconv"
	"strings"
)

type Currency interface {
	Format(amount int) string
	Code() string
	Symbol() string
}

type CurrencyImpl struct {
	code   string
	symbol string
	format func(currency Currency, amount int) string
}

func (c *CurrencyImpl) Code() string {
	return c.code
}

func (c *CurrencyImpl) Symbol() string {
	return c.symbol
}

func (c *CurrencyImpl) Format(amount int) string {
	return c.format(c, amount)
}

func padLeft(value int, filler string, minlen int) string {
	str := strconv.Itoa(value)

	if len(str) >= minlen {
		return str
	}

	str = strings.Repeat(filler, minlen-len(str)) + str
	return str
}

func addSeparetor(value, token string, position int) string {
	if len(value) <= position {
		return value
	}

	return value[:len(value)-position] + token + value[len(value)-position:]
}
