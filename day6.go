package main

import (
	"strconv"
	"strings"
)

func Challenge6_1(mapData string) string {
	satMap := ParseMapData(mapData)
	com := satMap.GetAdd("COM")
	orbitCount := com.CountOrbitsRoot()
	return strconv.Itoa(orbitCount)
}

func ParseMapData(mapData string) SatMap {
	entries := strings.Split(mapData, "\n")

	satMap := SatMap(make(map[string]*Object))

	for _, entry := range entries {
		parsed := strings.Split(entry, ")")
		planetName := strings.Trim(parsed[0], "\t\r\n ")
		satName := parsed[1]
		planet := satMap.GetAdd(planetName)
		satellite := satMap.GetAdd(satName)
		//fmt.Printf("%s <- %s\n", planetName, satName)
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

func (s Object) CountOrbits(depth int) int {
	indirectOrbitsDepth := 0
	for _, satellite := range s.satellites {
		indirectOrbitsDepth += satellite.CountOrbits(depth + 1)
	}
	return depth + indirectOrbitsDepth
}

func (s Object) CountOrbitsRoot() int {
	return s.CountOrbits(0)
}
