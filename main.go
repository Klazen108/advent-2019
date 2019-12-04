package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	challenge1()
}

func challenge1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your list of module masses, in csv format:\n>")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading module masses: %v\n", err)
		os.Exit(101)
	}
	moduleMassesStr := strings.Split(strings.Trim(text, "\r\n"), ",")
	moduleMasses := make([]float64, len(moduleMassesStr))
	for i, elem := range moduleMassesStr {
		moduleMasses[i], err = strconv.ParseFloat(elem, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid module mass: '%s'\n", elem)
			os.Exit(102)
		}
	}

	spacecraftFuel := getSpacecraftFuel(moduleMasses)
	fmt.Fprintf(os.Stderr, "Fuel required: %f\n", spacecraftFuel)
}

func getModuleFuel(moduleMass float64) float64 {
	return math.Floor(moduleMass/3) - 2
}

func getSpacecraftFuel(moduleMasses []float64) float64 {
	moduleFuels := make([]float64, len(moduleMasses))
	total := 0.0
	for i, elem := range moduleMasses {
		moduleFuels[i] = getModuleFuel(elem)
		total += moduleFuels[i]
	}
	return total
}
