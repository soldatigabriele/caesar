package caesar

// Decode returns a list of possible translations
// of a sentence using the Caesar Cipher. The sentence
// will be translated using all 26 possible translations
// by shifting the alphabet by one position at the time.
func Decode(input string) []string {

	letters := map[string]int{
		"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5,
		"G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11,
		"M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17,
		"S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23,
		"Y": 24, "Z": 25,
	}
	reverseLetters := reverseMap(letters)

	// solutions will store all the different decoded sentences
	var solutions []string

	// Let's shift for every letter
	for i := 1; i <= 26; i++ {
		solution := ""

		for _, v := range input {
			// If the value is not found in the map
			// (e.g. a punctuation marks), leave it as it is
			if _, ok := letters[string(v)]; ok != true {
				solution += string(v)
				continue
			}

			// If the decoded letter is past 26, let's get the rest
			// E.g. 25 -> 25, 26 -> 0, 27 -> 1, 28 -> 2
			dl := (letters[string(v)] + i) % 26
			solution += reverseLetters[dl]
		}
		solutions = append(solutions, solution)
	}
	return solutions
}

// reverseMap returns a new map with keys and values swapped
func reverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
