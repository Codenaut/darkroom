package main

import (
	"os"

	"github.com/codenaut/darkroom/contain"
	"github.com/codenaut/darkroom/square"
	"github.com/codenaut/darkroom/square_crop"
	"github.com/codenaut/darkroom/thumbnail"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "darkroom"

	app.Usage = "Image processing"

	app.Commands = []cli.Command{
		{
			Name:  "thumbnail",
			Usage: "create thumbnail",

			Action: func(ctx *cli.Context) error {
				return thumbnail.CreateThumbnail(ctx.GlobalString("input"), ctx.GlobalString("output"), ctx.String("postfix"), ctx.Int("width"), ctx.Int("height"))

			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "postfix, p",
					Value: "_thumb",
					Usage: "Thumbnail postfix",
				},
				cli.IntFlag{
					Name:  "height",
					Value: 0,
					Usage: "Thumbnail height",
				},
				cli.IntFlag{
					Name:  "width",
					Value: 0,
					Usage: "Thumbnail width",
				},
			},
		},
		{
			Name:  "square",
			Usage: "Make image square",

			Action: func(ctx *cli.Context) error {
				return square.CreateSquare(ctx.GlobalString("input"), ctx.GlobalString("output"), ctx.String("postfix"), ctx.Int("size"))

			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "postfix, p",
					Value: "_thumb",
					Usage: "Output postfix",
				},
				cli.IntFlag{
					Name:  "size",
					Value: 0,
					Usage: "Output size",
				},
			},
		},
		{
			Name:  "contain",
			Usage: "Contain image within given square",

			Action: func(ctx *cli.Context) error {
				return contain.CreateContain(ctx.GlobalString("input"), ctx.GlobalString("output"), ctx.String("postfix"), ctx.Int("size"))
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "postfix, p",
					Value: "_resized",
					Usage: "Output postfix",
				},
				cli.IntFlag{
					Name:  "size",
					Value: 0,
					Usage: "Max width/height",
				},
			},
		},
		{
			Name:  "square_crop",
			Usage: "Resize and crop image to fit within square",

			Action: func(ctx *cli.Context) error {
				return square_crop.CreateSquareCrop(ctx.GlobalString("input"), ctx.GlobalString("output"), ctx.String("postfix"), ctx.Int("size"))
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "postfix, p",
					Value: "_square",
					Usage: "Output postfix",
				},
				cli.IntFlag{
					Name:  "size",
					Value: 0,
					Usage: "Max width/height",
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Value: "",
			Usage: "Input image",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "",
			Usage: "Output image",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
