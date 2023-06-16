package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := range freq {
		if freq[i] != 0 {
			return false
		}
	}
	return true
}
