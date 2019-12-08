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

func Challenge6_2(mapData string) string {
	satMap := ParseMapData(mapData)

	youAncestry := satMap.GetAdd("YOU").GetAncestry()
	sanAncestry := satMap.GetAdd("SAN").GetAncestry()
	_, dist := FindLCA(youAncestry, sanAncestry)
	xferCount := dist - 2 //exclude you and santa
	return strconv.Itoa(xferCount)
}

func FindLCA(youAncestry []*Object, sanAncestry []*Object) (*Object, int) {
	//start at direct orbit
	for i := len(youAncestry) - 2; i >= 0; i-- {
		ya := youAncestry[i]
		if dist, found := Find(sanAncestry, ya); found {
			//fmt.Printf("%d/%d", len(youAncestry)-2-i, len(sanAncestry)-2-dist)
			youXfer := len(youAncestry) - 1 - i
			sanXfer := len(sanAncestry) - 1 - dist
			return sanAncestry[dist], youXfer + sanXfer
		}
	}
	return nil, 9999
}

func Find(haystack []*Object, needle *Object) (int, bool) {
	for i := len(haystack) - 1; i >= 0; i-- {
		if needle == haystack[i] {
			return i, true
		}
	}
	return 0, false
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
		satellite.parent = planet
	}

	return satMap
}

func (objMap SatMap) GetAdd(name string) *Object {
	object, ok := objMap[name]
	if !ok {
		object = &Object{name, make([]*Object, 0), nil}
		objMap[name] = object
	}
	return object
}

type SatMap map[string]*Object

type Object struct {
	name       string
	satellites []*Object
	parent     *Object
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

func (s *Object) GetAncestry() []*Object {
	if s.parent == nil {
		return make([]*Object, 0)
	}
	return append(s.parent.GetAncestry(), s)
}
