package c24parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateCategory(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Finanzen & Steuern", "Finance_Taxes"},
		{"DSL & Mobilfunk", "DSL_Mobile"},
		{"Einkommen", "Income"},
		{"Energie", "Energy"},
		{"Lebensmittel", "Groceries"},
		{"Mobilität", "Mobility"},
		{"Restaurant/ Café/ Bar", "Restaurant_Cafe"},
		{"Umbuchung", "Savings"},
		{"Versicherungen", "Insurance"},
		{"Weitere Ausgaben", "Other"},
		{"Weitere Einnahmen", "Other_Income"},
		{"Wohnen & Haushalt", "Housing"},
		{"Unexpected Category", "unexpected_category"},
	}

	for _, test := range tests {
		result := translateCategory(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestTranslateSubcategory(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Bäckerei", "bakery"},
		{"Drogerie", "drugstore"},
		{"Einrichtung & Haushaltswaren", "household_goods"},
		{"Elektrohandel", "electronics_store"},
		{"Festnetz, Internet und TV", "internet_tv"},
		{"Kapitalerträge", "capital_income"},
		{"Lohn/ Gehalt", "salary"},
		{"Miete", "rent"},
		{"Mobilfunk", "mobile_phone"},
		{"Restaurant/ Café/ Bar", "restaurant_cafe"},
		{"Rundfunkgebühren", "broadcast_fees"},
		{"Sonstige Versicherung", "other_insurance"},
		{"Sport Shop", "sports_shop"},
		{"Steuern und Abgaben", "taxes_and_fees"},
		{"Strom", "electricity"},
		{"Supermarkt", "supermarket"},
		{"Umbuchung", "Saving"},
		{"Weitere Ausgaben", "other_expenses"},
		{"Weitere Einnahmen", "other_income"},
		{"Öffentlicher Nahverkehr", "public_transport"},
		{"Unbekannte Unterkategorie", "unbekannte_unterkategorie"},
	}

	for _, test := range tests {
		result := translateSubcategory(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestSanitizeToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Finanzen & Steuern", "finanzen_steuern"},
		{"DSL & Mobilfunk", "dsl_mobilfunk"},
		{"Einkommen", "einkommen"},
		{"Energie", "energie"},
		{"Lebensmittel", "lebensmittel"},
		{"Mobilität", "mobilitt"},
		{"Restaurant/ Café/ Bar", "restaurant_caf_bar"},
		{"Umbuchung", "umbuchung"},
		{"Versicherungen", "versicherungen"},
		{"Weitere Ausgaben", "weitere_ausgaben"},
		{"Weitere Einnahmen", "weitere_einnahmen"},
		{"Wohnen & Haushalt", "wohnen_haushalt"},
		{"Unbekannte Kategorie", "unbekannte_kategorie"},
	}

	for _, test := range tests {
		result := sanitizeToSnakeCase(test.input)
		assert.Equal(t, test.expected, result)
	}
}
