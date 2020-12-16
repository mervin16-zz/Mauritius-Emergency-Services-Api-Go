package tests

import (
	helper "mes/helpers"
	"testing"
)

func TestGetServiceFileByLanguage(t *testing.T) {

	// Define test cases
	var tests = []struct {
		input    string
		expected string
	}{
		{"en", "services_en.json"},
		{"fr", "services_fr.json"},
		{"it", "services_en.json"},
		{"sada", "services_en.json"},
		{"", "services_en.json"},
	}

	// Iterate through each test cases
	for _, test := range tests {

		// Assert if errors
		if output := helper.GetServiceFileByLanguage(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, received {}", test.input, test.expected, output)
		}
	}
}
