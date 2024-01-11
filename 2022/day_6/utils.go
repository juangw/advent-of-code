package day_6

func containsDuplicate(strings []string) bool {
	disct := make(map[string]bool, len(strings))

	for _, str := range strings {
		if disct[str] {
			return true
		} else {
			disct[str] = true
		}
	}
	return false
}
