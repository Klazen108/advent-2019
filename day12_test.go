package main

import "testing"

func TestDay12_1(t *testing.T) {
	output := Challenge12_1()
	if output != "10845.0000" {
		t.Errorf("Expected 10845.0000, got %s", output)
	}
}
func TestDay12_2(t *testing.T) {
	output := Challenge12_2()
	if output != "100" {
		t.Errorf("Expected 100, got %s", output)
	}
}

func TestAttractPPos(t *testing.T) {
	adj1, adj2 := AttractP(3, 5)
	if adj1 != 1 {
		t.Errorf("Expected adj to be 1, got %f", adj1)
	}
	if adj2 != -1 {
		t.Errorf("Expected adj to be -1, got %f", adj2)
	}
}

func TestAttractPNeg(t *testing.T) {
	adj1, adj2 := AttractP(5, 3)
	if adj1 != -1 {
		t.Errorf("Expected adj to be -1, got %f", adj1)
	}
	if adj2 != 1 {
		t.Errorf("Expected adj to be 1, got %f", adj2)
	}
}

func TestAttractPZero(t *testing.T) {
	adj1, adj2 := AttractP(5, 5)
	if adj1 != 0 {
		t.Errorf("Expected adj to be 0, got %f", adj1)
	}
	if adj2 != 0 {
		t.Errorf("Expected adj to be 0, got %f", adj2)
	}
}

func TestAttractX(t *testing.T) {
	moonA, moonB := AttractX(
		Body{"", Vector3{1, 0, 0}, Vector3{0, 0, 0}},
		Body{"", Vector3{2, 0, 0}, Vector3{0, 0, 0}},
	)
	if moonA.velocity.x != 1 {
		t.Errorf("Expected moon A X velocity to be 1, got %f", moonA.velocity.x)
	}
	if moonB.velocity.x != -1 {
		t.Errorf("Expected moon B X velocity to be -1, got %f", moonB.velocity.x)
	}
}

func TestAttractY(t *testing.T) {
	moonA, moonB := AttractY(
		Body{"", Vector3{0, 1, 0}, Vector3{0, 0, 0}},
		Body{"", Vector3{0, 2, 0}, Vector3{0, 0, 0}},
	)
	if moonA.velocity.y != 1 {
		t.Errorf("Expected moon A Y velocity to be 1, got %f", moonA.velocity.y)
	}
	if moonB.velocity.y != -1 {
		t.Errorf("Expected moon B Y velocity to be -1, got %f", moonB.velocity.y)
	}
}

func TestAttractZ(t *testing.T) {
	moonA, moonB := AttractZ(
		Body{"", Vector3{0, 0, 1}, Vector3{0, 0, 0}},
		Body{"", Vector3{0, 0, 2}, Vector3{0, 0, 0}},
	)
	if moonA.velocity.z != 1 {
		t.Errorf("Expected moon A Z velocity to be 1, got %f", moonA.velocity.z)
	}
	if moonB.velocity.z != -1 {
		t.Errorf("Expected moon B Z velocity to be -1, got %f", moonB.velocity.z)
	}
}

func TestTick(t *testing.T) {
	moons := []Body{
		Body{"", Vector3{0, 0, 1}, Vector3{0, 0, 0}},
		Body{"", Vector3{0, 0, 2}, Vector3{0, 0, 0}},
	}
	moons = Tick(moons)
	if moons[0].velocity.z != 1 {
		t.Errorf("Expected moon A Z velocity to be 1, got %f", moons[0].velocity.z)
	}
	if moons[1].velocity.z != -1 {
		t.Errorf("Expected moon B Z velocity to be -1, got %f", moons[1].velocity.z)
	}
	if moons[0].position.z != 2 {
		t.Errorf("Expected moon A Z position to be 2, got %f", moons[0].position.z)
	}
	if moons[1].position.z != 1 {
		t.Errorf("Expected moon B Z position to be 1, got %f", moons[1].position.z)
	}
}

func TestTick4(t *testing.T) {
	moons := []Body{
		MoonAt("0", -1, 0, 2),
		MoonAt("1", 2, -10, -7),
		MoonAt("2", 4, -8, 8),
		MoonAt("3", 3, 5, -1),
	}
	moons = Tick(moons)
	if moons[0].position != (Vector3{2, -1, 1}) {
		t.Errorf("Expected moon 1 position %v, got %v", (Vector3{2, -1, 1}), moons[0].position)
	}
	if moons[0].velocity != (Vector3{3, -1, -1}) {
		t.Errorf("Expected moon 1 velocity %v, got %v", (Vector3{3, -1, -1}), moons[0].velocity)
	}
}

func TestEnergy(t *testing.T) {
	moon := Body{"", Vector3{2, 1, 3}, Vector3{-3, 2, 1}}
	if moon.TotalEnergy() != 36 {
		t.Errorf("Expected moon NRG to be 36, got %f", moon.TotalEnergy())
	}
}

func Test10(t *testing.T) {
	io := MoonAt("io", -8, -10, 0)
	europa := MoonAt("europa", 5, 5, 10)
	ganymede := MoonAt("ganymede", 2, -7, 3)
	callisto := MoonAt("callisto", 9, -8, -3)

	moons := []Body{io, europa, ganymede, callisto}
	for i := 0; i < 10; i++ {
		moons = Tick(moons)
	}

	if moons[0].position != (Vector3{-9, -10, 1}) {
		t.Errorf("Expected moon 1 at %v, got %v", (Vector3{-9, -10, 1}), moons[0].position)
	}
}

func TestAttract2(t *testing.T) {
	left := MoonAt("l", 0, 0, 0)
	mid := MoonAt("m", 1, 0, 0)
	right := MoonAt("r", 2, 0, 0)
	left, mid = Attract(left, mid)
	left, right = Attract(left, right)
	if left.velocity.x != 2.0 {
		t.Errorf("Expected left moon velocity %f, got %f", 2.0, left.velocity.x)
	}
}
