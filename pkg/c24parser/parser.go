package c24parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/13excite/c24-expences/pkg/models"
)

// Parser struct that holds the transactions and the CSV file
type Parser struct {
	transactions []models.Transaction
	file         *os.File
	csvReader    *csv.Reader
}

// NewParser returns a new Parser struct
func NewParser() *Parser {
	return &Parser{
		// Initialize the slice with length 0
		transactions: make([]models.Transaction, 0),
	}
}

// readCSV reads the CSV file and initializes the csv.Reader
func (p *Parser) readCSV(filename string) error {
	// Open the CSV file
	var err error
	p.file, err = os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	// Read CSV content
	p.csvReader = csv.NewReader(p.file)
	p.csvReader.Comma = ','
	return nil
}

// ParseFile parses the CSV file and stores the transactions in the Parser struct
func (p *Parser) ParseFile(filename string) error {
	// Close the file when the function returns
	defer p.file.Close()

	if err := p.readCSV(filename); err != nil {
		return err
	}
	_, err := p.csvReader.Read()
	if err != nil {
		return fmt.Errorf("error reading header: %v", err)
	}
	for {
		row, err := p.csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Error reading row: %v\n", err)
			continue
		}
		amount, err := p.parseAmount(row[2])
		if err != nil {
			fmt.Printf("Error parsing amount: %v\n", err)
			continue
		}
		// Parse date
		date, err := p.parseDate(row[1])
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			continue
		}
		var recipient string
		if row[0] == "SEPA-Überweisung" {
			recipient = strings.Split(row[3], ",")[0]
		} else {
			recipient = row[3]
		}

		p.transactions = append(p.transactions, models.Transaction{
			TransactionType: p.translateTransactionType(row[0]),
			Date:            date,
			Amount:          amount,
			Recipient:       recipient,
			Usage:           row[6],
			Category:        translateCategory(row[8]),
			Subcategory:     translateSubcategory(row[9]),
		})

	}
	return nil
}

// parseDate parses the date string to the format "YYYY-MM-DD"
func (p *Parser) parseDate(dateStr string) (string, error) {
	parsedDate, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		return "", err
	}
	return parsedDate.Format("2006-01-02"), nil
}

// GetTransactions returns the parsed transactions
func (p *Parser) GetTransactions() []models.Transaction {
	return p.transactions
}

// translateTransactionType translates the German transaction type to English
func (p *Parser) translateTransactionType(germanType string) string {
	switch germanType {
	case "Abbuchung":
		return "Debit"
	case "Zinszahlung":
		return "Interest"
	case "Kartenzahlung":
		return "Card"
	case "Pocket-Umbuchung":
		return "Pocket"
	case "SEPA-Überweisung":
		return "SEPA"
	case "SEPA-Lastschrift":
		return "SEPA_debit"
	case "Echtzeit-Überweisung":
		return "Instant_transfer"
	default:
		return germanType
	}
}

// parseAmount parses the amount string to a float64
func (p *Parser) parseAmount(amountStr string) (float64, error) {
	// Replace "," with "." and remove quotes
	replacer := strings.NewReplacer(",", ".", "\"", "")
	cleanAmount := replacer.Replace(amountStr)
	return strconv.ParseFloat(cleanAmount, 64)
}
