package main

import (
	"strconv"
	"strings"
)

func Challenge4_1(rangeStr string) string {
	valStrs := strings.Split(rangeStr, "-")
	min, _ := strconv.Atoi(valStrs[0])
	max, _ := strconv.Atoi(valStrs[1])

	count := 0
	for i := min; i <= max; i++ {
		curStr := strconv.Itoa(i)
		if MeetsRules(curStr) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func MeetsRules(val string) bool {
	hasAdjacent := false
	monoIncrease := true
	curMin := 0
	last := -1
	for _, char := range val {
		cVal := int(char - '0')
		if cVal < curMin {
			monoIncrease = false
			break
		}
		if cVal == last {
			hasAdjacent = true
		}
		last = cVal
		curMin = max(curMin, cVal)
	}
	return hasAdjacent && monoIncrease
}
