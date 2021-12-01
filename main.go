package main

import (
	"fmt"
	"image"
	"math"

	"github.com/cenkalti/dominantcolor"
	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
	"github.com/hazarek/pointilizm/angles"
	"github.com/hazarek/pointilizm/matrix"
	"github.com/hazarek/pointilizm/utils"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	median := gift.New(gift.Median(7, false))

	var pos utils.Vec
	img := utils.LoadImage("assets/lake.jpg")

	// brush directions
	dir := angles.GetAngles(img)
	dir.MapToRange(0.0, math.Pi*2)

	palette := utils.ToColorfulPalette(dominantcolor.FindN(img, 20))
	temp := palette


	medianImage := image.NewRGBA(img.Bounds())
	median.Draw(medianImage, img)
	ctx := gg.NewContextForImage(medianImage)

	// plot(dir, img, pos, ctx, palette, 3, 3)

	palette = utils.VaryPalette(temp, -20, -0.2, 0.4)
	plot(dir, img, pos, ctx, palette, 3, 30)

	palette = utils.VaryPalette(temp, 0, 0, 0.1)
	plot(dir, img, pos, ctx, palette, 3, 20)

	palette = utils.VaryPalette(temp, 10, 0.2, 0)
	plot(dir, img, pos, ctx, palette, 2, 10)

	palette = utils.VaryPalette(temp, 0, -0.3, 0.1)
	plot(dir, img, pos, ctx, palette, 2, 10)

	palette = utils.VaryPalette(temp, 0, 0.20, 0)
	plot(dir, img, pos, ctx, palette, 3, 20)

	palette = utils.VaryPalette(temp, -20, 0.5, 0.1)
	plot(dir, img, pos, ctx, palette, 3, 10)

	// utils.PaletteToImage("assets/palette.png", palette, 30, 4)
	ctx.SavePNG("assets/lakeP.png")
}
//g = grid resolution, s = scale
func plot(dir matrix.Mat, img image.Image, pos utils.Vec, ctx *gg.Context, p []colorful.Color, s float64, g int) {
	fmt.Println("drawing pass")
	for y := 0; y < img.Bounds().Max.Y; y += g {
		for x := 0; x < img.Bounds().Max.X; x += g {
			pos.X, pos.Y = float64(x), float64(y)
			pos.Randomize(-10.0, 10.0)
			clr := utils.NearestColor(img.At(int(pos.X), int(pos.Y)), p)
			rclr := utils.VaryColor(clr, utils.RandRange(-10.0, 10.0), utils.RandRange(-0.5, 0.5), 0)
			ctx.SetColor(rclr)
			ctx.Push()
			ctx.RotateAbout(dir.At(x, y), pos.X, pos.Y)
			// brushSize := dir.At(x,y)
			ctx.DrawEllipse(pos.X, pos.Y, 3,5)
			ctx.Pop()
			ctx.Fill()
		}
	}
}
