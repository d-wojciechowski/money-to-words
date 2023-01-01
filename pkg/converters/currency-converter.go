package converters

import (
	"ammount-in-words/pkg/dict"
	"ammount-in-words/pkg/utils"
	"math"
	"strings"
)

func ConvertToWordRepresentation(input string) (string, error) {
	if input == "0" {
		return "zero złotych", nil
	}

	sb := strings.Builder{}
	splitedMoney := utils.Split(utils.SanitizeAsMoney(input), []rune{',', '.'})
	if len(splitedMoney) > 0 && splitedMoney[0] != "" {
		sb.WriteString(convertIntegerPart(splitedMoney[0]))
	}
	if len(splitedMoney) > 1 && splitedMoney[1] != "" {
		sb.WriteString(convertDecimalPart(splitedMoney[1]))
	}

	return strings.TrimSpace(sb.String()), nil
}

func convertIntegerPart(money string) string {
	sb := &strings.Builder{}
	triplets := splitToTriplets(money)
	outputTriplet := make([]string, 0, len(triplets))
	for i, triplet := range triplets {
		amount := tripletToAmount(triplet)

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
	result := sb.String()
	suffix := getZlotyPrefix(result)
	return result + suffix
}

func getZlotyPrefix(result string) string {
	if result == "" {
		return "zero złotych "
	}
	for key, value := range dict.IntegerZlotySuffix {
		if strings.HasSuffix(result, key) {
			return value
		}
	}
	return "złotych "
}

func convertDecimalPart(money string) string {
	sb := &strings.Builder{}
	triplets := splitToTriplets(money)
	outputTriplet := make([]string, 0, len(triplets))
	for i, triplet := range triplets {
		amount := tripletToAmount(triplet)

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
	result := sb.String()
	suffix := getGroszyPrefix(result)
	return result + suffix
}

func getGroszyPrefix(result string) string {
	if result == "" {
		return "zero groszy "
	}
	if result == "jeden " {
		return "grosz "
	}
	for _, value := range dict.IntegerGroszSuffix {
		if strings.HasSuffix(result, value) {
			return "grosze "
		}
	}
	return "groszy "
}

func tripletToAmount(triplet string) string {
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
			return sb.String()
		} else if tempTriplet[0] != '0' {
			sb.WriteString(dict.UpperTenthsMap[string(tempTriplet[0])] + " ")
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 1 && tempTriplet[0] != '0' {
		sb.WriteString(dict.UnitMap[string(tempTriplet[0])] + " ")
	}
	return sb.String()
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
