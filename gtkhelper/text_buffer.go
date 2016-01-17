package gtkhelper

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <string.h>
import "C"
import (
	"github.com/gotk3/gotk3/gtk"
	"reflect"
	"unsafe"
)

func TextIterForwardWordEnd(ti *gtk.TextIter) bool {
	res := C.gtk_text_iter_forward_word_end((*C.GtkTextIter)(unsafe.Pointer(ti)))
	return res == 1
}
func TextIterIsWordStart(ti *gtk.TextIter) bool {
	return C.gtk_text_iter_starts_word((*C.GtkTextIter)(unsafe.Pointer(ti))) == 1
}
func TextIterForwardChar(ti *gtk.TextIter) bool {
	return C.gtk_text_iter_forward_char((*C.GtkTextIter)(unsafe.Pointer(ti)))==1
}
func TextIterWordStart(ti *gtk.TextIter) bool {
	for !TextIterIsWordStart(ti) {
		if !TextIterForwardChar(ti) {
			return false
		}
	}
	return true
}

type GChar C.gchar

func TextBufferRawSlice(tbuf *gtk.TextBuffer) *GChar {
	start := tbuf.GetStartIter()
	end := tbuf.GetEndIter()
	c := C.gtk_text_buffer_get_slice(
		(*C.GtkTextBuffer)(unsafe.Pointer(tbuf.Native())),
		(*C.GtkTextIter)(unsafe.Pointer(start)),
		(*C.GtkTextIter)(unsafe.Pointer(end)),
		gbool(true),
	)
	return (*GChar)(c)
}

func (ptr *GChar) String() string {
	nbytes := int(C.strlen((*C.char)(ptr)))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: uintptr(unsafe.Pointer(ptr)), Len: nbytes}))
}

func (ptr *GChar) Free() {
	C.g_free(C.gpointer(ptr))
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
