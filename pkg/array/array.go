package array

// Contains - returns true when "arr" contains "str"
func Contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}

	return false
}
