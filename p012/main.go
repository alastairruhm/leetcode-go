package p012

import (
	"math"
	"strings"
)

func intToRoman(num int) string {
	m := genMap()
	var digits []int
	radix := 0
	for num > 0 {
		d := num % 10
		num = num / 10
		digits = append(digits, d*int(math.Pow10(radix)))
		radix++
		// digits = append(digits, d)
	}
	// fmt.Printf("%+v\n", digits)
	result := ""
	for i := 0; i < len(digits); i++ {
		result = m[digits[i]] + result
	}
	return result
}

func genMap() map[int]string {
	m := make(map[int]string)
	i := 0
	m[0] = ""
	for i = 1; i < 10; i++ {
		if i < 5 {
			m[i] = strings.Repeat("I", i)
		} else {
			m[i] = "V" + strings.Repeat("I", i-5)
		}
	}
	m[4] = "IV"
	m[9] = "IX"
	for i = 10; i < 100; i = i + 10 {
		if i < 50 {
			m[i] = strings.Repeat("X", i/10)
		} else {
			m[i] = "L" + strings.Repeat("X", i/10-5)
		}
	}
	m[40] = "XL"
	m[90] = "XC"
	for i = 100; i < 1000; i = i + 100 {
		if i < 500 {
			m[i] = strings.Repeat("C", i/100)
		} else {
			m[i] = "D" + strings.Repeat("C", i/100-5)
		}
	}
	m[400] = "CD"
	m[900] = "CM"
	for i = 1000; i < 4000; i = i + 1000 {
		m[i] = strings.Repeat("M", i/1000)
	}
	// m[1] = 'I'
	// m[4] = 'IV'
	// m[5] = 'V'
	// m[9] = 'IX'
	// m[10] = 'X'
	// m[40] = 'XL'
	// m[50] = 'L'
	// m[90] = 'XC'
	// m[100] = 'C'
	// m[400] = 'CD'
	// m[500] = 'D'
	// m[900] = 'CM'
	// m[1000] = 'M'

	return m
}
