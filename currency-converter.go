package main

import (
	"math"
	"strings"
	"unicode"
)

var unitMap = map[string]string{
	"0": "zero",
	"1": "jeden",
	"2": "dwa",
	"3": "trzy",
	"4": "cztery",
	"5": "pięć",
	"6": "sześć",
	"7": "siedem",
	"8": "osiem",
	"9": "dziewięć",
}

var tenthsMap = map[string]string{
	"10": "dziesięć",
	"11": "jedenaście",
	"12": "dwanaście",
	"13": "trzynaście",
	"14": "czternaście",
	"15": "piętnaście",
	"16": "szesnaście",
	"17": "siedemnaście",
	"18": "osiemnaście",
	"19": "dziewiętnaście",
}

var upperTenthsMap = map[string]string{
	"2": "dwadzieścia",
	"3": "trzydzieści",
	"4": "czterdzieści",
	"5": "pięćdziesiąt",
	"6": "sześćdziesiąt",
	"7": "siedemdziesiąt",
	"8": "osiemdziesiąt",
	"9": "dziewięćdziesiąt",
}

var hundredsMap = map[string]string{
	"1": "sto",
	"2": "dwieście",
	"3": "trzysta",
	"4": "czterysta",
	"5": "pięćset",
	"6": "sześćset",
	"7": "siedemset",
	"8": "osiemset",
	"9": "dziewięćset",
}

var singular = []string{" ", "tysiąc ", "milion ", "miliard ", "bilion ", "biliard ", "trylion ", "tryliard ", "kwadrylion ", "kwadryliard ", "kwintylion ", "kwintyliard "}
var pluralSimple = []string{" ", "tysiące ", "miliony ", "miliardy ", "biliony ", "biliardy ", "tryliony ", "tryliardy ", "kwadryliony ", "kwadryliardy ", "kwintyliony ", "kwintyliardy "}
var pluralUpper = []string{" ", "tysięcy ", "milionów ", "miliardów ", "bilionów ", "biliardów ", "trylionów ", "tryliardów ", "kwadrylionów ", "kwadryliardów ", "kwintylionów ", "kwintyliardów "}

var relation = map[string][]string{
	"jeden ":  singular,
	"dwa ":    pluralSimple,
	"trzy ":   pluralSimple,
	"cztery ": pluralSimple,
}

func ConvertToWordRepresentation(money string) (string, error) {
	if money == "0" {
		return "zero", nil
	}

	sb := &strings.Builder{}
	sanitizedMoney, _ := sanitizeAndSplit(money)
	convertIntegerPart(sanitizedMoney, sb)

	return strings.TrimSpace(sb.String()), nil
}

func sanitizeAndSplit(input string) (string, error) {
	//allowed format is : [0-9]+[\.,]?[0-9]{0,2}
	//allowed characters are : [0-9\.,]
	sb := strings.Builder{}
	shouldStopOnNextComa := false
	digitsAfterComa := 0
	for _, r := range input {
		if digitsAfterComa == 2 {
			return sb.String(), nil
		}
		if unicode.IsDigit(r) {
			if shouldStopOnNextComa {
				digitsAfterComa++
			}
			sb.WriteRune(r)
			continue
		}
		if r == ',' || r == '.' {
			if shouldStopOnNextComa {
				return sb.String(), nil
			}
			sb.WriteRune(r)
			shouldStopOnNextComa = true
		}
	}
	return sb.String(), nil
}

func convertIntegerPart(money string, sb *strings.Builder) {
	triplets := toTriplets(money)
	outputTriplet := make([]string, 0, len(triplets))
	for i, triplet := range triplets {
		amount, _ := tripletToAmount(triplet)

		if relation[amount] != nil {
			amount = amount + relation[amount][i]
		} else if amount != "" {
			amount = amount + pluralUpper[i]
		}
		outputTriplet = append(outputTriplet, amount)
	}

	for i := len(outputTriplet); i > 0; i-- {
		sb.WriteString(outputTriplet[i-1])
	}
}

func tripletToAmount(triplet string) (string, error) {
	sb := strings.Builder{}
	var tempTriplet = triplet
	if len(tempTriplet) == 3 {
		if tempTriplet[0] != '0' {
			sb.WriteString(hundredsMap[string(tempTriplet[0])] + " ")
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 2 {
		if tempTriplet[0] == '1' {
			sb.WriteString(tenthsMap[tempTriplet] + " ")
			return sb.String(), nil
		} else if tempTriplet[0] != '0' {
			sb.WriteString(upperTenthsMap[string(tempTriplet[0])] + " ")
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 1 && tempTriplet[0] != '0' {
		sb.WriteString(unitMap[string(tempTriplet[0])] + " ")
	}
	return sb.String(), nil
}

func toTriplets(money string) []string {
	tempMoney := money
	tripletNumber := int(math.Ceil(float64(len(money)) / 3.0))
	triplets := make([]string, 0, tripletNumber)
	for i := tripletNumber; i > 0; i-- {
		right := i * 3
		if right > len(tempMoney) {
			right = len(tempMoney)
		}
		left := right - 3
		if left < 0 {
			left = 0
		}
		triplets = append(triplets, tempMoney[left:right])
		tempMoney = tempMoney[:left]
	}
	return triplets
}
