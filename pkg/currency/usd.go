package currency

var USD = CurrencyImpl{
	code:   "USD",
	symbol: "$",
	format: func(currency Currency, amount int) string {
		return currency.Symbol() + addSeparetor(padLeft(amount, "0", 3), ".", 2)
	},
}
