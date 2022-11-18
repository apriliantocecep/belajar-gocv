package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	// membuat window baru
	window := gocv.NewWindow("Belajar Gambar Objek")
	defer window.Close()

	// membuat canvas
	width := 640
	height := 480
	canvas := gocv.NewMatWithSize(width, height, gocv.MatTypeCV8UC3)

	// membuat warna
	red := color.RGBA{255, 0, 0, 0}
	green := color.RGBA{0, 255, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}
	purple := color.RGBA{106, 90, 205, 0}

	// membuat garis
	gocv.Line(&canvas, image.Pt(100, 150), image.Pt(300, 400), red, 1)

	// membuat persegi
	gocv.Rectangle(&canvas, image.Rect(150, 200, 350, 400), green, 2)

	// membuat lingkaran
	gocv.Circle(&canvas, image.Pt(220, 250), 80, blue, -1)

	// membuat ellips
	gocv.Ellipse(&canvas, image.Pt(300, 500), image.Pt(100, 50), 45, 130, 270, purple, 4)

	// loop
	for {
		// tampilkan canvas di window
		window.IMShow(canvas)

		// wait key to press
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
