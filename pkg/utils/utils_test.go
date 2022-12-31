package utils

import (
	"testing"
)

func TestCleanupOfInputForLoop(t *testing.T) {
	expectedResult := map[string]string{
		"200.00":             "200.00",
		"200,00 ":            "200,00",
		"200 , 00":           "200,00",
		"200 , 002":          "200,00",
		"A2C0D0E , G0F0H2I":  "200,00",
		"A2C0D0E , G0,F0H2I": "200,0",
	}

	for key, value := range expectedResult {
		result, err := SanitizeAsMoney(key)
		if err != nil {
			t.Errorf("Input cleanup failed when for implementation chosen, error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Input cleanup failed when for implementation chosen. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestCleanupOfInputRegex(t *testing.T) {
	expectedResult := map[string]string{
		"200.00":             "200.00",
		"200,00 ":            "200,00",
		"200 , 00":           "200,00",
		"200 , 002":          "200,00",
		"A2C0D0E , G0F0H2I":  "200,00",
		"A2C0D0E , G0F0,H2I": "200,00",
	}

	for key, value := range expectedResult {
		result, err := SanitizeAsMoney(key)
		if err != nil {
			t.Errorf("Input cleanup failed when regex implementation chosen, error: %v raised", err.Error())
		}
		if value != result {
			t.Errorf("Input cleanup failed when regex implementation chosen. Expected was: [%v], got: [%v].", value, result)
		}
	}
}

func TestSplitWithMultipleRunes(t *testing.T) {
	expectedResult := map[string][]string{
		"200.00":  {"200", "00"},
		"200,00":  {"200", "00"},
		"2,00.00": {"2", "00", "00"},
	}

	for key, value := range expectedResult {
		result := Split(key, []rune{',', '.'})
		for i := range value {
			if result[i] != value[i] {
				t.Errorf("Input cleanup failed when regex implementation chosen. Expected was: [%v], got: [%v].", value, result)
			}
		}
	}
}

func BenchmarkSanitizeInputForLoop(b *testing.B) {
	input := []string{
		"200.00",
		"200,00",
		"200 , 00",
		"200 , 002",
		"A2C0D0E , G0F0H2I",
	}

	for i := 0; i < b.N; i++ {
		for _, key := range input {
			_, _ = SanitizeAsMoney(key)
		}
	}
}
