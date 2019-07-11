package numerals

import "strings"

/*ConvertToRoman converts a given int to roman numeral*/
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
