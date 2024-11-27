package helpers

import "regexp"

// baki tangada mn baed
func EmailChecker(email string) bool {
	EmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return EmailRegex.MatchString(email)
}
