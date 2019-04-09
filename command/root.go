package command

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"strings"

	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// RootFlag is flag list for root command
var RootFlag = []cli.Flag{

	cli.StringFlag{
		Name:  "format, f",
		Usage: "image `format`. jpg or png",
	},
	cli.IntFlag{
		Name:  "width",
		Usage: "image width by pixel",
		Value: 200,
	},
	cli.IntFlag{
		Name:  "height",
		Usage: "image height by pixel",
		Value: 100,
	},
}

// DoGenerate is a Action for root command
func DoGenerate(c *cli.Context) error {

	outfile := c.Args().First()
	if outfile == "" {
		return errors.New("no output file name specified")
	}

	fFormat := c.String("format")
	format, err := detectImageFormat(fFormat, outfile)
	if err != nil {
		return err
	}

	width := c.Int("width")
	height := c.Int("height")

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.Transparent)
		}
	}

	f, err := os.Create(outfile)
	if err != nil {
		return errors.Wrap(err, "failed to create file.")
	}

	switch format {
	case "png":
		png.Encode(f, img)
	case "jpg":
		jpeg.Encode(f, img, &jpeg.Options{100})
	}

	return nil
}

func detectImageFormat(fFormat, outfile string) (string, error) {
	if fFormat != "" {
		if fFormat == "jpg" || fFormat == "png" {
			return fFormat, nil
		}
		return "", errors.New("specified format is illegal")
	}

	s := strings.Split(outfile, ".")
	i := len(s) - 1

	switch s[i] {
	case "png", "PNG":
		return "png", nil
	case "jpg", "jpeg", "JPG", "JPEG":
		return "jpg", nil
	}

	return "", errors.New("Failed to determine image format from file-name")
}
