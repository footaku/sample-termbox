package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func line(row int, item string) {
	runes := []rune(item)
	for col, r := range runes {
		termbox.SetCell(col, row, r, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func listItems(items []string) {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		panic(err)
	}

	for row, item := range items {
		line(row, item)
	}

	if err := termbox.Flush(); err != nil {
		panic(err)
	}
}

func highlight(next int, max int) (selected int) {
	switch {
	case next < 0:
		selected = max
	case next > max:
		selected = 0
	default:
		selected = next
	}
	width, _ := termbox.Size()

	for row := 0; row <= max; row++ {
		bg := termbox.ColorDefault

		if row == selected {
			bg = termbox.ColorWhite
		}

		for col := 0; col < width; col++ {
			char := termbox.CellBuffer()[(width*row)+col].Ch
			termbox.SetCell(col, row, char, termbox.ColorDefault, bg)
		}
	}
	return
}

func get(row int) string {
	var chars []rune
	width, _ := termbox.Size()
	for col := 0; col < width; col++ {
		chars = append(chars, termbox.CellBuffer()[(width*row)+col].Ch)
	}
	return string(chars)
}

func pollEvent() (result string) {
	currentRow := 0
	items := []string{"aaa", "bbbb", "ccccc", "d", "ee"}

	listItems(items)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
				currentRow = highlight(currentRow-1, len(items)-1)
			case termbox.KeyArrowDown:
				currentRow = highlight(currentRow+1, len(items)-1)
			case termbox.KeyEnter:
				result = get(currentRow)
				return
			}
		default:
			listItems(items)
		}

		if err := termbox.Flush(); err != nil {
			panic(err)
		}
	}
}

func start() (string, bool) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	switch result := pollEvent(); result {
	case "":
		return result, false
	default:
		return result, true
	}
}

func main() {
	result, _ := start()
	fmt.Print(result)
}
