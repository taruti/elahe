package main

import (
	"log"
	"unicode"

	"github.com/gotk3/gotk3/gtk"
	"github.com/taruti/elahe/gtkhelper"
	"github.com/taruti/enchant"
)

var dictionaries = loadSpellCheckers("en_US", "fi")

func loadSpellCheckers(languages ...string) []enchant.Dict {
	brk, err := enchant.NewEnchant()
	if err != nil {
		log.Printf("Failed to load enchant: %v", err)
		return nil
	}
	var ds []enchant.Dict
	for _, l := range languages {
		d, err := brk.LoadDict(l)
		if err != nil {
			log.Printf("Failed to load dictionary for %q: %v", l, err)
			continue
		}
		ds = append(ds, d)
	}
	return ds
}

// spellCheck is called only from the main gtk thread - thus no concurrency protection
// needed.
func spellCheck(tb *gtk.TextBuffer) {
	log.Println("Spellcheck begin")
	gptr := gtkhelper.TextBufferRawSlice(tb)
	if gptr == nil {
		return
	}
	defer gptr.Free()

	str := gptr.String()
	if str == "" {
		return
	}

	// Eat leading non-letters
	for idx, ch := range str {
		if unicode.IsLetter(ch) {
			str = str[idx:]
			break
		}
	}

	start := 0
	for idx, ch := range str {
		if !unicode.IsLetter(ch) {
			if start >= 0 {
				spellWord(str[start:idx])
				start = -1
			}
		} else if(start < 0) {
			start = idx
		}
	}
	// Skip trailing words by purpose
}

func spellWord(word string) {
	log.Println("word", word)
}
