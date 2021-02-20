package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

func main() {

	rand.Seed(time.Now().UnixNano() / 2)
	var w, h, n int
	var s string

	flag.IntVar(&w, "w", 400, "Width of the image")
	flag.IntVar(&h, "h", 200, "Height of the image")
	flag.IntVar(&n, "n", 5, "Number of images")
	flag.StringVar(&s, "t", "test", "Title of the image")

	flag.Parse()

	for i := 0; i < n; i++ {
		createImage(w, h, i, s)
	}

}

func createImage(width, height, i int, t string) {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))

	dc := gg.NewContext(width, height)
	dc.SetRGB(r, g, b)
	dc.Clear()

	// add a semi transparent overlay
	margin := 45.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	dc.SetColor(color.White)
	// dc.SetFontFace()

	s := fmt.Sprintf("%s %d", t, i+1)
	x = (float64(dc.Width()) / 2)
	y = (float64(dc.Height()) / 2)
	dc.DrawStringWrapped(s, x, y, 0.5, 0.5, w-20.0, 1.25, gg.AlignCenter)

	dc.SavePNG(fmt.Sprintf("%s_%d.png", strings.ReplaceAll(t, " ", "_"), i+1))
}

func createBanner(t string, i int) {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))
	dc := gg.NewContext(420, 200)
	dc.SetRGB(r, g, b)
	dc.Clear()

	// add a semi transparent overlay
	margin := 45.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	// add font and text
	dc.LoadFontFace("NotoSerif-Bold.ttf", 42)
	dc.SetColor(color.White)
	s := fmt.Sprintf("%s %d", t, i+1)
	x = (float64(dc.Width()) / 2)
	y = (float64(dc.Height()) / 2)
	dc.DrawStringWrapped(s, x, y, 0.5, 0.5, w-20.0, 1.25, gg.AlignCenter)

	dc.SavePNG(fmt.Sprintf("testdata/%s_banner_%d.png", strings.ReplaceAll(t, " ", "_"), i+1))
}
