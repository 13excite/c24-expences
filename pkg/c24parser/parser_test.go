package c24parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateTransactionType(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		input    string
		expected string
	}{
		{"Abbuchung", "Debit"},
		{"Zinszahlung", "Interest"},
		{"Kartenzahlung", "Card"},
		{"Pocket-Umbuchung", "Pocket"},
		{"SEPA-Überweisung", "SEPA"},
		{"SEPA-Lastschrift", "SEPA_debit"},
		{"Echtzeit-Überweisung", "Transfer"},
		{"Online-Kartenzahlung", "Online"},
		{"Unbekannter Typ", "Unbekannter Typ"},
	}

	for _, test := range tests {
		result := parser.translateTransactionType(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestParseAmount(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		input    string
		expected float64
		err      bool
	}{
		{"1234,56", 1234.56, false},
		{"1,23", 1.23, false},
		{"-37,20", -37.20, false},
		{"1234", 1234.0, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := parser.parseAmount(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestParseDate(t *testing.T) {
	parser := &Parser{}

	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"11.01.2025", "2025-01-11", false},
		{"2422.12.2021", "", true},
	}
	for _, test := range tests {
		parsedDate, err := parser.parseDate(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, parsedDate)
		}
	}
}
