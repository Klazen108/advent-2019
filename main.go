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
	challenge2()
}

func challenge2() {
	moduleMasses := readModuleMasses()

	spacecraftFuel := getSpacecraftFuel(moduleMasses, true)
	fmt.Fprintf(os.Stderr, "Fuel required: %f\n", spacecraftFuel)
}

func challenge1() {
	moduleMasses := readModuleMasses()

	spacecraftFuel := getSpacecraftFuel(moduleMasses, false)
	fmt.Fprintf(os.Stderr, "Fuel required: %f\n", spacecraftFuel)
}

func readModuleMasses() []float64 {
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
	return moduleMasses
}

func getModuleFuel(moduleMass float64, includeFuelFuel bool) float64 {
	moduleFuel := math.Floor(moduleMass/3) - 2
	if moduleFuel <= 0 {
		return 0
	}
	if includeFuelFuel {
		moduleFuel = moduleFuel + getModuleFuel(moduleFuel, includeFuelFuel)
	}
	return moduleFuel
}

func getSpacecraftFuel(moduleMasses []float64, includeFuelFuel bool) float64 {
	moduleFuels := make([]float64, len(moduleMasses))
	total := 0.0
	for i, elem := range moduleMasses {
		moduleFuels[i] = getModuleFuel(elem, includeFuelFuel)
		total += moduleFuels[i]
	}
	return total
}
