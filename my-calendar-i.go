package leetcode_solutions_golang

//https://leetcode.com/problems/my-calendar-i/
type MyCalendar struct {
	events [][]int
}

func CalendarConstructor() MyCalendar {
	return MyCalendar{
		events: make([][]int, 0),
	}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, event := range c.events {
		if (start >= event[0] && start <= event[1]) || (end > event[0] && end <= event[1]) || (start < event[0] && end > event[1]) {
			return false
		}
	}
	c.events = append(c.events, []int{start, end - 1})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
