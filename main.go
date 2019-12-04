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
	Challenge1_2("")
}

func Challenge1_2(massesStr string) string {
	var moduleMasses []float64
	if massesStr == "" {
		massesStr = readModuleMassesCsv()
	}
	moduleMasses = calcModuleMasses(massesStr)

	spacecraftFuel := getSpacecraftFuel(moduleMasses, true)
	fmt.Fprintf(os.Stderr, "Fuel required: %f\n", spacecraftFuel)
	return fmt.Sprintf("%.0f", spacecraftFuel)
}

func Challenge1_1(massesStr string) string {
	var moduleMasses []float64
	if massesStr == "" {
		massesStr = readModuleMassesCsv()
	}
	moduleMasses = calcModuleMasses(massesStr)

	spacecraftFuel := getSpacecraftFuel(moduleMasses, false)
	fmt.Fprintf(os.Stderr, "Fuel required: %f\n", spacecraftFuel)
	return fmt.Sprintf("%.0f", spacecraftFuel)
}

func readModuleMassesCsv() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your list of module masses, in csv format:\n>")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading module masses: %v\n", err)
		os.Exit(101)
	}
	return strings.Trim(text, "\r\n")
}

func calcModuleMasses(massCsv string) []float64 {
	moduleMassesStr := strings.Split(massCsv, ",")
	moduleMasses := make([]float64, len(moduleMassesStr))
	for i, elem := range moduleMassesStr {
		massf, err := strconv.ParseFloat(elem, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid module mass: '%s'\n", elem)
			os.Exit(102)
		}
		moduleMasses[i] = massf
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
