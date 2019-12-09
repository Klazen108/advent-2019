package main

import (
	"fmt"
	"strconv"
)

func Challenge8_1(imageData []byte) string {
	image := DecodeSIF(imageData, 25, 6)
	minLayer := image[0]
	//minLayerI := 0
	minCount := minLayer.CountDigits(0)
	for _, layer := range image {
		curCount := layer.CountDigits(0)
		//fmt.Printf("Cur layer: %d zeroes: %d\n", i, curCount)
		if curCount < minCount {
			minLayer = layer
			//minLayerI = i
			minCount = curCount
		}
	}
	//fmt.Printf("Min layer: %d\n%+v\n", minLayerI, minLayer)
	return strconv.Itoa(minLayer.CountDigits(1) * minLayer.CountDigits(2))
}

type SpaceImage []SIFLayer

type SIFLayer []byte

func DecodeSIF(bytes []byte, width int, height int) SpaceImage {
	layers := make([]SIFLayer, 0)
	layerSize := width * height
	for offset := 0; offset < len(bytes)-layerSize; offset += layerSize {
		layers = append(layers, bytes[offset:offset+layerSize])
	}
	for layerI, layer := range layers {
		for i, curByte := range layer {
			layer[i] = curByte - '0'
		}

		if len(layer) != layerSize {
			panic(fmt.Sprintf("Bad layer length in layer %d", layerI))
		}
	}
	return layers
}

func (layer SIFLayer) CountDigits(digit int) int {
	count := 0
	for _, curByte := range layer {
		if int(curByte) == digit {
			count++
		}
	}
	return count
}
