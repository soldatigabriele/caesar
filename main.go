package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	caesar "github.com/soldatigabriele/caesar/caesar"
)

// This app will translate a sentence using the Caesar Cipher.
// https://en.wikipedia.org/wiki/Caesar_cipher
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded text: ")
	// Read the encoded string and convert it to upper case
	enc, _ := reader.ReadString('\n')
	enc = strings.ToUpper(enc)

	// Get the list of possible solutions
	solutions := caesar.Decode(enc)

	// Now we want to count the english words in
	// every sentence, so we can guess the correct one
	solutionsCount := make(map[string]int)
	for _, solution := range solutions {
		solutionsCount[solution] = 0
		words := strings.Split(solution, string(32))
		for _, word := range words {
			val, err := checkWord(word)
			if err != nil {
				fmt.Println(err)
				break
			}
			if val == true {
				solutionsCount[solution]++
			}
		}

	}
	for k, v := range solutionsCount {
		if v == 0 {
			delete(solutionsCount, k)
		}
	}

	s := sortedKeys(solutionsCount)
	if len(s) == 0 {
		fmt.Println("Could not find any solution to this encryption")

		return
	}

	fmt.Printf("The solution is likely to be: \n\n %s \n\n", s[0])
	if len(s) > 1 {
		fmt.Printf("but it could also be one of the following: \n\n")
		for i := 1; i < len(s); i++ {
			fmt.Printf("%s \n", s[i])
		}
	}
}

// checkWord checks if a word is a valid English word by
// looking at a dictionary of 3000 popular English words
func checkWord(w string) (bool, error) {
	// format the word to be lower case
	w = strings.ToLower(w)
	// open the dictionary of popular English words
	f, err := os.Open("dictionary.txt")
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if scanner.Text() == w {
			return true, nil
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	// could not find any match, but no error is provided
	return false, nil
}

// Define a sorted Map to get the string with the
// highest possibility of being an english sentence
type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
