package proptests

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{500, "D"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman),
			func(t *testing.T) {
				got := ConvertToRoman(test.Arabic)
				if got != test.Roman {
					t.Errorf("got %q want %q", got, test.Roman)
				}
			})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic),
			func(t *testing.T) {
				got := ConvertToArabic(test.Roman)
				if got != test.Arabic {
					t.Errorf("got %d want %d", got, test.Arabic)
				}
			})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic

	}
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
