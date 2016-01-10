package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/taruti/enchant"
)

func loadSpellCheckers(languages ...string) []enchant.Dict {
	brk, err := enchant.NewEnchant()
	if err != nil {
		log.Print("Failed to load enchant: %v", err)
		return nil
	}
	var ds []enchant.Dict
	for _, l := range languages {
		d, err := brk.LoadDict(l)
		if err != nil {
			log.Print("Failed to load dictionary for %q: %v", l, err)
			continue
		}
		ds = append(ds, d)
	}
	return ds
}

func spellCheck(tb *gtk.TextBuffer) {
}
