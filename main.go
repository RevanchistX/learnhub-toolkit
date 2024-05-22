package main


import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/kbinani/screenshot"
	"image"
	"log"
	"os"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			for i := 0; i < screenshot.NumActiveDisplays(); i++ {
				img := generateImage(i)
				drawImage(&ops, img)
			}
			e.Frame(gtx.Ops)
			w.Invalidate()
		}
	}
}

func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Filter = paint.FilterNearest
	imageOp.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func generateImage(index int) image.Image {
	bounds := screenshot.GetDisplayBounds(index)
	capturedImg, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	return image.Image(capturedImg)
}
