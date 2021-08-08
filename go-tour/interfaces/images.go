package main

import (
	"fmt"
	"image"
)

// Package image defines the "Image" interface:
/*
	package image

	type Image interface {
		ColorModel() color.Model
		Bounds() Rectange
		At(x, y int) color.Color
	}
*/

// Note: The "Rectange" return value of the "Bounds" method is actually an "image.Rectange", as the declaration is inside the package "image".

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())        // (0,0)-(100,100)
	fmt.Println(m.At(0, 0).RGBA()) // 0 0 0 0
}

// The "color.Color" and "color.Model" types are also interfaces, but we'll ignore that by using predefined implementations "color.RGBA" and "color.RGBAModel".
// These interfaces and types are specified by "image/color" package.
