package main

import (
	"fmt"
	"strconv"

	"github.com/klazen108/advent-2019/intcode"
)

func Challenge11_1(pgmStr string) string {
	program := ParseProgram(pgmStr)
	computer := intcode.NewComputer(program)
	position := Point{0, 0}
	worldMap := make(map[Point]int)
	direction := Up
	doGo := true

	for doGo {
		//fmt.Printf("at %+v\n", position)
		computer.ProvideInput(intcode.Byte(worldMap[position]))
		computer.Execute()
		color := computer.GetOutput()
		curDir := computer.GetOutput()
		worldMap[position] = int(color)
		if curDir == 0 {
			curDir = -1
		}
		direction += Direction(curDir)
		if direction < 0 {
			direction += 4
		}
		if direction > 3 {
			direction -= 4
		}
		switch direction {
		case Up:
			position = Point{position.x, position.y + 1}
		case Right:
			position = Point{position.x + 1, position.y}
		case Down:
			position = Point{position.x, position.y - 1}
		case Left:
			position = Point{position.x - 1, position.y}
		}
		doGo = computer.IsPendingInput()
	}
	return strconv.Itoa(len(worldMap))
}

func Challenge11_2(pgmStr string) string {
	program := ParseProgram(pgmStr)
	computer := intcode.NewComputer(program)
	position := Point{0, 0}
	worldMap := make(map[Point]int)
	worldMap[Point{0, 0}] = 1
	direction := Up
	doGo := true

	for doGo {
		//fmt.Printf("at %+v\n", position)
		computer.ProvideInput(intcode.Byte(worldMap[position]))
		computer.Execute()
		color := computer.GetOutput()
		curDir := computer.GetOutput()
		worldMap[position] = int(color)
		if curDir == 0 {
			curDir = -1
		}
		direction += Direction(curDir)
		if direction < 0 {
			direction += 4
		}
		if direction > 3 {
			direction -= 4
		}
		switch direction {
		case Up:
			position = Point{position.x, position.y + 1}
		case Right:
			position = Point{position.x + 1, position.y}
		case Down:
			position = Point{position.x, position.y - 1}
		case Left:
			position = Point{position.x - 1, position.y}
		}
		doGo = computer.IsPendingInput()
	}
	fmt.Printf("painted")
	left := 9999999
	bottom := 9999999
	right := -9999999
	top := -9999999
	for point := range worldMap {
		if point.x < left {
			left = point.x
		}
		if point.x > right {
			right = point.x
		}
		if point.y < bottom {
			bottom = point.y
		}
		if point.y > top {
			top = point.y
		}
	}
	mapImage := ""
	for y := top; y >= bottom; y-- {
		for x := left; x <= right; x++ {
			if worldMap[Point{x, y}] == 1 {
				mapImage += "#"
			} else {
				mapImage += " "
			}
		}
		mapImage += "\n"
	}
	return mapImage
}

type Direction int

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)
