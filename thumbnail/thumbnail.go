package thumbnail

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func CreateThumbnail(input string, output string, postfix string, width int, height int) error {
	fi, err := os.Stat(input)
	if err != nil {
		return err
	}
	if fi.Mode().IsDir() {
		if files, err := ioutil.ReadDir(input); err == nil {
			for _, f := range files {
				if err := CreateThumbnail(filepath.Join(input, f.Name()), "", postfix, width, height); err != nil {
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
	thumbnail := imaging.Resize(src, width, height, imaging.Lanczos)

	dest := output
	if dest == "" {
		ts := strings.SplitN(input, ".", 2)
		dest = ts[0] + postfix
		if len(ts) > 1 {
			dest += "." + ts[1]
		}
	}

	if err := imaging.Save(thumbnail, dest); err != nil {
		return fmt.Errorf("Could not save file: %s (%s)", dest, err)
	} else {
		return nil
	}
}
