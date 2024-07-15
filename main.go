package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/check"
	"ascii/file"
	"ascii/graphics"
)

func main() {
	args := os.Args[1:]
	data := check.Usage(args)

	text := check.Text(data.Text)

	banner := "standard"
	if data.Banner != "" {
		banner = data.Banner
	}
	asciiArtFile := check.ArtFile(banner)
	bannerFileContents := file.ReadArtFile(asciiArtFile)

	inputArgs := strings.Split(text, "\\n")

	output := graphics.Graphic(inputArgs, bannerFileContents, data.Color, data.ColorSubstr)
	fmt.Print(output)
}
