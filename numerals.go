package numerals

import (
	"strings"
)

/*RomanNumeral defines a struct to represent a roman numeral and its integer value*/
type RomanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []RomanNumeral

/*RomanNumerals creates an array of roman numeral structs*/
func (r romanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

/*ConvertToRoman converts a given int to roman numeral*/
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

/*ConvertToArabic converts a given roman numeral to arabic represntation as an int*/
func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			// build the two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			// get the value of the two character string
			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total++
			}
		} else {
			total++
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}
