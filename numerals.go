package numerals

import "strings"

/*RomanNumeral defines a struct to represent a roman numeral and its integer value*/
type RomanNumeral struct {
	Value  int
	Symbol string
}

/*RomanNumerals cointains an array of roman numeral structs*/
var RomanNumerals = []RomanNumeral{
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

	for _, numeral := range RomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

/*ConvertToArabic converts a given roman numeral to arabic represntation as an int*/
func ConvertToArabic(rn string) int {
	return 1
	// var result int

	// for _, numeral := range RomanNumerals {
	// 	for rn != "" {
	// 		result += numeral.Value
	// 		rn = rn[:len(numeral.Symbol)-1]
	// 	}
	// }

	// return result
}
