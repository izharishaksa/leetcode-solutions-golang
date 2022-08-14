//https://leetcode.com/problems/task-scheduler-ii/

package leetcode_solutions_golang

func taskSchedulerII(tasks []int, space int) int64 {
	taskType := make(map[int]int64, 0)
	curDay := int64(0)
	for i := 0; i < len(tasks); i++ {
		cur := tasks[i]
		lastDay, isExist := taskType[cur]
		if !isExist {
			curDay++
			taskType[cur] = curDay
		} else {
			if curDay-lastDay < int64(space) {
				curDay = lastDay + int64(space) + 1
				taskType[cur] = curDay
			} else {
				curDay++
				taskType[cur] = curDay
			}
		}
	}
	return curDay
}
