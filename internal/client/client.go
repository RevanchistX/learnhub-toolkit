package client

import (
	"bytes"
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/kbinani/screenshot"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func AppClient() {

	//
	//go func() {
	//	window := new(app.Window)
	//	err := loadConfig(window)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	if err := loop(window); err != nil {
	//		log.Fatal(err)
	//	}
	//	os.Exit(0)
	//}()
	app.Main()
}

func loadConfig(window *app.Window) error {
	window.Option(app.Title("LearnHub ToolKit Client"), app.Maximized.Option())
	return nil
}

func loop(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			layoutContext := app.NewContext(&ops, e)
			op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(1, 1))).Add(&ops)
			for i := 0; i < screenshot.NumActiveDisplays(); i++ {
				generatedImage := generateImage(i)
				drawImage(&ops, generatedImage, i, screenshot.GetDisplayBounds(i))
				//TODO Draw on sharer's screen
				//drawBounds(&ops, generatedImage.Bounds())
			}
			e.Frame(layoutContext.Ops)
			window.Invalidate()
		}
	}
}

func drawImage(ops *op.Ops, imageToDraw image.Image, index int, bounds image.Rectangle) {
	imageOp := paint.NewImageOp(imageToDraw)
	imageOp.Filter = paint.FilterNearest
	imageOp.Add(ops)
	//TODO resize image to fit container
	op.Offset(image.Pt(bounds.Dx()/2*index, 0)).Add(ops)
	paint.PaintOp{}.Add(ops)
}

//func generateImage(index int) image.Image {
//	bounds := screenshot.GetDisplayBounds(index)
//	capturedImg, err := screenshot.CaptureRect(bounds)
//	if err != nil {
//		panic(err)
//	}
//	return image.Image(capturedImg)
//}

func generateImage(index int) image.Image {
	resp, err := http.Get("http://localhost:8080/streaming")
	if err != nil {
		log.Fatalln(err)
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	img, _, err := image.Decode(bytes.NewReader(resBody))
	if err != nil {
		log.Fatalln(err)
	}
	return img
}
