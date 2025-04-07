package c24parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateCategory(t *testing.T) {
	tests := []struct {
		input    string
		recipent string
		expected string
	}{
		{"Finanzen & Steuern", "", "Finance_Taxes"},
		{"DSL & Mobilfunk", "", "DSL_Mobile"},
		{"Einkommen", "", "Income"},
		{"Energie", "", "Energy"},
		{"Lebensmittel", "", "Groceries"},
		{"Mobilität", "", "Mobility"},
		{"Restaurant/ Café/ Bar", "", "Restaurant_Cafe"},
		{"Umbuchung", "", "Savings"},
		{"Versicherungen", "", "Insurance"},
		{"Weitere Ausgaben", "", "Other"},
		{"Weitere Einnahmen", "", "Other_Income"},
		{"Wohnen & Haushalt", "", "Housing"},
		{"Unexpected Category", "", "unexpected_category"},
		{"Weitere Ausgaben", "Espresso House", "Restaurant_Cafe"},
		{"Weitere Ausgaben", "GITHUB", "Work"},
		{"Weitere Ausgaben", "Fahrschule", "Driving_Lessons"},
		{"Weitere Ausgaben", "Viethouse", "Restaurant_Cafe"},
		{"Weitere Ausgaben", "Asia Mark", "Groceries"},
		{"Weitere Ausgaben", "JUICE FACTORY", "Restaurant_Cafe"},
		{"Weitere Ausgaben", "Unbekannter Empfänger", "Other"},
	}

	for _, test := range tests {
		result := translateCategory(test.input, test.recipent)
		assert.Equal(t, test.expected, result)
	}
}

func TestTranslateSubcategory(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Bäckerei", "Bakery"},
		{"Drogerie", "Drugstore"},
		{"Einrichtung & Haushaltswaren", "Household_goods"},
		{"Elektrohandel", "Electronics_store"},
		{"Festnetz, Internet und TV", "Internet_tv"},
		{"Kapitalerträge", "Capital_income"},
		{"Lohn/ Gehalt", "Salary"},
		{"Miete", "Rent"},
		{"Mobilfunk", "Mobile_phone"},
		{"Restaurant/ Café/ Bar", "Restaurant_cafe"},
		{"Rundfunkgebühren", "Broadcast_fees"},
		{"Sonstige Versicherung", "Other_insurance"},
		{"Sport Shop", "Sports_shop"},
		{"Steuern und Abgaben", "Taxes_and_fees"},
		{"Strom", "Electricity"},
		{"Supermarkt", "Supermarket"}, // nolint:all
		{"Umbuchung", "Saving"},
		{"Weitere Ausgaben", "Other_expenses"},
		{"Weitere Einnahmen", "Other_income"},
		{"Öffentlicher Nahverkehr", "Public_transport"},
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
