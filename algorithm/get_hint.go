package algorithm

import "fmt"

func GetHint(secret string, guess string) string {
	table := make(map[byte]int)
	var a, b int
	for i := range secret {
		if secret[i] == guess[i] {
			a++
		} else {
			table[secret[i]]++
			if table[secret[i]]<1 {
				b++
			}
			table[guess[i]]--
			if table[guess[i]]>-1 {
				b++
			}
		}

	}

	return fmt.Sprintf("%dA%dB", a, b)
}
