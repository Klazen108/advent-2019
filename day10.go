package main

import (
	"fmt"
	"strings"
)

func Challenge10_1(mapStr string) Detector {
	return GetBestPosition(GetAsteroids(mapStr))
}

func Challenge10_2(mapStr string) string {
	return ""
}

func GetAsteroids(mapStr string) []*Detector {
	asteroidList := make([]*Detector, 0)
	for y, line := range strings.Split(strings.ReplaceAll(mapStr, "\r", ""), "\n") {
		for x, char := range line {
			if string(char) == "#" {
				asteroidList = append(asteroidList, &Detector{Point{x, y}, 0})
			}
		}
	}
	return asteroidList
}

func GetVisibleAsteroids(asteroid Detector, asteroidList []*Detector) []*Detector {
	detectedAsteroids := make([]*Detector, 0)
	for _, target := range asteroidList {
		if asteroid.position == target.position {
			continue
		}
		sightLine := Line{asteroid.position, target.position}

		isBlocked := false
		for _, blocker := range asteroidList {
			if blocker.position == asteroid.position || blocker.position == target.position {
				continue
			}
			if IsOnLine(sightLine, blocker.position) {
				//fmt.Printf("%+v -> %+v blocked by %+v\n", detector.position, target.position, blocker.position)
				isBlocked = true
				break
			}
		}
		if !isBlocked {
			fmt.Printf("%+v -> %+v detected\n", asteroid.position, target.position)
			detectedAsteroids = append(detectedAsteroids, target)
		}
	}
	return detectedAsteroids
}

func GetBestPosition(asteroidList []*Detector) Detector {
	for _, detector := range asteroidList {
		asteroids := GetVisibleAsteroids(*detector, asteroidList)
		detector.detectedAsteroids = len(asteroids)
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
