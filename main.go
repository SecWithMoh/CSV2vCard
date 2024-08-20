package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Contact struct {
	FirstName  string
	LastName   string
	Company    string
	JobTitle   string
	PhoneNumber string
}

func parseCSV(filename string) ([]Contact, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// استفاده از map برای ذخیره مخاطبین یکتا بر اساس شماره تلفن
	uniqueContacts := make(map[string]Contact)

	for _, record := range records[1:] { // Skip the header
		contact := Contact{
			FirstName:   record[0],
			LastName:    record[1],
			Company:     record[2],
			JobTitle:    record[3],
			PhoneNumber: record[4],
		}

		// اضافه کردن فقط مخاطبین یکتا بر اساس شماره تلفن
		if _, exists := uniqueContacts[contact.PhoneNumber]; !exists {
			uniqueContacts[contact.PhoneNumber] = contact
		}
	}

	// تبدیل map به slice از مخاطبین
	var contacts []Contact
	for _, contact := range uniqueContacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func generateVCard(contact Contact) string {
	vCard := fmt.Sprintf(
		"BEGIN:VCARD\nVERSION:3.0\nN:%s;%s;;;\nFN:%s %s\nORG:%s\nTITLE:%s\nTEL;TYPE=CELL:%s\nEND:VCARD\n",
		contact.LastName, contact.FirstName, contact.FirstName, contact.LastName, contact.Company, contact.JobTitle, contact.PhoneNumber,
	)
	return vCard
}

func saveVCard(filename string, vCards []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, vCard := range vCards {
		_, err := file.WriteString(vCard + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	csvFilename := "contacts.csv"
	vcardFilename := "contacts.vcf"

	contacts, err := parseCSV(csvFilename)
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	var vCards []string
	for _, contact := range contacts {
		vCard := generateVCard(contact)
		vCards = append(vCards, vCard)
	}

	err = saveVCard(vcardFilename, vCards)
	if err != nil {
		fmt.Println("Error saving vCard file:", err)
		return
	}

	fmt.Println("vCard file generated successfully:", vcardFilename)
}
