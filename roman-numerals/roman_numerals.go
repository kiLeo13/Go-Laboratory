package numerals

import "strings"

func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {

    var arabic uint16

    for _, numeral := range romanNumerals {
        for strings.HasPrefix(roman, numeral.Symbol) {
            arabic += numeral.Value
            roman = strings.TrimPrefix(roman, numeral.Symbol)
        }
    }

    return arabic
}

type RomanNumeral struct {
    Value  uint16
    Symbol string
}

var romanNumerals = []RomanNumeral{
    {1000, "M"},
    {900,  "CM"},
    {500,  "D"},
    {400,  "CD"},
    {100,  "C"},
    {90,   "XC"},
    {50,   "L"},
    {40,   "XL"},
    {10,   "X"},
    {9,    "IX"},
    {5,    "V"},
    {4,    "IV"},
    {1,    "I"},
}