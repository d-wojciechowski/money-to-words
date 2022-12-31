package converters

import "testing"

func TestUnits(t *testing.T) {
	expectedResult := map[string]string{
		"0": "zero złotych",
		"1": "jeden złoty",
		"2": "dwa złote",
		"3": "trzy złote",
		"4": "cztery złote",
		"5": "pięć złotych",
		"6": "sześć złotych",
		"7": "siedem złotych",
		"8": "osiem złotych",
		"9": "dziewięć złotych",
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
		"10": "dziesięć złotych",
		"11": "jedenaście złotych",
		"12": "dwanaście złotych",
		"13": "trzynaście złotych",
		"14": "czternaście złotych",
		"15": "piętnaście złotych",
		"16": "szesnaście złotych",
		"17": "siedemnaście złotych",
		"18": "osiemnaście złotych",
		"19": "dziewiętnaście złotych",
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
		"20": "dwadzieścia złotych",
		"21": "dwadzieścia jeden złoty",
		"32": "trzydzieści dwa złote",
		"43": "czterdzieści trzy złote",
		"54": "pięćdziesiąt cztery złote",
		"65": "sześćdziesiąt pięć złotych",
		"76": "siedemdziesiąt sześć złotych",
		"87": "osiemdziesiąt siedem złotych",
		"98": "dziewięćdziesiąt osiem złotych",
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
		"100": "sto złotych",
		"101": "sto jeden złoty",
		"111": "sto jedenaście złotych",
		"221": "dwieście dwadzieścia jeden złoty",
		"321": "trzysta dwadzieścia jeden złoty",
		"432": "czterysta trzydzieści dwa złote",
		"543": "pięćset czterdzieści trzy złote",
		"654": "sześćset pięćdziesiąt cztery złote",
		"765": "siedemset sześćdziesiąt pięć złotych",
		"876": "osiemset siedemdziesiąt sześć złotych",
		"987": "dziewięćset osiemdziesiąt siedem złotych",
		"999": "dziewięćset dziewięćdziesiąt dziewięć złotych",
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
		"1000":       "jeden tysiąc złotych",
		"1000000":    "jeden milion złotych",
		"1000000000": "jeden miliard złotych",
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
		"1001":          "jeden tysiąc jeden złoty",
		"1101":          "jeden tysiąc sto jeden złoty",
		"2101":          "dwa tysiące sto jeden złoty",
		"3101":          "trzy tysiące sto jeden złoty",
		"4101":          "cztery tysiące sto jeden złoty",
		"5101":          "pięć tysięcy sto jeden złoty",
		"9101":          "dziewięć tysięcy sto jeden złoty",
		"1001001":       "jeden milion jeden tysiąc jeden złoty",
		"2001101":       "dwa miliony jeden tysiąc sto jeden złoty",
		"3002101":       "trzy miliony dwa tysiące sto jeden złoty",
		"4003101":       "cztery miliony trzy tysiące sto jeden złoty",
		"5004101":       "pięć milionów cztery tysiące sto jeden złoty",
		"6005101":       "sześć milionów pięć tysięcy sto jeden złoty",
		"9009101":       "dziewięć milionów dziewięć tysięcy sto jeden złoty",
		"1001001001":    "jeden miliard jeden milion jeden tysiąc jeden złoty",
		"2002001101":    "dwa miliardy dwa miliony jeden tysiąc sto jeden złoty",
		"3003002101":    "trzy miliardy trzy miliony dwa tysiące sto jeden złoty",
		"4004003101":    "cztery miliardy cztery miliony trzy tysiące sto jeden złoty",
		"5005004101":    "pięć miliardów pięć milionów cztery tysiące sto jeden złoty",
		"6006005101":    "sześć miliardów sześć milionów pięć tysięcy sto jeden złoty",
		"9009009101":    "dziewięć miliardów dziewięć milionów dziewięć tysięcy sto jeden złoty",
		"1001001001001": "jeden bilion jeden miliard jeden milion jeden tysiąc jeden złoty",
		"2002002001101": "dwa biliony dwa miliardy dwa miliony jeden tysiąc sto jeden złoty",
		"3003003002101": "trzy biliony trzy miliardy trzy miliony dwa tysiące sto jeden złoty",
		"4004004003101": "cztery biliony cztery miliardy cztery miliony trzy tysiące sto jeden złoty",
		"5005005004101": "pięć bilionów pięć miliardów pięć milionów cztery tysiące sto jeden złoty",
		"6006006005101": "sześć bilionów sześć miliardów sześć milionów pięć tysięcy sto jeden złoty",
		"9009009009101": "dziewięć bilionów dziewięć miliardów dziewięć milionów dziewięć tysięcy sto jeden złoty",
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

func TestUnitsGroszy(t *testing.T) {
	expectedResult := map[string]string{
		"0.0": "zero złotych zero groszy",
		"0.1": "zero złotych jeden grosz",
		"0.2": "zero złotych dwa grosze",
		"0.3": "zero złotych trzy grosze",
		"0.4": "zero złotych cztery grosze",
		"0.5": "zero złotych pięć groszy",
		"0.6": "zero złotych sześć groszy",
		"0.7": "zero złotych siedem groszy",
		"0.8": "zero złotych osiem groszy",
		"0.9": "zero złotych dziewięć groszy",
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

func TestTenthsGroszy(t *testing.T) {
	expectedResult := map[string]string{
		"0.10": "zero złotych dziesięć groszy",
		"0.11": "zero złotych jedenaście groszy",
		"0.12": "zero złotych dwanaście groszy",
		"0.13": "zero złotych trzynaście groszy",
		"0.14": "zero złotych czternaście groszy",
		"0.15": "zero złotych piętnaście groszy",
		"0.16": "zero złotych szesnaście groszy",
		"0.17": "zero złotych siedemnaście groszy",
		"0.18": "zero złotych osiemnaście groszy",
		"0.19": "zero złotych dziewiętnaście groszy",
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

func TestUpperTenthsGroszy(t *testing.T) {
	expectedResult := map[string]string{
		"0.20": "zero złotych dwadzieścia groszy",
		"0.21": "zero złotych dwadzieścia jeden groszy",
		"0.32": "zero złotych trzydzieści dwa grosze",
		"0.43": "zero złotych czterdzieści trzy grosze",
		"0.54": "zero złotych pięćdziesiąt cztery grosze",
		"0.65": "zero złotych sześćdziesiąt pięć groszy",
		"0.76": "zero złotych siedemdziesiąt sześć groszy",
		"0.87": "zero złotych osiemdziesiąt siedem groszy",
		"0.98": "zero złotych dziewięćdziesiąt osiem groszy",
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

func TestComplete(t *testing.T) {
	expectedResult := map[string]string{
		"0.20": "zero złotych dwadzieścia groszy",
		"1.21": "jeden złoty dwadzieścia jeden groszy",
		"2.32": "dwa złote trzydzieści dwa grosze",
		"3.43": "trzy złote czterdzieści trzy grosze",
		"4.54": "cztery złote pięćdziesiąt cztery grosze",
		"5.65": "pięć złotych sześćdziesiąt pięć groszy",
		"6.76": "sześć złotych siedemdziesiąt sześć groszy",
		"7.87": "siedem złotych osiemdziesiąt siedem groszy",
		"8.98": "osiem złotych dziewięćdziesiąt osiem groszy",
		"9.99": "dziewięć złotych dziewięćdziesiąt dziewięć groszy",
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
