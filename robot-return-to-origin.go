package leetcode_solutions_golang

//https://leetcode.com/problems/robot-return-to-origin/
func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for i := 0; i < len(moves); i++ {
		c := moves[i]
		switch c {
		case 'R':
			x++
			break
		case 'L':
			x--
			break
		case 'U':
			y++
			break
		case 'D':
			y--
			break
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}
