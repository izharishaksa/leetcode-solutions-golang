//https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/

package leetcode_solutions_golang

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	s := make(map[string]bool, 0)
	for _, v := range supplies {
		s[v] = true
	}
	for _, v := range recipes {
		s[v] = false
	}
	result := make([]string, 0)
	for {
		total := 0
		for i := 0; i < len(ingredients); i++ {
			val, isCreated := s[recipes[i]]
			isCreated = val
			if isCreated {
				continue
			}
			isCreated = true
			for _, v := range ingredients[i] {
				val, isAvailable := s[v]
				isAvailable = val
				if !isAvailable {
					isCreated = false
					break
				}
			}
			if isCreated {
				total++
				s[recipes[i]] = true
				result = append(result, recipes[i])
			}
		}

		if total == 0 {
			break
		}
	}

	return result
}
