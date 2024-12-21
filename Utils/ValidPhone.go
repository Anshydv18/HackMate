package utils

import "regexp"

func IsValidPhone(phone string) bool {
	PhonePattern := `^(\+91|91|0)?[6-9]\d{9}$`

	re := regexp.MustCompile(PhonePattern)

	return re.MatchString(phone)
}

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	return re.MatchString(email)
}
