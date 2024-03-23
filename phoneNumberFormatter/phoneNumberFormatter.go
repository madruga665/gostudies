package phoneNumberFormatter

import (
	"errors"
	"fmt"
	"strings"
)

func PhoneNumberFormatter(phoneNumber string) (string, error) {
	if phoneNumber == "" {
		return "", errors.New("no phone number provided")
	}

	if len(phoneNumber) < 11 && len(phoneNumber) > 0 {
		return "", errors.New("invalid phone number")
	}

	splitString := strings.Split(phoneNumber, "")
	areaCode := splitString[0] + splitString[1]
	prefixNumber, uniqueNumber := "", ""

	for _, digit := range splitString[2:7] {
		prefixNumber += digit
	}

	for _, digit := range splitString[7:11] {
		uniqueNumber += digit
	}

	result := fmt.Sprintf("(%s) %s-%s", areaCode, prefixNumber, uniqueNumber)

	return result, nil
}
