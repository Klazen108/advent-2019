package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Challenge3_1(wireListing string) string {
	intersectionSet := make(map[Point]bool)
	wires := strings.Split(wireListing, "\n")
	//doing pairs, so go from first wire to second-to-last
	for curWireIndex := 0; curWireIndex < len(wires)-1; curWireIndex++ {
		//doing pairs, so go from one after current wire to last
		for compWireIndex := curWireIndex + 1; compWireIndex < len(wires); compWireIndex++ {
			wireIntersections := CheckWires(wires[curWireIndex], wires[compWireIndex])
			for _, intersection := range wireIntersections {
				intersectionSet[intersection] = true
			}
		}
	}
	//var minPoint Point
	minDist := 99999999
	for intersection := range intersectionSet {
		if intersection == (Point{0, 0}) {
			continue
		}
		curDist := ManhattanDistance(intersection, Point{0, 0})
		if curDist < minDist {
			//minPoint = intersection
			minDist = curDist
		}
	}
	return strconv.Itoa(minDist)
}

func ManhattanDistance(a Point, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func CheckWires(wire1 string, wire2 string) []Point {
	w1Points := GetWirePoints(wire1)
	w2Points := GetWirePoints(wire2)

	w1Lines := make([]Line, 0)
	for i := 0; i < len(w1Points)-1; i++ {
		w1Lines = append(w1Lines, Line{w1Points[i], w1Points[i+1]})
	}
	w2Lines := make([]Line, 0)
	for i := 0; i < len(w2Points)-1; i++ {
		w2Lines = append(w2Lines, Line{w2Points[i], w2Points[i+1]})
	}

	intersections := make([]Point, 0)
	for _, line1 := range w1Lines {
		for _, line2 := range w2Lines {
			intersection, doesIntersect := CheckLineIntersection(line1, line2)
			if doesIntersect {
				intersections = append(intersections, intersection)
			}
		}
	}
	return intersections
}

func CheckLineIntersection(line1 Line, line2 Line) (Point, bool) {
	//assume the lines are horizontal and vertical only

	//first, check endpoints
	if line1.start == line2.start {
		return line1.start, true
	}
	if line1.start == line2.end {
		return line1.start, true
	}
	if line1.end == line2.start {
		return line1.end, true
	}
	if line1.end == line2.end {
		return line1.end, true
	}

	//If the lines are collinear but their endpoints do not meet, then they do not cross
	//this does not account for overlap, we assume that doesn't happen
	if line1.GetAlignment() == line2.GetAlignment() {
		return Point{0, 0}, false
	}

	return CheckCross(line1, line2)
}

func CheckCross(line1 Line, line2 Line) (Point, bool) {
	var hLine Line
	var vLine Line

	if line1.GetAlignment() == Horizontal {
		hLine = line1
		vLine = line2
	} else {
		hLine = line2
		vLine = line1
	}

	//l2 must be vertical, so take the x of line 1 and y of line 2
	minX1 := min(hLine.start.x, hLine.end.x)
	maxX1 := max(hLine.start.x, hLine.end.x)
	minY2 := min(vLine.start.y, vLine.end.y)
	maxY2 := max(vLine.start.y, vLine.end.y)
	yAxis := hLine.start.y //line 1 is horizontal, Y is same
	xAxis := vLine.start.x //line 2 is vertical, X is same
	if minY2 <= yAxis && yAxis <= maxY2 {
		if minX1 < xAxis && xAxis < maxX1 {
			return Point{xAxis, yAxis}, true
		}
	}
	return Point{0, 0}, false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetWirePoints(wire string) []Point {
	moves := strings.Split(wire, ",")
	points := make([]Point, 0)
	points = append(points, Point{0, 0})
	for _, move := range moves {
		direction := string(move[0])
		distance, err := strconv.Atoi(move[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading distance in move: %s\n", move)
			os.Exit(104)
		}
		curPoint := points[len(points)-1]
		switch direction {
		case "U":
			curPoint.y -= distance
		case "D":
			curPoint.y += distance
		case "L":
			curPoint.x -= distance
		case "R":
			curPoint.x += distance
		}
		points = append(points, curPoint)
	}
	return points
}

type Point struct {
	x int //I deeply regret int64s so lets not
	y int
}

type Line struct {
	start Point
	end   Point
}

func (l Line) GetAlignment() Alignment {
	if l.start.y == l.end.y {
		return Horizontal
	}
	return Vertical
}

type Alignment int

const (
	Horizontal Alignment = 0
	Vertical   Alignment = 1
)
