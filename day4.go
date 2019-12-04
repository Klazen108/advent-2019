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

func Challenge4_2(rangeStr string) string {
	valStrs := strings.Split(rangeStr, "-")
	min, _ := strconv.Atoi(valStrs[0])
	max, _ := strconv.Atoi(valStrs[1])

	count := 0
	for i := min; i <= max; i++ {
		curStr := strconv.Itoa(i)
		if MeetsRules2(curStr) {
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

func MeetsRules2(val string) bool {
	if !MeetsRules(val) {
		return false
	}

	runs := make(map[int]int)
	for _, char := range val {
		cVal := int(char - '0')

		run := 0
		mapRun, ok := runs[cVal]
		if ok {
			run = mapRun
		}
		run++
		runs[cVal] = run
	}

	for _, count := range runs {
		if count == 2 {
			return true
		}
	}

	return false
}
