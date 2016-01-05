package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func createMainWindow() (*gtk.Window, error) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

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

	win.Add(vbox)
	return win, nil
}

func main() {
	gtk.Init(nil)

	win, err := createMainWindow()
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.ShowAll()
	gtk.Main()
}
