package main

import (
	"fmt"
	"os"
)

var validCountryCodes = []string{"US", "VI"}

func getCountryNameFormat(countryCode string) string {
	switch countryCode {
	case "US":
		return "%[1]s %[2]s %[3]s"
	case "VI":
		return "%[2]s %[3]s %[1]s"
	default:
		return "%[1]s %[2]s %[3]s"
	}
}

func getCountryNameFormatWithoutMiddleName(countryCode string) string {
	switch countryCode {
	case "US":
		return "%[1]s %[2]s"
	case "VI":
		return "%[2]s %[1]s"
	default:
		return "%[1]s %[2]s"
	}
}

func reorderName(firstName string, lastName string, middleName string, countryCode string) string {
	format := ""

	if middleName != "" {
		format = getCountryNameFormat(countryCode)
	} else {
		format = getCountryNameFormatWithoutMiddleName(countryCode)
	}

	return fmt.Sprintf(format, firstName, lastName, middleName)
}

func validateName(firstName string, lastName string, middleName string) bool {
	if firstName == "" || lastName == "" {
		return false
	}

	return true
}

func isCountryCodeValid(countryCode string) bool {
	for _, code := range validCountryCodes {
		if code == countryCode {
			return true
		}
	}

	return false
}

// func parseArgs() (firstName string, lastName string, middleName string, countryCode string) {

// 	flag.StringVar(&firstName, "firstName", firstName, "First name")
// 	flag.StringVar(&lastName, "lastName", lastName, "Last name")
// 	flag.StringVar(&middleName, "middleName", middleName, "Middle name")
// 	flag.StringVar(&countryCode, "countryCode", countryCode, "Country code")

// 	flag.Parse()

// 	return
// }

func parseArgs() (firstName string, lastName string, middleName string, countryCode string) {

	args := os.Args[1:]

	if len(args) > 4 || len(args) < 2 {
		fmt.Println("Invalid number of arguments")
	}

	firstName = args[0]
	lastName = args[1]

	if len(args) == 4 {
		middleName = args[2]
		countryCode = args[3]
	}

	if len(args) == 3 {
		countryCode = args[2]
	}

	return
}

func main() {
	firstName, lastName, middleName, countryCode := parseArgs()

	fmt.Printf("firstName: %s, lastName: %s, middleName: %s, countryCode: %s\n", firstName, lastName, middleName, countryCode)

	if !isCountryCodeValid(countryCode) {
		fmt.Println("Invalid country code")
		return
	}

	if !validateName(firstName, lastName, middleName) {
		fmt.Println("Invalid name")
		return
	}

	name := reorderName(firstName, lastName, middleName, countryCode)
	println(name)
}
