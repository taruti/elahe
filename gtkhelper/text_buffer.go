package gtkhelper

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"
import (
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

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

func (ptr *GChar) Free() {
	C.g_free(C.gpointer(ptr))
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
