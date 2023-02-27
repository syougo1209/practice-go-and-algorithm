func isAnagram(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	sb_map := map[byte]int{}
	tb_map := map[byte]int{}
	if len(sb) != len(tb) {
		return false
	}
	for _, s := range sb {
		sb_map[s] += 1
	}

	for _, s := range tb {
		tb_map[s] += 1
	}
	for _, s := range sb {
		if sb_map[s] != tb_map[s] {
			return false
		}
	}
	return true
}
