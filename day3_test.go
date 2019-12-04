package main

import (
	"testing"
)

func TestDay3Example1(t *testing.T) {
	result := Challenge3_1("R8,U5,L5,D3\nU7,R6,D4,L4")
	if result != "6" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "6")
	}
}

func TestDay3Example2(t *testing.T) {
	result := Challenge3_1("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	if result != "159" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "159")
	}
}

func TestDay3Horizontal1Vertical2(t *testing.T) {
	intersection, doesIntersect := CheckLineIntersection(
		Line{
			Point{0, 0}, Point{5, 0},
		},
		Line{
			Point{2, -1}, Point{2, 1},
		},
	)
	if !doesIntersect {
		t.Errorf("Expected intersection")
	}
	if intersection != (Point{2, 0}) {
		t.Errorf("Expected intersection at %v was %v", Point{2, 0}, intersection)
	}
}

func TestDay3Horizontal1Vertical2Flipped(t *testing.T) {
	intersection, doesIntersect := CheckLineIntersection(
		Line{
			Point{0, 0}, Point{-5, 0},
		},
		Line{
			Point{-2, 1}, Point{-2, -1},
		},
	)
	if !doesIntersect {
		t.Errorf("Expected intersection")
	}
	if intersection != (Point{-2, 0}) {
		t.Errorf("Expected intersection at %v was %v", Point{-2, 0}, intersection)
	}
}

func TestDay3Vertical1Horizontal2(t *testing.T) {
	intersection, doesIntersect := CheckLineIntersection(
		Line{
			Point{0, 0}, Point{0, -5},
		},
		Line{
			Point{-2, -1}, Point{2, -1},
		},
	)
	if !doesIntersect {
		t.Errorf("Expected intersection")
	}
	if intersection != (Point{0, -1}) {
		t.Errorf("Expected intersection at %v was %v", Point{0, -1}, intersection)
	}
}

func TestDay3Vertical1Horizontal2Flipped(t *testing.T) {
	intersection, doesIntersect := CheckLineIntersection(
		Line{
			Point{0, 0}, Point{0, 5},
		},
		Line{
			Point{2, 1}, Point{-2, 1},
		},
	)
	if !doesIntersect {
		t.Errorf("Expected intersection")
	}
	if intersection != (Point{0, 1}) {
		t.Errorf("Expected intersection at %v was %v", Point{0, 1}, intersection)
	}
}

func TestDay3Vertical1Horizontal2NoIntersectRight(t *testing.T) {
	intersection, doesIntersect := CheckLineIntersection(
		Line{
			Point{0, 0}, Point{0, -5},
		},
		Line{
			Point{2, -1}, Point{5, -1},
		},
	)
	if doesIntersect {
		t.Errorf("Expected no intersection, got one at %v", intersection)
	}
}
