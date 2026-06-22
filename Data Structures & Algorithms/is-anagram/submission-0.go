import "maps"

func isAnagram(s string, t string) bool {

	sLetters := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sLetters[s[i]]++
	}

	tLetters := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tLetters[t[i]]++
	}

	return maps.Equal(sLetters, tLetters)
}
