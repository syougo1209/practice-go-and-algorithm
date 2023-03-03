package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	count1, count2 := [26]int{}, [26]int{}
	for i := range s1 {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if count1[i] == count2[i] {
			matches++
		}
	}
	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		count2[s2[r]-'a']++
		if count2[s2[r]-'a'] == count1[s2[r]-'a'] {
			matches++
		} else if count2[s2[r]-'a'] == count1[s2[r]-'a']+1 {
			matches--
		}
		count2[s2[l]-'a']--
		if count2[s2[l]-'a'] == count1[s2[l]-'a'] {
			matches++
		} else if count2[s2[l]-'a']+1 == count1[s2[l]-'a'] {
			matches--
		}
		l++
	}
	return matches == 26
}
