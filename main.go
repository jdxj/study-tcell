package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

var (
	defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(0x606367)
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	err = s.Init()
	if err != nil {
		panic(err)
	}

	s.SetStyle(boxStyle)

	s.Show()
	time.Sleep(time.Hour)

}
