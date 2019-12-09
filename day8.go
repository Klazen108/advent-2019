package main

import (
	"fmt"
	"strconv"
)

func Challenge8_1(imageData []byte, width int, height int) string {
	image := DecodeSIF(imageData, width, height)
	minLayer := image.layers[0]
	//minLayerI := 0
	minCount := minLayer.CountDigits(0)
	for _, layer := range image.layers {
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

func Challenge8_2(imageData []byte, width int, height int) {
	image := DecodeSIF(imageData, width, height)
	image.PrintImage()
	//image.DumpImage()
}

type SpaceImage struct {
	layers []SIFLayer
	width  int
	height int
}

type SIFLayer []byte

func DecodeSIF(bytes []byte, width int, height int) SpaceImage {
	layers := make([]SIFLayer, 0)
	layerSize := width * height
	for offset := 0; offset <= len(bytes)-layerSize; offset += layerSize {
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
	return SpaceImage{layers, width, height}
}

func (image SpaceImage) DumpImage() {
	for _, layer := range image.layers {
		fmt.Printf("%+v\n", layer)
	}
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

func (image SpaceImage) PrintImage() {
	for y := 0; y < image.height; y++ {
		for x := 0; x < image.width; x++ {
			i := y*image.width + x

			abort := false
			for _, layer := range image.layers {
				switch layer[i] {
				case 1:
					fmt.Printf("⬜")
					abort = true
				case 0:
					fmt.Printf("⬛")
					abort = true
				}
				if abort {
					break
				}
			}
			if !abort {
				panic("eek")
			}
		}
		fmt.Printf("\n")
	}
}

func (image SpaceImage) PrintImage2() {
	finalImage := make([]byte, image.width*image.height)
	for i := 0; i < image.width*image.height; i++ {
		finalImage[i] = 2
	}
	for li := len(image.layers) - 1; li >= 0; li-- {
		for i := 0; i < image.width*image.height; i++ {
			layer := image.layers[li]

			if layer[i] == 2 {
				continue
			}
			finalImage[i] = layer[i]
		}
	}

	for y := 0; y < image.height; y++ {
		for x := 0; x < image.width; x++ {
			i := y*image.width + x
			switch finalImage[i] {
			case 2:
				panic("eek")
			case 1:
				fmt.Printf("⬜")
			case 0:
				fmt.Printf("⬛")
			}
		}
		fmt.Printf("\n")
	}
}
