package main

import (
	"log"

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
func (cw *ComposeWin) spellCheck() {
	if len(dictionaries) == 0 {
		return
	}

	i0 := cw.textBuf.GetStartIter()
	i1 := cw.textBuf.GetStartIter()
	idx := 0
	stats := make([]int32, len(dictionaries))
	for gtkhelper.TextIterWordStart(i0) && gtkhelper.TextIterForwardWordEnd(i1) {
		word, err := cw.textBuf.GetText(i0, i1, true)
		if err != nil {
			log.Println("GetText:", err)
			return
		}
		for i, d := range dictionaries {
			if d.Check(word) {
				stats[i]++
			}
		}
		// Only handle first chars when trying to determine language
		if idx >= 31 {
			break
		}
		idx++
		gtkhelper.TextIterForwardChar(i0)
	}
	best := 0
	for i, v := range stats[1:] {
		if v > stats[best] {
			best = i + 1
		}
	}
	dict := dictionaries[best]

	// Remove all existing tags
	cw.textBuf.RemoveTag(errorTextTag, cw.textBuf.GetStartIter(), cw.textBuf.GetEndIter())

	i0 = cw.textBuf.GetStartIter()
	i1 = cw.textBuf.GetStartIter()
	for gtkhelper.TextIterWordStart(i0) && gtkhelper.TextIterForwardWordEnd(i1) {
		word, err := cw.textBuf.GetText(i0, i1, true)
		if err != nil {
			log.Println("GetText:", err)
			return
		}
		if !dict.Check(word) {
			cw.textBuf.ApplyTag(errorTextTag, i0, i1)
		}
		gtkhelper.TextIterForwardChar(i0)
	}
}
