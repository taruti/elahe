package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

func gridAttachLabel(grid *gtk.Grid, name string, left, top int) error {
	lab, err := gtk.LabelNew(name)
	if err != nil {
		return err
	}
	lab.SetMarginStart(2)
	lab.SetMarginEnd(2)
	lab.SetHAlign(gtk.ALIGN_END)
	grid.Attach(lab, left, top, 1, 1)
	return nil
}

func gridAttachEntry(grid *gtk.Grid, left, top int) error {
	ent, err := gtk.EntryNew()
	if err != nil {
		return err
	}
	ent.SetHAlign(gtk.ALIGN_FILL)
	ent.SetHExpand(true)
	ent.Connect("activate", func() {
		wdg, err := grid.GetChildAt(left, top+1)
		if err!=nil {
			return
		}
		wdg.GrabFocus()
	})
	grid.Attach(ent, left, top, 1, 1)
	return nil
}

func gridAttachLabelEntry(grid *gtk.Grid, name string, left, top int) error {
	err := gridAttachLabel(grid, name, left, top)
	if err != nil {
		return err
	}
	return gridAttachEntry(grid, left+1, top)
}

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
	tte, err := gtk.TextTagNew("error")
	if err != nil {
		return nil, err
	}
	tte.SetProperty("underline", pango.UNDERLINE_ERROR)
	ttt.Add(tte)

	tbuf, err := gtk.TextBufferNew(ttt)
	if err != nil {
		return nil, err
	}
	tbuf.Connect("changed", func() { spellCheck(tbuf, tte) })

	tv, err := gtk.TextViewNewWithBuffer(tbuf)
	if err != nil {
		return nil, err
	}
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if err != nil {
		return nil, err
	}
	vbox.PackStart(tv, true, true, 2)

	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}
	grid.SetColumnSpacing(4)

	for i, name := range []string{"Subject", "From", "To"} {
		err = gridAttachLabelEntry(grid, name, 0, i)
		if err != nil {
			return nil, err
		}
	}

	vbox.PackStart(grid, false, false, 2)

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
