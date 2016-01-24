package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
	"github.com/taruti/elahe/gtkhelper"
)

var (
	errorTextTag *gtk.TextTag
)

type ComposeWin struct {
	textBuf *gtk.TextBuffer
	overlay *gtkhelper.OverlayWithSpinner
}

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

func gridAttachEntry(grid *gtk.Grid, left, top int, fun func()) error {
	ent, err := gtk.EntryNew()
	if err != nil {
		return err
	}
	ent.SetHAlign(gtk.ALIGN_FILL)
	ent.SetHExpand(true)
	ent.Connect("activate", fun)
	grid.Attach(ent, left, top, 1, 1)
	return nil
}

func gridAttachLabelEntry(grid *gtk.Grid, name string, left, top int, fun func()) error {
	err := gridAttachLabel(grid, name, left, top)
	if err != nil {
		return err
	}
	return gridAttachEntry(grid, left+1, top, fun)
}

func createMainTextView(cw *ComposeWin) (*gtk.TextView, error) {
	ttt, err := gtk.TextTagTableNew()
	if err != nil {
		return nil, err
	}
	ttt.Add(errorTextTag)

	tbuf, err := gtk.TextBufferNew(ttt)
	if err != nil {
		return nil, err
	}
	cw.textBuf = tbuf
	tbuf.Connect("changed", cw.spellCheck)

	tv, err := gtk.TextViewNewWithBuffer(tbuf)
	return tv, err
}

func createMainWindow() (*gtk.Window, error) {
	var err error

	errorTextTag, err = gtk.TextTagNew("error")
	if err != nil {
		return nil, err
	}
	errorTextTag.SetProperty("underline", pango.UNDERLINE_ERROR)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	cw := &ComposeWin{}

	tv, err := createMainTextView(cw)
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
		i := i
		fun := func() {
			wdg, err := grid.GetChildAt(1, i+1)
			if err != nil {
				cw.overlay.StartSpin()
				return
			}
			wdg.GrabFocus()
		}
		err = gridAttachLabelEntry(grid, name, 0, i, fun)
		if err != nil {
			return nil, err
		}
	}

	vbox.PackStart(grid, false, false, 2)

	ov, err := gtkhelper.NewOverlayWithSpinner()
	if err != nil {
		return nil, err
	}
	ov.Add(vbox)
	cw.overlay = ov

	win.Add(&ov.Container.Widget)
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
