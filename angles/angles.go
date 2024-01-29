package angles

import (
	"image"
	"math"

	"github.com/disintegration/gift"
	"github.com/setanarut/pointilizm/matrix"
)

// GetAngles get vector field angles
func GetAngles(im image.Image) matrix.Mat {
	var sigma float32 = 30
	normalize := false
	size := im.Bounds().Max

	//kernels
	ScharrX := gift.New(
		gift.Convolution(
			[]float32{
				3, 0, -3,
				10, 1, -10,
				3, 0, -3,
			},
			normalize, false, false, 0,
		),
		gift.GaussianBlur(sigma),
	)
	ScharrY := gift.New(
		gift.Convolution(
			[]float32{
				3, 10, 3,
				0, 0, 0,
				-3, -10, -3,
			},
			normalize, false, false, 0,
		),
		gift.GaussianBlur(sigma),
	)
	//Apply Scharr filters
	ScharrYImage := image.NewGray16(ScharrY.Bounds(im.Bounds()))
	ScharrXImage := image.NewGray16(ScharrX.Bounds(im.Bounds()))
	ScharrY.Draw(ScharrYImage, im)
	ScharrX.Draw(ScharrXImage, im)

	// angle matrix
	angles := matrix.NewMatrix(size.X, size.Y)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {

			dx := ScharrXImage.Gray16At(x, y).Y
			dy := ScharrYImage.Gray16At(x, y).Y

			// calculate vector angle
			dir := math.Atan2(float64(dx), float64(dy))

			angles.Set(x,y, dir)
		}
	}


	return angles
}
