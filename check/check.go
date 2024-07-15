package check

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"ascii/types"
)

// Usage checks if the program was supplied the expected command-line arguments, exits on usage error.
func Usage(args []string) types.Data {
	if len(args) == 0 {
		PrintUsage()
	} else if len(args) > 4 {
		PrintUsage()
	}
	var out types.Data

	if len(args) == 1 {
		inValid, _ := Expressions(args[0])
		if inValid {
			PrintUsage()
		} else {
			out.Text = args[0]
		}
	} else if len(args) == 2 {
		isValid, color := Expressions(args[0])
		if isValid {
			out.Color = color
			out.Text = args[1]
		} else {
			PrintUsage()
		}
	} else if len(args) == 3 {
		isValid, color := Expressions(args[0])
		if isValid {
			out.Color = color
			out.ColorSubstr = args[1]
			out.Text = args[2]
		} else {
			PrintUsage()
		}
	} else if len(args) == 4 {
		isValid, color := Expressions(args[0])
		if isValid {
			out.Color = color
			out.ColorSubstr = args[1]
			out.Text = args[2]
			out.Banner = args[3]
		} else {
			PrintUsage()
		}
	}
	return out
}

// PrintUsage prints the program usage information, then exits the program with the error code 1.
func PrintUsage() {
	usage := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
	fmt.Println(usage)
	os.Exit(1)
}

func Text(text string) string {
	text = strings.ReplaceAll(text, "\n", "\\n")

	if text == "" {
		os.Exit(1)
	} else if text == "\\n" {
		fmt.Println()
		os.Exit(1)
	}

	out := ""
	for _, char := range text {
		if char > '~' {
			fmt.Println("Error: Non-ASCII character found! Cannot display the graphic representation")
			os.Exit(1)
		} else if char >= ' ' {
			out += string(char)
		}
	}

	return out
}

func ArtFile(banner string) string {
	switch banner {
	case "standard":
		return "standard.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	case "shadow":
		return "shadow.txt"
	case "lean":
		return "lean.txt"
	default:
		fmt.Printf("Invalid banner: %q\n", banner)
		PrintUsage()
	}

	return ""
}

// Expressions is used to check if the flags to be matched are the same.
func Expressions(s string) (bool, string) {
	re := regexp.MustCompile(`^--color=(.+)`)
	matches := re.FindStringSubmatch(s)
	if matches == nil {
		return false, ""
	}
	return true, matches[1]
}
