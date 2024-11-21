package utils

func CheckEmpty(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}
	return true
}
