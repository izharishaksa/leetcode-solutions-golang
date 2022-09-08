//https://leetcode.com/problems/minimum-window-substring/

package leetcode_solutions_golang

func minWindow(s string, t string) string {
	targetCount := [58]int{}
	curCount := [58]int{}
	for _, i := range t {
		targetCount[i-'A']++
	}
	left, right := 0, 0
	answer := ""
	for left < len(s) {
		for right < len(s) && !isTargetMet(curCount, targetCount) {
			if targetCount[s[right]-'A'] > 0 {
				curCount[s[right]-'A']++
			}
			right++
		}
		if isTargetMet(curCount, targetCount) && (answer == "" || len(answer) > right-left) {
			answer = s[left:right]
		}
		if targetCount[s[left]-'A'] > 0 {
			curCount[s[left]-'A']--
		}
		left++
	}
	return answer
}

func isTargetMet(curCount, targetCount [58]int) bool {
	for i := 0; i < 58; i++ {
		if curCount[i] < targetCount[i] {
			return false
		}
	}
	return true
}
