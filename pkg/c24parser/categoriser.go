package c24parser

import (
	"regexp"
	"strings"
)

// translateCategory translates the German category to English and converts to snake_case
func translateCategory(germanCategory string) string {
	// TODO: Improve parsing of categories
	switch germanCategory {
	// skip Shopping
	case "Finanzen & Steuern":
		return "Finance_Taxes"
	case "DSL & Mobilfunk":
		return "DSL_Mobile"
	case "Einkommen":
		return "Income"
	case "Energie":
		return "Energy"
	case "Lebensmittel":
		return "Groceries"
	case "Mobilität":
		return "Mobility"
	case "Restaurant/ Café/ Bar":
		return "Restaurant_Cafe"
	case "Umbuchung":
		return "Savings"
	case "Versicherungen":
		return "Insurance"
	case "Weitere Ausgaben":
		return "Other"
	case "Weitere Einnahmen":
		return "Other_Income"
	case "Wohnen & Haushalt":
		return "Housing"
	default:
		return sanitizeToSnakeCase(germanCategory)
	}
}

// translateSubcategory translates the German subcategory to English and converts to snake_case
func translateSubcategory(germanSubcategory string) string {
	// TODO: Improve parsing of subcategories
	translationMap := map[string]string{
		"Bäckerei":                     "bakery",
		"Drogerie":                     "drugstore",
		"Einrichtung & Haushaltswaren": "household_goods",
		"Elektrohandel":                "electronics_store",
		"Festnetz, Internet und TV":    "internet_tv",
		"Kapitalerträge":               "capital_income",
		"Lohn/ Gehalt":                 "salary",
		"Miete":                        "rent",
		"Mobilfunk":                    "mobile_phone",
		"Restaurant/ Café/ Bar":        "restaurant_cafe",
		"Rundfunkgebühren":             "broadcast_fees",
		"Sonstige Versicherung":        "other_insurance",
		"Sport Shop":                   "sports_shop",
		"Steuern und Abgaben":          "taxes_and_fees",
		"Strom":                        "electricity",
		"Supermarkt":                   "supermarket",
		"Umbuchung":                    "Saving",
		"Weitere Ausgaben":             "other_expenses",
		"Weitere Einnahmen":            "other_income",
		"Öffentlicher Nahverkehr":      "public_transport",
	}

	if translation, exists := translationMap[germanSubcategory]; exists {
		return translation
	}
	return sanitizeToSnakeCase(germanSubcategory)
}

// sanitizeToSnakeCase removes spaces, special characters, and converts to snake_case
func sanitizeToSnakeCase(input string) string {
	// Replace spaces and special characters with underscores
	re := regexp.MustCompile(`[\\s&/-]+`)
	input = re.ReplaceAllString(input, "_")
	// Remove all non-alphanumeric or underscore characters
	input = regexp.MustCompile(`[^a-zA-Z0-9_]`).ReplaceAllString(input, "")
	// Convert to lowercase
	return strings.ToLower(input)
}
