package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file := getFileContents()
	filteredFile := filterFileContents(file)
	blocks := splitIntoBlocks(filteredFile)
	formattedBlocks := filterBlocks(blocks)
	fmt.Println(formattedBlocks[0])
}

func filterBlocks(blocks []string) []string {
	var newBlocks []string
	for _, block := range blocks {
		if len(block) > 5 {
			newBlocks = append(newBlocks, formatBlock(block))
		}
	}
	return newBlocks
}

func formatBlock(block string) string {
	lines := strings.Split(block, " ")
	var (
		pair string
		pairs []string
	)
	if len(lines) > 1 {
		if !strings.Contains(lines[0], ":") { 
			lines[1] = fmt.Sprintf("%s %s", lines[0], lines[1]) 
			lines = lines[1:]
		}
	}
	for _, line := range lines {
		if strings.Contains(line, ":") {
			pairs = append(pairs, pair)
			pair = ""
			pair += line
		} else {
			pair += " " + line
		}
	}
	return strings.Join(pairs, "\n")
}

func splitIntoBlocks(file string) []string {
	lines := strings.Split(file, "\n")
	var (
		block  string
		blocks []string
	)
	for _, line := range lines {
		if len(line) > 0 {
			if []rune(line)[0] != ' ' {
				blocks = append(blocks, block)
				block = ""
				block += line
			} else {
				block += line
			}
		}
	}
	return blocks
}

func filterFileContents(file string) string {
	lines := strings.Split(file, "\n")
	newStr := ""
	for _, line := range lines {
		if len(line) > 0 {
			if []rune(line)[0] != '#' {
				newStr += line + "\n"
			}
		}
	}
	return newStr
}

func getFileContents() string {
	file, err := os.ReadFile("../languages-on-github.yml")
	if err != nil {
		fmt.Println("There was a problem reading the file")
		os.Exit(1)
	}
	return string(file)
}
