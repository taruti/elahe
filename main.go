package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func createMainWindow() (*gtk.Box, error) {
	ttt, err := gtk.TextTagTableNew()
	if err != nil {
		return nil, err
	}
	tbuf, err := gtk.TextBufferNew(ttt)
	if err != nil {
		return nil, err
	}
	tv, err := gtk.TextViewNewWithBuffer(tbuf)
	if err != nil {
		return nil, err
	}
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if err != nil {
		return nil, err
	}
	vbox.PackStart(tv, true, true, 2)
	return vbox, nil
}

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	wdg, err := createMainWindow()
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Add(wdg)

	win.ShowAll()
	gtk.Main()
}
