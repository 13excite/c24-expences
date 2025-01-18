package c24parser

import (
	"regexp"
	"strings"
)

// translateCategory translates the German category to English and converts to snake_case
func translateCategory(germanCategory, recipient string) string {
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
		return advancedCategoriser(recipient)
	case "Weitere Einnahmen":
		return "Other_Income"
	case "Wohnen & Haushalt":
		return "Housing"
	case "Wellness & Beauty":
		return "Beauty"
	default:
		return sanitizeToSnakeCase(germanCategory)
	}
}

func advancedCategoriser(recipient string) string {
	if strings.Contains(recipient, "Espresso House") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "GITHUB") {
		return "Work"
	}
	if strings.Contains(recipient, "Fahrschule") {
		return "Driving_Lessons"
	}
	if strings.Contains(recipient, "Viethouse") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "Asia Mark") {
		return "Groceries"
	}
	if strings.Contains(recipient, "JUICE FACTORY") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "SIHOO") {
		return "Housing"
	}
	return "Other"
}

// translateSubcategory translates the German subcategory to English and converts to snake_case
func translateSubcategory(germanSubcategory string) string {
	// TODO: Improve parsing of subcategories
	translationMap := map[string]string{
		"Bäckerei":                     "Bakery",
		"Drogerie":                     "Drugstore",
		"Einrichtung & Haushaltswaren": "Household_goods",
		"Elektrohandel":                "Electronics_store",
		"Festnetz, Internet und TV":    "Internet_tv",
		"Kapitalerträge":               "Capital_income",
		"Lohn/ Gehalt":                 "Salary",
		"Miete":                        "Rent",
		"Mobilfunk":                    "Mobile_phone",
		"Restaurant/ Café/ Bar":        "Restaurant_cafe",
		"Rundfunkgebühren":             "Broadcast_fees",
		"Sonstige Versicherung":        "Other_insurance",
		"Sport Shop":                   "Sports_shop",
		"Steuern und Abgaben":          "Taxes_and_fees",
		"Strom":                        "Electricity",
		"Supermarkt":                   "Supermarket",
		"Umbuchung":                    "Saving",
		"Weitere Ausgaben":             "Other_expenses",
		"Weitere Einnahmen":            "Other_income",
		"Öffentlicher Nahverkehr":      "Public_transport",
		"friseur":                      "Haircut",
		"Behörden":                     "Authorities",
		"Erstattung":                   "Refund",
		"Bonus Energievertrag":         "Energy_bonus",
		"Getränkehandel":               "Supermarket",
		"Heimwerken & Garten":          "Building_garden",
	}

	if translation, exists := translationMap[germanSubcategory]; exists {
		return translation
	}
	return sanitizeToSnakeCase(germanSubcategory)
}

// sanitizeToSnakeCase removes spaces, special characters, and converts to snake_case
func sanitizeToSnakeCase(input string) string {
	// Replace spaces and special characters with underscores
	re := regexp.MustCompile(`[\s&/-]+`)
	input = re.ReplaceAllString(input, "_")
	// Remove all non-alphanumeric or underscore characters
	input = regexp.MustCompile(`[^a-zA-Z0-9_]`).ReplaceAllString(input, "")
	// Convert to lowercase
	return strings.ToLower(input)
}
