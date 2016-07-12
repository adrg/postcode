package postcode

func inSlice(needle string, haystack []string) bool {
	if needle == "" || len(haystack) == 0 {
		return false
	}

	for _, val := range haystack {
		if needle == val {
			return true
		}
	}

	return false
}
