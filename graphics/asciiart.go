package graphics

import (
	"fmt"
	"os"
	"strings"
)

// This function is used to show the graphics by calculating the value of each line and returning the string version of the graphic.
func Graphic(str, banner []string, colorString, colorSubstr string) string {
	var lines int
	result := ""

	color, ok := ParseColor(colorString)
	if !ok && colorString != "" {
		fmt.Printf("Invalid color: %q\n", colorString)
		os.Exit(1)
	}

	doColor := func(i int, lineString string) {
		for _, runechar := range lineString {
			index := (runechar-32)*9 + rune(i)
			if colorString == "" {
				result += banner[index]
			} else {
				result += fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", color.R, color.G, color.B, banner[index])
			}
		}
	}

	for _, lineString := range str {
		if lineString == "" {
			lines = 1
		} else {
			lines = 8
		}

		if colorSubstr == "" {
			for i := 1; i <= lines; i++ {
				doColor(i, lineString)
				result += "\n"
			}
			continue
		}

		for i := 1; i <= lines; i++ {
			substrSplits := strings.Split(lineString, colorSubstr)
			for j, split := range substrSplits {
				for _, runechar := range split {
					index := (runechar-32)*9 + rune(i)
					result += banner[index]
				}

				if j != len(substrSplits)-1 {
					doColor(i, colorSubstr)
				}

			}
			result += "\n"
		}

	}
	return result
}

type Color struct {
	R, G, B int
}

// This functions contains the colors to be used when adding a color to a string
func ParseColor(s string) (Color, bool) {
	colorMap := map[string]Color{
		"red":       {255, 0, 0},
		"green":     {0, 255, 0},
		"blue":      {0, 0, 255},
		"black":     {0, 0, 0},
		"white":     {255, 255, 255},
		"yellow":    {255, 255, 0},
		"cyan":      {0, 255, 255},
		"magenta":   {255, 0, 255},
		"gray":      {128, 128, 128},
		"maroon":    {128, 0, 0},
		"olive":     {128, 128, 0},
		"purple":    {128, 0, 128},
		"teal":      {0, 128, 128},
		"navy":      {0, 0, 128},
		"silver":    {192, 192, 192},
		"lime":      {0, 255, 0},
		"orange":    {255, 165, 0},
		"brown":     {165, 42, 42},
		"darkgreen": {0, 100, 0},
		"gold":      {255, 215, 0},
	}

	color, ok := colorMap[s]
	return color, ok
}
