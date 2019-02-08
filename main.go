package main

import (
	"log"
	"os"

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

	app.Action = func(c *cli.Context) error {
		log.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
