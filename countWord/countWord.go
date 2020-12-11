package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	END_MARK    = "EOF"
	END_MARK_V2 = '^'
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	flush(stdin)

	fmt.Println("input with V1: ")
	contents := getContents(stdin)
	fmt.Println("total number of words: ", getCount(contents))

	flush(stdin)
	fmt.Println("input with V2: ")
	wordsV2 := strings.Fields(getContentsV2(stdin))
	fmt.Println("total number of words V2: ", len(wordsV2))
}

func getContents(stdin *bufio.Reader) string {
	var lines []string
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if strings.Index(line, END_MARK) != -1 {
			break
		}
	}

	return strings.Join(lines, " ")
}

func getContentsV2(stdin *bufio.Reader) string {
	contents, _ := stdin.ReadString(END_MARK_V2)
	return contents[:strings.Index(contents, string(END_MARK_V2))]
}

func getCount(contents string) int {
	var count int
	var previous, current bool
	for _, content := range contents {
		current = unicode.IsSpace(content)
		if previous && !current {
			count++
		}
		previous = current
	}

	if !unicode.IsSpace((rune)(contents[strings.Index(contents, END_MARK)-1])) {
		count++
	}

	return count
}

func flush(stdin *bufio.Reader) {
	for i := 0; i < stdin.Buffered(); i++ {
		stdin.ReadByte()
	}
}
