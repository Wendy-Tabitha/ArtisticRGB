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
	// Get command-line arguments, excluding the program name.
	args := os.Args[1:]

	// Check the usage of the arguments and parse the input data.
	data := check.Usage(args)

	// Validate that the color substring does not contain newline characters.
	if strings.Contains(data.ColorSubstr, "\n") || strings.Contains(data.ColorSubstr, `\n`) {
		fmt.Printf("Invalid character in color substring\n")
		return
	}

	// Process the input text.
	text := check.Text(data.Text)

	// Determine the banner style, defaulting to "standard".
	banner := "standard"
	if data.Banner != "" {
		banner = data.Banner
	}

	// Validate and read the ASCII art file for the chosen banner.
	asciiArtFile := check.ArtFile(banner)
	bannerFileContents := file.ReadArtFile(asciiArtFile)

	// Split the input text by the newline escape sequence.
	inputArgs := strings.Split(text, "\\n")

	// Generate the ASCII art graphic with the specified parameters.
	output := graphics.Graphic(inputArgs, bannerFileContents, data.Color, data.ColorSubstr)

	// Print the generated ASCII art graphic to the output.
	fmt.Print(output)
}
