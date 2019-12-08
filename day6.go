package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Challenge6_1(mapData string) string {
	satMap := ParseMapData(mapData)
	com := satMap.GetAdd("COM")
	orbitCount := com.CountOrbits()
	return strconv.Itoa(orbitCount)
}

func ParseMapData(mapData string) SatMap {
	entries := strings.Split(mapData, "\n")

	satMap := SatMap(make(map[string]*Object))

	for _, entry := range entries {
		parsed := strings.Split(entry, ")")
		planetName := parsed[0]
		satName := parsed[1]
		planet := satMap.GetAdd(planetName)
		satellite := satMap.GetAdd(satName)
		fmt.Printf("%s <- %s\n", planetName, satName)
		planet.satellites = append(planet.satellites, satellite)
	}

	return satMap
}

func (objMap SatMap) GetAdd(name string) *Object {
	object, ok := objMap[name]
	if !ok {
		object = &Object{name, make([]*Object, 0)}
		objMap[name] = object
	}
	return object
}

type SatMap map[string]*Object

type Object struct {
	name       string
	satellites []*Object
}

func (s Object) CountOrbits() int {
	directOrbits := len(s.satellites)
	indirectOrbits := 0
	for _, satellite := range s.satellites {
		indirectOrbits += satellite.CountOrbits()
	}
	fmt.Printf("%s: %d + %d\n", s.name, directOrbits, indirectOrbits)
	return directOrbits + indirectOrbits
}
