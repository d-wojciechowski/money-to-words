package main

import (
	"fmt"
	"math"
	"strings"
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

var singular = []string{"", "tysiąc", "milion", "miliard", "bilion", "biliard", "trylion", "tryliard", "kwadrylion", "kwadryliard", "kwintylion", "kwintyliard"}
var pluralSimple = []string{"", "tysiące", "miliony", "miliardy", "biliony", "biliardy", "tryliony", "tryliardy", "kwadryliony", "kwadryliardy", "kwintyliony", "kwintyliardy"}
var pluralUpper = []string{"", "tysięcy", "milionów", "miliardów", "bilionów", "biliardów", "trylionów", "tryliardów", "kwadrylionów", "kwadryliardów", "kwintylionów", "kwintyliardów"}

var relation = map[string][]string{
	"jeden":  singular,
	"dwa":    pluralSimple,
	"trzy":   pluralSimple,
	"cztery": pluralSimple,
}

type pair struct {
	k int
	v string
}

func ConvertToWordRepresentation(money string) (string, error) {
	if money == "0" {
		return "zero", nil
	}

	triplets := toTriplets(money)
	amountChan := make(chan *pair)
	defer close(amountChan)

	for tripletOrder, triplet := range triplets {
		go calculateAndAddTripletToChan(triplet, tripletOrder, amountChan)
	}

	outputTriplet := consumeTripletsFromChannel(len(triplets), amountChan)
	return strings.TrimSpace(strings.Join(outputTriplet, " ")), nil
}

func consumeTripletsFromChannel(expectedSize int, amountChan chan *pair) []string {
	var p *pair
	lastIdx := expectedSize - 1
	outputTriplet := make([]string, expectedSize)
	for i := 0; i < expectedSize; i++ {
		p = <-amountChan
		outputTriplet[lastIdx-p.k] = p.v
	}
	return outputTriplet
}

func calculateAndAddTripletToChan(triplet string, tripletOrder int, amountChan chan *pair) {
	amount, _ := getTripletWithSuffix(triplet, tripletOrder)
	amountChan <- &pair{
		v: amount,
		k: tripletOrder,
	}
}

func getTripletWithSuffix(triplet string, tripletOrder int) (string, error) {
	amount, _ := tripletToAmount(triplet)

	if relation[amount] != nil {
		amount = fmt.Sprintf("%s %s", amount, relation[amount][tripletOrder])
	} else if amount != "" {
		amount = fmt.Sprintf("%s %s", amount, pluralUpper[tripletOrder])
	}
	return amount, nil
}

func tripletToAmount(triplet string) (string, error) {
	convertedTriplet := make([]string, 0, 3)
	var tempTriplet = triplet
	if len(tempTriplet) == 3 {
		if tempTriplet[0] != '0' {
			convertedTriplet = append(convertedTriplet, hundredsMap[string(tempTriplet[0])])
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 2 {
		if tempTriplet[0] == '1' {
			convertedTriplet = append(convertedTriplet, tenthsMap[tempTriplet])
			return strings.Join(convertedTriplet, " "), nil
		} else if tempTriplet[0] != '0' {
			convertedTriplet = append(convertedTriplet, upperTenthsMap[string(tempTriplet[0])])
		}
		tempTriplet = tempTriplet[1:]
	}
	if len(tempTriplet) == 1 && tempTriplet[0] != '0' {
		convertedTriplet = append(convertedTriplet, unitMap[string(tempTriplet[0])])
	}
	return strings.Join(convertedTriplet, " "), nil
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
