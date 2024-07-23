# Ascii-art-color
* Ascii-art-color is a program written in Go that takes a string as input and outputs a graphic representation of the string using ASCII characters in different colors.
* The output should manipulate colors using the flag "--color=thecolor", "substring" and "string to be colored", in which --color is the flag and "thecolor" is the color desired by the user and "substring" is the substring that you can choose to be colored from the string. 
* Colors can be achieved using different notations (color code systems, like RGB, hsl, ANSI...), it is up to you to choose which one you want to use.
* In the project below we implemented the RGB notation which also uses RGB hexadecimal color codes.

## Features
- Colors can be achieved using different notions i.e  RGB, hsl, ANSI...
- The substring to be coloured can be a single letter or more.
- If the substring is not specified the whole string should be colored.
- The flag must have exactly the same format as mentioned above, any other formats must return the following usage message:
```bash
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<thecolor> <substring to be colored> "something"
or
  go run . --color=<thecolor>  "something"
```
- Additionally, the program must still be able to run with a single [STRING] argument. EX: go run . [STRING]

## Instructions to run locally

To clone this repository, copy the command below on your terminal:

```bash
git clone https://github.com/Wendy-Tabitha/ArtisticRGB
```

Go to the project directory
```bash
cd ascii-art-color
```
## Usage
- To run the program, use the command below;
```bash
go run . [OPTION] [STRING]
```
EX: go run . --color=<color> <substring to be colored> "something"

- Theprogram can also run in this instances below;
```bash
go run . [OPTION] [SUBSTRING] [STRING] [BANNER]

go run . [OPTION] [SUBSTRING] [STRING]
```

- The program also runs with a different flag notation by using a hexadecimal representation of a color in the RGB and the hexadecimal must be seven characters i.e. for example the color code below represents red.The RGB colors should be written in lowercase letters only whereas the RGB hex color codes can use either capital or small letters.

```bash
go run . --color=yellow world
or
go run . --color=#FF0000 hello
or
go run . --color=#ff0000 hello
```

## RGB Colors to be parsed at the flag are the ones below
> red, green, blue, black, white, yellow, cyan, magenta, gray, maroon, olive, purple, teal, navy, silver, lime, orange, brown, darkgreen and gold.

> For the RGB hex color codes you can get them here: [RGBhexcolorcodes](https://www.rapidtables.com/web/color/RGB_Color.html)

## Banner Format
* Each character in the ASCII representation has a height of 8 lines.
* Characters are separated by a newline character \n.
* Incase of an empty string such as "\n" it takes a height of 1 line.

The program uses only standard.txt banner format for graphical representation using ASCII characters:

>  The program automatically takes the banner format of `standard`.

### Running Tests
To run unit tests, navigate to the project directory and run the following command:
```bash
go test -v
```

### Formatting program
To format the program, navigate to the project directory and run the following command:
```bash
gofmt -w -s .
```

## AUTHORS
- [Wendy-Tabitha](https://github.com/Wendy-Tabitha)