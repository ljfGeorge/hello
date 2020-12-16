package hello

import "rsc.io/quote/v3"

func Hello() string {
	//testingasdasdasdsadasdsadsad
	return "Hello, world."
}

func Proverb() string {
	return quote.Concurrency()
}
