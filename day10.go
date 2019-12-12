package main

import (
	"strings"
)

func Challenge10_1(mapStr string, mode int) string {
	return ""
}

func GetBestPosition(mapStr string) Detector {
	asteroidList := make([]*Detector, 0)
	for y, line := range strings.Split(strings.ReplaceAll(mapStr, "\r", ""), "\n") {
		for x, char := range line {
			if string(char) == "#" {
				asteroidList = append(asteroidList, &Detector{Point{x, y}, 0})
			}
		}
	}

	for di, detector := range asteroidList {
		for ti, target := range asteroidList {
			if di == ti {
				continue
			}
			sightLine := Line{detector.position, target.position}

			isBlocked := false
			for bi, blocker := range asteroidList {
				if bi == di || bi == ti {
					continue
				}
				if IsOnLine(sightLine, blocker.position) {
					//fmt.Printf("%+v -> %+v blocked by %+v\n", detector.position, target.position, blocker.position)
					isBlocked = true
					break
				}
			}
			if !isBlocked {
				//fmt.Printf("%+v -> %+v detected\n", detector.position, target.position)
				detector.detectedAsteroids++
			}
		}
	}

	maxDetects := 0
	var maxDetector *Detector
	for _, detector := range asteroidList {
		//fmt.Printf("%+v\n", detector)
		if detector.detectedAsteroids > maxDetects {
			maxDetects = detector.detectedAsteroids
			maxDetector = detector
		}
	}

	return *maxDetector
}

type Detector struct {
	position          Point
	detectedAsteroids int
}

func IsOnLine(line Line, point Point) bool {
	a := line.start
	b := line.end
	c := point
	crossproduct := (float32(c.y)-float32(a.y))*(float32(b.x)-float32(a.x)) - (float32(c.x)-float32(a.x))*(float32(b.y)-float32(a.y))

	if absf(crossproduct) > 0.01 {
		return false
	}

	dotproduct := (float32(c.x)-float32(a.x))*(float32(b.x)-float32(a.x)) + (float32(c.y)-float32(a.y))*(float32(b.y)-float32(a.y))
	if dotproduct < 0 {
		return false
	}

	squaredlengthba := (float32(b.x)-float32(a.x))*(float32(b.x)-float32(a.x)) + (float32(b.y)-float32(a.y))*(float32(b.y)-float32(a.y))
	if dotproduct > squaredlengthba {
		return false
	}

	return true
}

func absf(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}
