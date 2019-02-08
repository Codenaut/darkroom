package square

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func CreateSquare(input string, output string, postfix string, size int) error {
	if fi, err := os.Stat(input); err != nil {
		return err
	} else if fi.Mode().IsDir() {
		if files, err := ioutil.ReadDir(input); err == nil {
			for _, f := range files {
				if err := CreateSquare(filepath.Join(input, f.Name()), "", postfix, size); err != nil {
					return err
				}
			}
			return nil
		} else {
			return fmt.Errorf("Could not read directory: %s (%s)", input, err)
		}
	}

	src, err := imaging.Open(input)
	if err != nil {
		return fmt.Errorf("Could not open file: %s (%s)", input, err)
	}
	if size > 0 {
		src = imaging.Resize(src, 0, size, imaging.Lanczos)
	}
	imgSize := src.Bounds().Size()
	if imgSize.X == imgSize.Y {
		return nil
	}
	maxSize := imgSize.X
	if imgSize.Y > maxSize {
		maxSize = imgSize.Y
	}
	fillColor := color.NRGBA{0xff, 0xff, 0xff, 0xff}
	newImage := imaging.New(maxSize, maxSize, fillColor)
	processed := imaging.OverlayCenter(newImage, src, 1.0)

	dest := output
	if dest == "" {
		ts := strings.SplitN(input, ".", 2)
		dest = ts[0] + postfix
		if len(ts) > 1 {
			dest += "." + ts[1]
		}
	}

	if err := imaging.Save(processed, dest); err != nil {
		return fmt.Errorf("Could not save file: %s (%s)", dest, err)
	} else {
		return nil
	}
}
