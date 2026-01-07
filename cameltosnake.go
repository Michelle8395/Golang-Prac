func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	result := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		// reject numbers or punctuation
		if c < 'A' || (c > 'Z' && c < 'a') || c > 'z' {
			return s
		}

		// reject consecutive uppercase letters
		if i > 0 && c >= 'A' && c <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}

		// reject ending with uppercase letter
		if i == len(s)-1 && c >= 'A' && c <= 'Z' {
			return s
		}

		// insert underscore before uppercase letters (except first char)
		if i > 0 && c >= 'A' && c <= 'Z' {
			result += "_"
		}

		result += string(c)
	}

	return result
}
