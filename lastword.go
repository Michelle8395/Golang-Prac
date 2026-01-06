func LastWord(s string) string {
	result := ""
	i := len(s) - 1

	// If string is empty
	if len(s) == 0 {
		return "\n"
	}

	// Skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// If no word found
	if i < 0 {
		return "\n"
	}

	// Collect characters of the last word backwards
	for i >= 0 && s[i] != ' ' {
		result = string(s[i]) + result
		i--
	}

	return result + "\n"
}
