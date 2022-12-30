package main

import (
	"ammount-in-words/pkg/dict"
	"ammount-in-words/pkg/utils"
	"math"
	"strings"
)

func ConvertToWordRepresentation(money string) (string, error) {
	if money == "0" {
		return "zero", nil
	}

	sb := &strings.Builder{}
	sanitizedMoney, _ := utils.SanitizeAsMoney(money)
	splitedMoney := utils.Split(sanitizedMoney, []rune{',', '.'})
	if splitedMoney[0] != "" {
		convertIntegerPart(splitedMoney[0], sb)
	}

	return strings.TrimSpace(sb.String()), nil
}

func convertIntegerPart(money string, sb *strings.Builder) {
	triplets := splitToTriplets(money)
	outputTriplet := make([]string, 0, len(triplets))
	for i, triplet := range triplets {
		amount, _ := tripletToAmount(triplet)

		if dict.Relation[amount] != nil {
			amount = amount + dict.Relation[amount][i]
		} else if amount != "" {
			amount = amount + dict.PluralUpper[i]
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
			sb.WriteString(dict.HundredsMap[string(tempTriplet[0])] + " ")
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 2 {
		if tempTriplet[0] == '1' {
			sb.WriteString(dict.TenthsMap[tempTriplet] + " ")
			return sb.String(), nil
		} else if tempTriplet[0] != '0' {
			sb.WriteString(dict.UpperTenthsMap[string(tempTriplet[0])] + " ")
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 1 && tempTriplet[0] != '0' {
		sb.WriteString(dict.UnitMap[string(tempTriplet[0])] + " ")
	}
	return sb.String(), nil
}

func splitToTriplets(money string) []string {
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
