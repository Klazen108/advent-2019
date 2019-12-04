package main

import "testing"

func TestChallenge1Example(t *testing.T) {
	total := Challenge1("100756")
	if total != "33583" {
		t.Errorf("Total was incorrect, got: %s, want: %s.", total, "33583")
	}
}

func TestChallenge1(t *testing.T) {
	total := Challenge1("83326,84939,135378,105431,119144,124375,138528,88896,98948,85072,112576,144497,112824,98892,81551,139462,73213,93261,130376,118425,132905,54627,134676,140435,131410,128441,96755,94866,89490,122118,106596,77531,84941,57494,97518,136224,69247,147209,92814,63436,79819,109335,85698,110103,79072,52282,73957,68668,105394,149663,91954,66479,55778,126377,75471,75662,71910,113031,133917,76043,65086,117882,134854,60690,67495,62434,67758,95329,123078,128541,108213,93543,147937,148262,56212,148586,73733,110763,149243,133232,95817,68261,123872,93764,147297,51555,110576,89485,109570,88052,132786,70585,105973,85898,149990,114463,147536,67786,139193,112322")
	if total != "3408471" {
		t.Errorf("Total was incorrect, got: %s, want: %s.", total, "3408471")
	}
}

func TestChallenge2Example(t *testing.T) {
	total := Challenge2("100756")
	if total != "50346" {
		t.Errorf("Total was incorrect, got: %s, want: %s.", total, "50346")
	}
}

func TestChallenge2(t *testing.T) {
	total := Challenge2("83326,84939,135378,105431,119144,124375,138528,88896,98948,85072,112576,144497,112824,98892,81551,139462,73213,93261,130376,118425,132905,54627,134676,140435,131410,128441,96755,94866,89490,122118,106596,77531,84941,57494,97518,136224,69247,147209,92814,63436,79819,109335,85698,110103,79072,52282,73957,68668,105394,149663,91954,66479,55778,126377,75471,75662,71910,113031,133917,76043,65086,117882,134854,60690,67495,62434,67758,95329,123078,128541,108213,93543,147937,148262,56212,148586,73733,110763,149243,133232,95817,68261,123872,93764,147297,51555,110576,89485,109570,88052,132786,70585,105973,85898,149990,114463,147536,67786,139193,112322")
	if total != "5109803" {
		t.Errorf("Total was incorrect, got: %s, want: %s.", total, "5109803")
	}
}
