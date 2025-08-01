// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.White,                        // 背景色
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // 红
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, // 橙
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // 黄
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // 绿
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // 蓝
	color.RGBA{0x4B, 0x00, 0x82, 0xFF}, // 靛
	color.RGBA{0x8B, 0x00, 0xFF, 0xFF}, // 紫
}

const (
	whiteIndex = 0 // first color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		colorIndex := uint8(i%(len(palette)-1) + 1)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
