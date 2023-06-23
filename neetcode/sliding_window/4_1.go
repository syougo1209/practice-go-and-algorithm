package main

func checkInclusion(s1 string, s2 string) bool {
	s1hash := make([]int, 26)
	s2hash := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		s1hash[s1[i]-'a']++
		s2hash[s2[i]-'a']++
	}
	matches := 0

	for i := 0; i < 26; i++ {
		if s1hash[i] == s2hash[i] {
			matches++
		}
	}
	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			break
		}
		sr := s2[r] - 'a'
		s2hash[sr]++
		if s1hash[sr] == s2hash[sr] {
			matches++
		} else if s2hash[sr]-1 == s1hash[sr] {
			matches--
		}

		sl := s2[l] - 'a'
		s2hash[sl]--
		if s1hash[sl] == s2hash[sl] {
			matches++
		} else if s2hash[sl]+1 == s1hash[sl] {
			matches--
		}
		l++
	}
	return matches == 26
}
