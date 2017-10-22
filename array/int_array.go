package array

import (
	"sort"
)

// return the largest sum of sub array
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	preSum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum := 0
		if preSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum = nums[i] + preSum
		}
		if currentSum > max {
			max = currentSum
		}
		preSum = currentSum
	}
	return max
}

type Interval struct {
	Start int
	End   int
}

// merge intervals to avoid intersect
func MergeInterval(intervals []Interval) []Interval {
	var ret []Interval
	if len(intervals) < 1 {
		return ret
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= ret[len(ret)-1].End {
			if intervals[i].End > ret[len(ret)-1].End {
				ret[len(ret)-1].End = intervals[i].End
			}
		} else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

// insert a new interval to a non-overlapping intervals
func InsertInterval(intervals []Interval, newInterval Interval) []Interval {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var ret []Interval
	ret = append(ret, newInterval)
	for i := 0; i < len(intervals); i++ {
		// flag represent whether intervals[i] intersect with ret
		flag := false
		for j := 0; j < len(ret); j++ {
			if intervals[i].Start <= ret[j].Start {
				if ret[j].End <= intervals[i].End {
					ret[j].Start, ret[j].End = intervals[i].Start, intervals[i].End
					flag = true
					break
				} else if ret[j].Start <= intervals[i].End {
					ret[j].Start = intervals[i].Start
					flag = true
					break
				}
			} else {
				if intervals[i].End <= ret[j].End {
					flag = true
					break
				} else if intervals[i].Start <= ret[j].End {
					ret[j].End = intervals[i].End
					flag = true
					break
				}
			}
		}
		if !flag {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}
