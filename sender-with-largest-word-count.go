package leetcode_solutions_golang

//https://leetcode.com/problems/sender-with-largest-word-count/
func largestWordCount(messages []string, senders []string) string {
	wordCountBySender := make(map[string]int)
	for i := 0; i < len(messages); i++ {
		wordCountBySender[senders[i]] += countWords(messages[i])
	}
	mostSenderName := ""
	mostSenderCount := 0
	for k, v := range wordCountBySender {
		if v > mostSenderCount {
			mostSenderName = k
			mostSenderCount = v
		} else if v == mostSenderCount {
			if k > mostSenderName {
				mostSenderName = k
			}
		}
	}
	return mostSenderName
}

func countWords(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result++
		}
	}
	return result + 1
}
