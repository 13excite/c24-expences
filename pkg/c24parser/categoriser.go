// Package c24parser provides functions to categorize transactions
// and subcategories from the C24 CSV file.
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
		// saving for household expenses
		if strings.Contains(recipient, "Haushalt") {
			return "Rent"
		}
		return "Savings"
	case "Versicherungen":
		return "Insurance"
	case "Freizeit & Unterhaltung":
		return advancedCategoriser(recipient)
	// if category is "Weitere Ausgaben" categorise based on recipient
	// to get more specific category
	case "Weitere Ausgaben":
		return advancedCategoriser(recipient)
	case "Weitere Einnahmen":
		return "Other_Income"
	case "Wohnen & Haushalt":
		if strings.Contains(recipient, "Norbert") {
			return "Rent"
		}
		return "Housing"
	case "Wellness & Beauty":
		return "Beauty"
	default:
		return sanitizeToSnakeCase(germanCategory)
	}
}

// advancedCategoriser categorises based on the recipient
// CHANGE IT IF YOU WANT TO ADD MORE CATEGORIES
func advancedCategoriser(recipient string) string {
	if strings.Contains(recipient, "Espresso House") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "reisen_urlaub") {
		return "Travel_Vacation"
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
	if strings.Contains(recipient, "HYUNDAI") || strings.Contains(recipient, "Hyundai") {
		return "Mobility"
	}
	if strings.Contains(recipient, "Anastasiia") || strings.Contains(recipient, "ANASTASIIA") {
		return "Nastya"
	}
	if strings.Contains(recipient, "SIHOO") {
		return "Housing"
	}
	if strings.Contains(recipient, "Herzensbackere") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "OVHcloud") {
		return "Housing"
	}
	if strings.Contains(recipient, "DOMKELLER") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "WEINBAUER") {
		return "Travel_Vacation"
	}
	if strings.Contains(recipient, "Vinothek") {
		return "Restaurant_Cafe"
	}
	if strings.Contains(recipient, "DATART") {
		return "Housing"
	}
	if strings.Contains(recipient, "METZGEREI") {
		return "Groceries"
	}
	if strings.Contains(recipient, "Richter Erz") {
		return "Groceries"
	}
	if strings.Contains(recipient, "Solntcev") || strings.Contains(recipient, "Haushalt") {
		return "Rent"
	}
	if strings.Contains(recipient, "Norbert") {
		return "Rent"
	}
	if strings.Contains(recipient, "KLIVER") {
		return "Groceries"
	}
	return "Other"
}

// translateSubcategory translates the German subcategory to English and converts to snake_case
func translateSubcategory(germanSubcategory, recipient string) string {
	// TODO: Improve parsing of subcategories
	if germanSubcategory == "Weitere Ausgaben" || germanSubcategory == "Saving" {
		return advancedCategoriser(recipient)
	}

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
		"Supermarkt":                   "Supermarket", // nolint:all
		"Umbuchung":                    "Saving",
		"Weitere Einnahmen":            "Other_income",
		"Öffentlicher Nahverkehr":      "Public_transport",
		"friseur":                      "Haircut",
		"Behörden":                     "Authorities",
		"Erstattung":                   "Refund",
		"Bonus Energievertrag":         "Energy_bonus",
		"Getränkehandel":               "Supermarket",
		"Heimwerken & Garten":          "Building_garden",
		"hotel_urlaubswohnungen":       "Hotel_vacation",
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
