package main

import "fmt"

func Challenge12_1() string {
	io := MoonAt("io", 5, 4, 4)
	europa := MoonAt("europa", -11, -11, -3)
	ganymede := MoonAt("ganymede", 0, 7, 0)
	callisto := MoonAt("callisto", -13, 2, 10)

	moons := []Body{io, europa, ganymede, callisto}
	for i := 0; i < 1000; i++ {
		moons = Tick(moons)
	}
	//fmt.Printf("%+v", moons)
	totEnergy := 0.0
	for _, moon := range moons {
		totEnergy += moon.TotalEnergy()
	}
	return fmt.Sprintf("%.4f", totEnergy)
}

func MoonAt(name string, x, y, z float64) Body {
	return Body{name, Vector3{x, y, z}, Vector3{0, 0, 0}}
}

type Body struct {
	name     string
	position Vector3
	velocity Vector3
}

type Vector3 struct {
	x float64
	y float64
	z float64
}

func Tick(world []Body) []Body {
	//Apply gravity
	for ai, _ := range world {
		for bi, _ := range world {
			if ai >= bi { //skip already processed pairs & self pairing
				continue
			}
			//fmt.Printf("%d,%d -> BEFORE a: %v b: %v\n", ai, bi, world[ai], world[bi])
			world[ai], world[bi] = Attract(world[ai], world[bi])
			//fmt.Printf("%d,%d -> AFTER  a: %v b: %v\n", ai, bi, world[ai], world[bi])
		}
	}

	//Apply velocity
	for ai, _ := range world {
		world[ai] = Advance(world[ai])
	}
	return world
}

func (moon Body) PotentialEnergy() float64 {
	return Absf(moon.position.x) + Absf(moon.position.y) + Absf(moon.position.z)
}

func (moon Body) KineticEnergy() float64 {
	return Absf(moon.velocity.x) + Absf(moon.velocity.y) + Absf(moon.velocity.z)
}

func (moon Body) TotalEnergy() float64 {
	return moon.KineticEnergy() * moon.PotentialEnergy()
}

func Absf(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func Advance(moon Body) Body {
	moon.position.x += moon.velocity.x
	moon.position.y += moon.velocity.y
	moon.position.z += moon.velocity.z
	return moon
}

func Attract(moonA Body, moonB Body) (Body, Body) {
	moonA, moonB = AttractX(moonA, moonB)
	moonA, moonB = AttractY(moonA, moonB)
	moonA, moonB = AttractZ(moonA, moonB)
	return moonA, moonB
}

func AttractP(posA float64, posB float64) (float64, float64) {
	if posA < posB {
		return 1, -1
	} else if posA > posB {
		return -1, 1
	}
	return 0, 0
}

func AttractX(moonA Body, moonB Body) (Body, Body) {
	adjAX, adjBX := AttractP(moonA.position.x, moonB.position.x)
	moonA.velocity.x += adjAX
	moonB.velocity.x += adjBX
	return moonA, moonB
}

func AttractY(moonA Body, moonB Body) (Body, Body) {
	adjAY, adjBY := AttractP(moonA.position.y, moonB.position.y)
	moonA.velocity.y += adjAY
	moonB.velocity.y += adjBY
	return moonA, moonB
}

func AttractZ(moonA Body, moonB Body) (Body, Body) {
	adjAZ, adjBZ := AttractP(moonA.position.z, moonB.position.z)
	moonA.velocity.z += adjAZ
	moonB.velocity.z += adjBZ
	return moonA, moonB
}
