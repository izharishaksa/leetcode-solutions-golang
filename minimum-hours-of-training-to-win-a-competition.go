//https://leetcode.com/problems/minimum-hours-of-training-to-win-a-competition/

package leetcode_solutions_golang

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	total := 0
	for i := 0; i < len(energy); i++ {
		initialEnergy -= energy[i]
		if initialEnergy <= 0 {
			total += -initialEnergy + 1
			initialEnergy = 1
		}
		diff := experience[i] - initialExperience
		if diff >= 0 {
			total += diff + 1
			initialExperience += diff + 1
		}
		initialExperience += experience[i]
	}
	return total
}
