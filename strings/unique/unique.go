package unique

// Unique returns a copy of an array of strings with duplicates removed
func Unique(slice []string) []string {
	var unique []string
	m := map[string]bool{}

	for _, v := range slice {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
