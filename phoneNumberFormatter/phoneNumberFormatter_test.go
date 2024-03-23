package phoneNumberFormatter

import (
	"testing"
)

func TestPhoneNumberFormatter(t *testing.T) {
	expectedResult := "(13) 98179-7296"
	result, _ := PhoneNumberFormatter("13981797296")

	if result != expectedResult {
		t.Errorf("O resultado era pra ser %s mas saiu %s", expectedResult, result)

	}
}

func TestPhoneNumberFormatter_NoPhoneNumberProvided(t *testing.T) {
	expectedError := "no phone number provided"
	_, error := PhoneNumberFormatter("")

	if error.Error() != expectedError {
		t.Error("error diferente do esperado")
	}
}

func TestPhoneNumberFormatter_InvalidPhoneNumber(t *testing.T) {
	expectedError := "invalid phone number"
	_, error := PhoneNumberFormatter("123")

	if error.Error() != expectedError {
		t.Error("error diferente do esperado")
	}
}
