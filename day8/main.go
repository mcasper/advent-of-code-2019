package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	width  int
	height int
	rows   [][]int
	index  int
}

func NewLayer(width, height, index int) Layer {
	var rows [][]int
	for i := 0; i < height; i++ {
		rows = append(rows, []int{})
	}
	return Layer{
		width,
		height,
		rows,
		index,
	}
}

func (l *Layer) PixelAt(index int) int {
	var flatPixels []int

	for _, row := range l.rows {
		for _, layerPixel := range row {
			flatPixels = append(flatPixels, layerPixel)
		}
	}

	return flatPixels[index]
}

func (l *Layer) AddPixel(pixel int) {
	for i, row := range l.rows {
		if len(row) == l.width {
			continue
		} else {
			l.rows[i] = append(l.rows[i], pixel)
			break
		}
	}
}

func (l *Layer) Full() bool {
	if len(l.rows) != l.height {
		return false
	}

	for _, row := range l.rows {
		if len(row) != l.width {
			return false
		}
	}

	return true
}

func (l Layer) PixelCount(pixel int) int {
	count := 0

	for _, row := range l.rows {
		for _, layerPixel := range row {
			if layerPixel == pixel {
				count++
			}
		}
	}

	return count
}

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	split := strings.Split(strings.Trim(string(input), "\n"), "")

	if err != nil {
		log.Fatal(err)
	}

	var ints []int

	for _, line := range split {
		integer, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, integer)
	}

	result := Execute(ints)
	fmt.Printf("Part 1 result: %v\n", result)

	Execute2(ints)
	fmt.Println("Open up ./result.png")
}

func Execute(ints []int) int {
	width := 25
	height := 6

	if len(ints)%(width*height) != 0 {
		log.Fatal("Don't have a complete set of layers")
	}

	layers := []Layer{}
	layer := NewLayer(width, height, 0)

	for i, pixel := range ints {
		layer.AddPixel(pixel)
		if layer.Full() {
			layers = append(layers, layer)
			layer = NewLayer(width, height, i+1)
		}
	}

	fewestZerosCount := 10000000
	var fewestZerosLayer Layer

	for _, layer := range layers {
		zerosCount := layer.PixelCount(0)
		if zerosCount < fewestZerosCount {
			fewestZerosCount = zerosCount
			fewestZerosLayer = layer
		}
	}

	return fewestZerosLayer.PixelCount(1) * fewestZerosLayer.PixelCount(2)
}

func Execute2(ints []int) {
	width := 25
	height := 6

	if len(ints)%(width*height) != 0 {
		log.Fatal("Don't have a complete set of layers")
	}

	layers := []Layer{}
	layer := NewLayer(width, height, 0)

	for i, pixel := range ints {
		layer.AddPixel(pixel)
		if layer.Full() {
			layers = append(layers, layer)
			layer = NewLayer(width, height, i+1)
		}
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for _, layer := range layers {
				pixel := layer.PixelAt(x + (width * y))
				if pixel == 2 {
					continue
				} else if pixel == 1 {
					img.Set(x, y, color.White)
					break
				} else {
					img.Set(x, y, color.Black)
					break
				}
			}
		}
	}

	f, _ := os.Create("result.png")
	png.Encode(f, img)
}
