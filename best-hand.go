package leetcode_solutions_golang

func bestHand(ranks []int, suits []byte) string {
	isFlush := func(suits []byte) bool {
		for i := 1; i < 5; i++ {
			if suits[i] != suits[0] {
				return false
			}
		}
		return true
	}
	if isFlush(suits) {
		return "Flush"
	}

	isThreeOfAKind := func(ranks []int) bool {
		count := make([]int, 14)
		for _, i := range ranks {
			count[i]++
		}
		for _, i := range count {
			if i >= 3 {
				return true
			}
		}
		return false
	}
	if isThreeOfAKind(ranks) {
		return "Three of a Kind"
	}

	isPair := func(ranks []int) bool {
		for i := 0; i < 5; i++ {
			for j := i + 1; j < 5; j++ {
				if ranks[i] == ranks[j] {
					return true
				}
			}
		}
		return false
	}
	if isPair(ranks) {
		return "Pair"
	}

	return "High Card"
}
