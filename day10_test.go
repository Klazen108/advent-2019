package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Challenge10Input() string {
	return `
	..#..###....#####....###........#
	.##.##...#.#.......#......##....#
	#..#..##.#..###...##....#......##
	..####...#..##...####.#.......#.#
	...#.#.....##...#.####.#.###.#..#
	#..#..##.#.#.####.#.###.#.##.....
	#.##...##.....##.#......#.....##.
	.#..##.##.#..#....#...#...#...##.
	.#..#.....###.#..##.###.##.......
	.##...#..#####.#.#......####.....
	..##.#.#.#.###..#...#.#..##.#....
	.....#....#....##.####....#......
	.#..##.#.........#..#......###..#
	#.##....#.#..#.#....#.###...#....
	.##...##..#.#.#...###..#.#.#..###
	.#..##..##...##...#.#.#...#..#.#.
	.#..#..##.##...###.##.#......#...
	...#.....###.....#....#..#....#..
	.#...###..#......#.##.#...#.####.
	....#.##...##.#...#........#.#...
	..#.##....#..#.......##.##.....#.
	.#.#....###.#.#.#.#.#............
	#....####.##....#..###.##.#.#..#.
	......##....#.#.#...#...#..#.....
	...#.#..####.##.#.........###..##
	.......#....#.##.......#.#.###...
	...#..#.#.........#...###......#.
	.#.##.#.#.#.#........#.#.##..#...
	.......#.##.#...........#..#.#...
	.####....##..#..##.#.##.##..##...
	.#.#..###.#..#...#....#.###.#..#.
	............#...#...#.......#.#..
	.........###.#.....#..##..#.##...`
}

func TestDay10_1(t *testing.T) {
	input := Challenge10Input()
	output := Challenge10_1(input)
	if output.detectedAsteroids != 314 {
		t.Errorf("Expected 314 detected asteroids, got %+v", output)
	}
}

func TestDay10_2(t *testing.T) {
	input := Challenge10Input()
	output := Challenge10_1(input)
	fmt.Printf("%+v", output)
	output = Challenge10_2(input, 200)
	ans := strconv.Itoa(output.position.x*100 + output.position.y)

	if ans != "1000" {
		t.Errorf("Expected 314 detected asteroids, got %+v", output)
	}
}

func TestDay10_1_Example1(t *testing.T) {
	input := `
.#..#
.....
#####
....#
...##`
	output := Challenge10_1(input)
	if output.detectedAsteroids != 8 {
		t.Errorf("Expected 8 detected asteroids, got %+v", output)
	}
	if output.position.x != 3 || output.position.y != 4 {
		t.Errorf("Expected asteroid in 3,4 to be selected, got %+v", output)
	}
}

func TestDay10_1_Example2(t *testing.T) {
	input := `
......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
	output := Challenge10_1(input)
	if output.detectedAsteroids != 33 {
		t.Errorf("Expected 33 detected asteroids, got %+v", output)
	}
	if output.position.x != 5 || output.position.y != 8 {
		t.Errorf("Expected asteroid in 5,8 to be selected, got %+v", output)
	}
}

func TestDay10_2_Ex1_BestPos(t *testing.T) {
	input := `
.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	output := Challenge10_1(input)
	if output.position.x != 8 || output.position.y != 3 {
		t.Errorf("Expected asteroid in 8,3 to be selected, got %+v", output)
	}
}

func TestDay10_2_Example2(t *testing.T) {
	input := `
.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	output := Challenge10_2(input, 3) //9,1
	if output.position.x != 9 || output.position.y != 1 {
		t.Errorf("Expected 9,1, got %+v", output.position)
	}
}

func TestDay10_2_Example2_AfterLoop(t *testing.T) {
	input := `
.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	output := Challenge10_2(input, 31)
	if output.position.x != 8 || output.position.y != 0 {
		t.Errorf("Expected 8,0, got %+v", output.position)
	}
}

func TestDay10_2_ExampleBig(t *testing.T) {
	input := `
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	output := Challenge10_2(input, 1)
	if output.position.x != 11 || output.position.y != 12 {
		t.Errorf("1-Expected 11,12, got %+v", output.position)
	}
	output = Challenge10_2(input, 2)
	if output.position.x != 12 || output.position.y != 1 {
		t.Errorf("2-Expected 12,1, got %+v", output.position)
	}
	output = Challenge10_2(input, 20)
	if output.position.x != 16 || output.position.y != 0 {
		t.Errorf("20-Expected 16,0, got %+v", output.position)
	}
	output = Challenge10_2(input, 100)
	if output.position.x != 10 || output.position.y != 16 {
		t.Errorf("100-Expected 10,16, got %+v", output.position)
	}
	output = Challenge10_2(input, 201)
	if output.position.x != 10 || output.position.y != 9 {
		t.Errorf("201-Expected 10,9, got %+v", output.position)
	}
	output = Challenge10_2(input, 299)
	if output.position.x != 11 || output.position.y != 1 {
		t.Errorf("299-Expected 11,1, got %+v", output.position)
	}
}

func TestGetSightLinesUp(t *testing.T) {
	src := Detector{Point{2, 2}, 0}
	target := Detector{Point{2, 1}, 0}
	output := GetSightLines(src, []Detector{target})
	if output[0].angle != 0 {
		t.Errorf("Expected 2,1 to be at angle 0 from 2,2, got %f", output[0].angle)
	}
}

func TestGetSightLinesRight(t *testing.T) {
	src := Detector{Point{2, 2}, 0}
	target := Detector{Point{3, 2}, 0}
	output := GetSightLines(src, []Detector{target})
	if output[0].angle-math.Pi/2 > 0.001 {
		t.Errorf("Expected 3,2 to be at angle pi/2 from 2,2, got %f", output[0].angle)
	}
}

func TestGetSightLinesLeft(t *testing.T) {
	src := Detector{Point{2, 2}, 0}
	target := Detector{Point{1, 2}, 0}
	output := GetSightLines(src, []Detector{target})
	if output[0].angle-3*math.Pi/2 > 0.001 {
		t.Errorf("Expected 1,2 to be at angle 3*pi/2 from 2,2, got %f", output[0].angle)
	}
}
