package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var rules = uint8(110)

func nextState(index uint8) uint8 {
	return rules >> index & 1
}

func calculateIndex(prev, current, next uint8) uint8 {
	ret := prev<<2 | current<<1 | next<<0
	return ret
}

func main() {
	size := flag.Int("size", 100, "image width")
	filename := flag.String("output", "final.png", "png output result")
	if *size < 2 {
		fmt.Fprintf(os.Stderr, "image size too small\n")
		flag.Usage()
		os.Exit(0)
	}
	if *filename == "" {
		fmt.Fprintf(os.Stderr, "output must not be empty\n")
		flag.Usage()
		os.Exit(0)
	}

	flag.Parse()
	fd, err := os.Create(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create new image file %v failed: %v\n", *filename, err)
		os.Exit(1)
	}
	defer fd.Close()

	currentStates := make([]uint8, *size, *size)
	currentStates[len(currentStates)-2] = 1

	img := image.NewRGBA(image.Rect(0, 0, *size-1, *size-1))

	for h := 0; h < *size-2; h++ {
		nextStates := make([]uint8, *size, *size)
		for index := 1; index < *size-1; index++ {
			prev := currentStates[index-1]
			current := currentStates[index]
			next := currentStates[index+1]

			if current == 1 {
				img.Set(h, index-1, color.Black)
			} else {
				img.Set(h, index-1, color.White)
			}
			nextStates[index] = nextState(calculateIndex(prev, current, next))
		}
		currentStates = nextStates
	}
	err = png.Encode(fd, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encoding image failed: %v\n", err)
		os.Exit(1)
	}
}
