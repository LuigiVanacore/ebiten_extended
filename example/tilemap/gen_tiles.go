//go:build ignore
// +build ignore

// Run with: go run gen_tiles.go
// Creates tiles.png for the tilemap example.
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 64, 32))
	c1 := color.RGBA{80, 120, 80, 255}
	c2 := color.RGBA{100, 140, 100, 255}
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if x < 32 {
				img.Set(x, y, c1)
			} else {
				img.Set(x, y, c2)
			}
		}
	}
	f, _ := os.Create("tiles.png")
	defer f.Close()
	png.Encode(f, img)
}
