package leetcode_solutions_golang

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	s1Count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
	}

	s2Count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s2Count[s2[i]-'a']++
	}

	if isSame(s2Count, s1Count) {
		return true
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		s2Count[s2[i]-'a']--
		s2Count[s2[len(s1)+i]-'a']++
		if isSame(s2Count, s1Count) {
			return true
		}
	}

	return false
}

func isSame(source, target []int) bool {
	for i := 0; i < 26; i++ {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}
