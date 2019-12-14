package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Challenge10_1(mapStr string) Detector {
	return GetBestPosition(GetAsteroids(mapStr))
}

func Challenge10_2(mapStr string, target int) Detector {
	asteroidList := GetAsteroids(mapStr)
	detector := GetBestPosition(asteroidList)
	//fmt.Printf("best detector: %+v\n", detector)
	i := 0
	for {
		visibleAsteroids := GetVisibleAsteroids(detector, asteroidList)
		fmt.Printf("Visible Asteroids: %d\n", len(visibleAsteroids))
		//if the number of remaining asteroids won't reach target after they're all gone
		//then wipe em out and go next loop
		if i+len(visibleAsteroids) < target {
			asteroidList = Remove(asteroidList, visibleAsteroids)
			i += len(visibleAsteroids)
			//fmt.Printf("skippin'")
			continue
		}
		//otherwise this loop is the one
		sightLines := GetSightLines(detector, visibleAsteroids)
		//sort by angle
		sort.Slice(sightLines, func(i, j int) bool { return sightLines[i].angle < sightLines[j].angle })
		//for _, ass := range sightLines {
		//fmt.Printf("\t%+v\n", ass)
		//}
		return sightLines[target-1-i].asteroid
	}
}

func Remove(haystack []*Detector, needles []Detector) []*Detector {
	for _, needle := range needles {
		for i, hay := range haystack {
			if hay.position == needle.position {
				haystack = append(haystack[:i], haystack[i+1:]...)
			}
		}
	}
	return haystack
}

func GetSightLines(detector Detector, asteroids []Detector) []SightLine {
	lines := make([]SightLine, 0)
	for _, asteroid := range asteroids {
		source := detector.position
		target := asteroid.position

		//invert source & target y, because y is down but should be up
		angle := math.Atan2(float64(source.y)-float64(target.y), float64(target.x)-float64(source.x)) - (math.Pi / 2) //up is 0
		if angle < 0 {
			angle = 2*math.Pi + angle
		} //clockwise
		if angle != 0 {
			angle = 2*math.Pi - angle
		}

		sl := SightLine{angle, asteroid}
		lines = append(lines, sl)
	}
	return lines
}

type SightLine struct {
	angle    float64
	asteroid Detector
}

func GetAsteroids(mapStr string) []*Detector {
	mapStr = strings.Trim(mapStr, " \r\n")
	asteroidList := make([]*Detector, 0)
	for y, line := range strings.Split(strings.ReplaceAll(mapStr, "\r", ""), "\n") {
		line = strings.Trim(line, " \r\n")
		for x, char := range line {
			if string(char) == "#" {
				asteroidList = append(asteroidList, &Detector{Point{x, y}, 0})
			}
		}
	}
	return asteroidList
}

func GetVisibleAsteroids(asteroid Detector, asteroidList []*Detector) []Detector {
	detectedAsteroids := make([]Detector, 0)
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
			//fmt.Printf("%+v -> %+v detected\n", asteroid.position, target.position)
			detectedAsteroids = append(detectedAsteroids, *target)
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
