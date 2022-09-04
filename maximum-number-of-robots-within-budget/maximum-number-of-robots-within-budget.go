//https://leetcode.com/problems/maximum-number-of-robots-within-budget/

package maximum_number_of_robots_within_budget

import (
	"container/heap"
	"fmt"
)

type ChargeTime struct {
	Value int
	Index int
}

type Heap []ChargeTime

func (t Heap) Len() int {
	return len(t)
}

func (t Heap) Less(i, j int) bool {
	return t[i].Value > t[j].Value
}

func (t Heap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *Heap) Push(x interface{}) {
	*t = append(*t, x.(ChargeTime))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	maxLen := 0
	right := 0
	pq := &Heap{}
	heap.Init(pq)
	usedRobot := make(map[int]bool)
	usedBudget := int64(0)
	fmt.Println(usedBudget)
	totalCost := int64(0)
	for i := 0; i < len(chargeTimes); i++ {
		left := i
		isOk := true
		for isOk && right-left <= maxLen {
			isOk = false
			if pq.Len() == 0 {
				if int64(chargeTimes[i]+runningCosts[i]) <= budget {
					usedBudget = int64(chargeTimes[i] + runningCosts[i])
					totalCost += int64(runningCosts[i])
					heap.Push(pq, ChargeTime{chargeTimes[i], i})
					usedRobot[i] = true
					right = i
					isOk = true
					maxLen = 1
				}
			} else if right+1 < len(chargeTimes) {
				maxChargeTime := -1
				for len(*pq) > 0 {
					top := heap.Pop(pq).(ChargeTime)
					if _, isUsed := usedRobot[top.Index]; !isUsed {
						continue
					}
					maxChargeTime = top.Value
					heap.Push(pq, top)
					break
				}
				if chargeTimes[right+1] > maxChargeTime {
					maxChargeTime = chargeTimes[right+1]
				}
				newBudget := int64(maxChargeTime) + int64(len(usedRobot)+1)*(totalCost+int64(runningCosts[right+1]))
				if newBudget <= budget {
					right++
					heap.Push(pq, ChargeTime{chargeTimes[right], right})
					totalCost += int64(runningCosts[right])
					usedBudget = newBudget
					usedRobot[right] = true
					isOk = true
				} else {
					totalCost -= int64(runningCosts[left])
					delete(usedRobot, i)
					maxChargeTime := -1
					for len(*pq) > 0 {
						top := heap.Pop(pq).(ChargeTime)
						if _, isUsed := usedRobot[top.Index]; !isUsed {
							continue
						}
						maxChargeTime = top.Value
						heap.Push(pq, top)
						break
					}
					usedBudget = int64(maxChargeTime) + int64(len(usedRobot))*totalCost

					break
				}
			}
			if right != left {
				maxLen = max(maxLen, right-left+1)
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
