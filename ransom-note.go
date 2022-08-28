//https://leetcode.com/problems/ransom-note/

package leetcode_solutions_golang

func canConstruct(ransomNote string, magazine string) bool {
	totalA := [26]int{}
	for _, c := range ransomNote {
		totalA[c-'a']++
	}

	totalB := [26]int{}
	for _, c := range magazine {
		totalB[c-'a']++
	}

	for key, i := range totalA {
		if i > totalB[key] {
			return false
		}
	}

	return true
}
