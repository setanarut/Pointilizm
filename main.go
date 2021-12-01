// CC BY-NC 3.0
// hazarek

package main

import (
	"fmt"
	"image"
	"math"

	"github.com/cenkalti/dominantcolor"
	"github.com/fogleman/gg"
	"github.com/hazarek/pointilizm/angles"
	"github.com/hazarek/pointilizm/matrix"
	"github.com/hazarek/pointilizm/utils"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	var pos utils.Vec
	img := utils.LoadImage("assets/lake.jpg")

	// brush directions
	dir := angles.GetAngles(img)
	dir.MapToRange(0.0, math.Pi*2)

	palette := utils.ToColorfulPalette(dominantcolor.FindN(img, 20))
	temp := palette
	ctx := gg.NewContextForImage(img)

	// plot(dir, img, pos, ctx, palette, 3, 3)

	palette = utils.VaryPalette(temp, -20, -0.2, 0.4)
	plot(dir, img, pos, ctx, palette, 3, 30)

	palette = utils.VaryPalette(temp, 0, 0, 0.1)
	plot(dir, img, pos, ctx, palette, 3, 20)

	palette = utils.VaryPalette(temp, 10, 0.2, 0)
	plot(dir, img, pos, ctx, palette, 2, 10)

	// palette = utils.VaryPalette(temp, 0, 0, 0.1)
	// plot(dir, img, pos, ctx, palette, 2, 10)

	// palette = utils.VaryPalette(temp, 0, 0.20, 0)
	// plot(dir, img, pos, ctx, palette, 3, 20)

	palette = utils.VaryPalette(temp, -20, 0.5, 0.1)
	plot(dir, img, pos, ctx, palette, 3, 20)

	// utils.PaletteToImage("assets/palette.png", palette, 30, 4)
	ctx.SavePNG("assets/lakeP2.png")
}

func plot(dir matrix.Mat, img image.Image, pos utils.Vec, ctx *gg.Context, p []colorful.Color, s float64, g int) {
	fmt.Println(p[0])
	for y := 0; y < img.Bounds().Max.Y; y += g {
		for x := 0; x < img.Bounds().Max.X; x += g {
			pos.X, pos.Y = float64(x), float64(y)
			pos.Randomize(-10.0, 10.0)
			clr := utils.NearestColor(img.At(int(pos.X), int(pos.Y)), p)
			ctx.SetColor(clr)
			ctx.Push()
			ctx.RotateAbout(dir.At(x, y), pos.X, pos.Y)
			ctx.DrawEllipse(pos.X, pos.Y, 5, 3*dir.At(x,y))
			ctx.Pop()
			ctx.Fill()
		}
	}
}
