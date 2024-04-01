package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run . [OPTIONS] [STRING] [BANNER] || Example: go run . \"test\" standard || Options: --output=, --align=")
		return
	}
	argStr := os.Args[1]
	var styleBanner string
	var outputFile string
	var align string
	thirdBanner := false
	var width int

	if len(argStr) >= 8 {
		if argStr[:2] == "--" {
			if argStr[:8] == "--align=" || argStr[:9] == "--output=" {
			} else {
				fmt.Println("Wrong flag. (--output= || --align=)")
				return
			}
		}
		if argStr[:8] == "--align=" {
			width, _, _ = getTerminalSize()
			align = argStr[8:]
			align = strings.ToLower(align)
			if align == "" {
				fmt.Println("Missing align name!")
				return
			} else {
				switch align {
				case "left", "right", "center", "justify":
				default:
					fmt.Println("Wrong align! (right, left, center, justify)")
					return
				}
			}
			argStr = os.Args[2]
			thirdBanner = true
			if strings.Contains(os.Args[2], "--output=") {
				fmt.Println("Can't use output flag and align flag same time!")
				return
			}
		} else if argStr[:9] == "--output=" {
			outputFile = argStr[9:]
			if outputFile == "" {
				fmt.Println("Missing output name!")
				return
			}
			if len(os.Args) < 3 {
				fmt.Println("Missing string!")
				return
			}
			argStr = os.Args[2]
			thirdBanner = true
			if strings.Contains(os.Args[2], "--align=") {
				fmt.Println("Can't use output flag and align flag same time!")
				return
			}
		}
	}

	if len(os.Args) == 2 {
		styleBanner = "standard"
	} else if len(os.Args) == 3 {
		if thirdBanner {
			styleBanner = "standard"
		} else {
			styleBanner = strings.ToLower(os.Args[2])
		}
	} else if len(os.Args) == 4 {
		styleBanner = strings.ToLower(os.Args[3])
	} else {
		fmt.Println("Usage: go run . [OPTIONS] [STRING] [BANNER] || Example: go run . \"test\" standard || Options: --output=, --align=")
		return
	}

	sepArgs := strings.Split(argStr, "\\n")

	file, err := os.ReadFile(styleBanner + ".txt")
	if err != nil {
		fmt.Println(styleBanner + " banner does not exist.")
		return
	}

	lines := strings.Split(string(file), "\n")
	if align != "" {
		printAsciiArtAlign(sepArgs, lines, align, width)
	} else if outputFile != "" {
		createdFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Something went wrong while creating output file.")
		}
		printAsciiArtToFile(sepArgs, lines, createdFile)
	} else {
		printAsciiArt(sepArgs, lines)
	}
}

func printAsciiArtAlign(sentences []string, textFile []string, position string, w int) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}
		wordCount := 1
		for _, char := range word {
			if char == ' ' {
				wordCount++
			}
		}
		wordLen := 0
		for i := 0; i < len(word); i++ {
			for lineIndex, line := range textFile {
				if lineIndex == (int(word[i])-32)*9+2 {
					wordLen += len(line)
					break
				}
			}
		}
		var spacesForJustify int
		if wordCount == 1 && position == "justify" {
			position = "center"
		} else if wordCount == 1 {
			spacesForJustify = (w - wordLen) / wordCount
		} else {
			spacesForJustify = (w - wordLen) / (wordCount - 1)
		}
		spaces := w/2 - wordLen/2
		for h := 1; h < 9; h++ {
			if position == "center" {
				for i := 1; i <= spaces; i++ {
					fmt.Print(" ")
				}
			} else if position == "right" {
				for i := 1; i <= spaces*2; i++ {
					fmt.Print(" ")
				}
			}
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						if position == "justify" && i != len(word)-1 && word[i] == ' ' {
							fmt.Print(line)
							for i := 1; i <= spacesForJustify; i++ {
								fmt.Print(" ")
							}
						} else {
							fmt.Print(line)
						}
						break
					}
				}
			}
			if position == "center" {
				for i := 1; i <= spaces; i++ {
					fmt.Print(" ")
				}
			} else if position == "left" {
				for i := 1; i <= spaces*2; i++ {
					fmt.Print(" ")
				}
			}

			fmt.Println()
		}
	}
}

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	size := strings.Split(string(out), " ")
	width, err := strconv.Atoi(strings.TrimSpace(size[1]))
	if err != nil {
		return 0, 0, err
	}

	height, err := strconv.Atoi(strings.TrimSpace(size[0]))
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

func printAsciiArtToFile(sentences []string, textFile []string, toFile *os.File) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				_, err := toFile.WriteString("\n")
				if err != nil {
					log.Fatal(err)
				}
			}
			continue
		}
		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						_, err := toFile.WriteString(line)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
			_, err := toFile.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	_, err := toFile.WriteString("\n")
	if err != nil {
		log.Fatal(err)
	}
}

func printAsciiArt(sentences []string, textFile []string) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}
		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						fmt.Print(line)
					}
				}
			}
			fmt.Println()
		}
	}
}
