package currency

var BRL = CurrencyImpl{
	code:   "BRL",
	symbol: "R$",
	format: func(currency Currency, amount int) string {
		prefix := currency.Symbol() + " "
		return prefix + addSeparetor(padLeft(amount, "0", 3), ",", 2)
	},
}
