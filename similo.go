package similo

import "strings"

// FindSimilarity returns a number between 0 and 1 that represents
// the similarity between 2 strings.
func FindSimilarity(first, second string) float64 {
	if first == second {
		return 1.0
	}

	matches := 0.0
	total := float64(len(first))

	if len(second) > len(first) {
		total = float64(len(second))
	}

	for i := 0; i < len(first); i++ {
		if i < len(second) && first[i] == second[i] {
			matches = matches + 1.0
		}
	}

	return matches / total
}

// ContainsSimilars Looks on the lookup string for similars to the base
// string that mach the acceptance or higher, returns true if it has at
// least one
func ContainsSimilars(base string, lookup string, acceptance float64) bool {
	similars := FindSimilars(base, lookup, acceptance)
	return len(similars) > 0
}

// FindSimilars Looks on the lookup string for similars to the base
// string that mach the acceptance or higher, and returns these inside
// a string array
func FindSimilars(base string, lookup string, acceptance float64) []string {
	total := float64(len(base))
	matches := []string{}

	for start := 0; start < len(base); start++ {
		for end := start + 1; end <= len(base); end++ {
			substr := base[start:end]

			if strings.Contains(lookup, substr) {
				match := float64(len(substr)) / total
				if match >= acceptance {
					matches = append(matches, substr)
				}
			}
		}
	}

	return matches
}
