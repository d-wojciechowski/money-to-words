package main

import "testing"

func TestUnits(t *testing.T) {
	expectedResult := map[string]string{
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

	for key, value := range expectedResult {
		result, err := ConvertToWordRepresentation(key)
		if err != nil {
			t.Errorf("Unit conversion should not fail, but error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Unit conversion failed. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestTenths(t *testing.T) {
	expectedResult := map[string]string{
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

	for key, value := range expectedResult {
		result, err := ConvertToWordRepresentation(key)
		if err != nil {
			t.Errorf("Unit conversion should not fail, but error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Unit conversion failed. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestUpperTenths(t *testing.T) {
	expectedResult := map[string]string{
		"20": "dwadzieścia",
		"21": "dwadzieścia jeden",
		"32": "trzydzieści dwa",
		"43": "czterdzieści trzy",
		"54": "pięćdziesiąt cztery",
		"65": "sześćdziesiąt pięć",
		"76": "siedemdziesiąt sześć",
		"87": "osiemdziesiąt siedem",
		"98": "dziewięćdziesiąt osiem",
	}

	for key, value := range expectedResult {
		result, err := ConvertToWordRepresentation(key)
		if err != nil {
			t.Errorf("Unit conversion should not fail, but error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Unit conversion failed. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestHundreds(t *testing.T) {
	expectedResult := map[string]string{
		"100": "sto",
		"101": "sto jeden",
		"111": "sto jedenaście",
		"221": "dwieście dwadzieścia jeden",
		"321": "trzysta dwadzieścia jeden",
		"432": "czterysta trzydzieści dwa",
		"543": "pięćset czterdzieści trzy",
		"654": "sześćset pięćdziesiąt cztery",
		"765": "siedemset sześćdziesiąt pięć",
		"876": "osiemset siedemdziesiąt sześć",
		"987": "dziewięćset osiemdziesiąt siedem",
		"999": "dziewięćset dziewięćdziesiąt dziewięć",
	}

	for key, value := range expectedResult {
		result, err := ConvertToWordRepresentation(key)
		if err != nil {
			t.Errorf("Unit conversion should not fail, but error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Unit conversion failed. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestSingularMoreThanTriples(t *testing.T) {
	expectedResult := map[string]string{
		"1000":       "jeden tysiąc",
		"1000000":    "jeden milion",
		"1000000000": "jeden miliard",
	}

	for key, value := range expectedResult {
		result, err := ConvertToWordRepresentation(key)
		if err != nil {
			t.Errorf("Unit conversion should not fail, but error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Unit conversion failed. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestPluralMoreThanTriples(t *testing.T) {
	expectedResult := map[string]string{
		"1001":          "jeden tysiąc jeden",
		"1101":          "jeden tysiąc sto jeden",
		"2101":          "dwa tysiące sto jeden",
		"3101":          "trzy tysiące sto jeden",
		"4101":          "cztery tysiące sto jeden",
		"5101":          "pięć tysięcy sto jeden",
		"9101":          "dziewięć tysięcy sto jeden",
		"1001001":       "jeden milion jeden tysiąc jeden",
		"2001101":       "dwa miliony jeden tysiąc sto jeden",
		"3002101":       "trzy miliony dwa tysiące sto jeden",
		"4003101":       "cztery miliony trzy tysiące sto jeden",
		"5004101":       "pięć milionów cztery tysiące sto jeden",
		"6005101":       "sześć milionów pięć tysięcy sto jeden",
		"9009101":       "dziewięć milionów dziewięć tysięcy sto jeden",
		"1001001001":    "jeden miliard jeden milion jeden tysiąc jeden",
		"2002001101":    "dwa miliardy dwa miliony jeden tysiąc sto jeden",
		"3003002101":    "trzy miliardy trzy miliony dwa tysiące sto jeden",
		"4004003101":    "cztery miliardy cztery miliony trzy tysiące sto jeden",
		"5005004101":    "pięć miliardów pięć milionów cztery tysiące sto jeden",
		"6006005101":    "sześć miliardów sześć milionów pięć tysięcy sto jeden",
		"9009009101":    "dziewięć miliardów dziewięć milionów dziewięć tysięcy sto jeden",
		"1001001001001": "jeden bilion jeden miliard jeden milion jeden tysiąc jeden",
		"2002002001101": "dwa biliony dwa miliardy dwa miliony jeden tysiąc sto jeden",
		"3003003002101": "trzy biliony trzy miliardy trzy miliony dwa tysiące sto jeden",
		"4004004003101": "cztery biliony cztery miliardy cztery miliony trzy tysiące sto jeden",
		"5005005004101": "pięć bilionów pięć miliardów pięć milionów cztery tysiące sto jeden",
		"6006006005101": "sześć bilionów sześć miliardów sześć milionów pięć tysięcy sto jeden",
		"9009009009101": "dziewięć bilionów dziewięć miliardów dziewięć milionów dziewięć tysięcy sto jeden",
	}

	for key, value := range expectedResult {
		result, err := ConvertToWordRepresentation(key)
		if err != nil {
			t.Errorf("Unit conversion should not fail, but error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Unit conversion failed. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func BenchmarkPluralMoreThanTriples(b *testing.B) {
	input := []string{
		"1001", "1101", "2101", "3101", "4101", "5101", "9101", "1001001", "2001101", "3002101", "4003101", "5004101",
		"6005101", "9009101", "1001001001", "2002001101", "3003002101", "4004003101", "5005004101", "6006005101",
		"9009009101", "1001001001001", "2002002001101", "3003003002101", "4004004003101", "5005005004101",
		"6006006005101", "9009009009101",
	}

	for i := 0; i < b.N; i++ {
		for _, key := range input {
			ConvertToWordRepresentation(key)
		}
	}
}
