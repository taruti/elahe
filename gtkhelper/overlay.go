package gtkhelper

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <string.h>
import "C"
import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"runtime"
	"unsafe"
)

type OverlayWithSpinner struct {
	gtk.Container
	spin *gtk.Spinner
}

func NewOverlayWithSpinner() (*OverlayWithSpinner, error) {
	spin, err := gtk.SpinnerNew()
	if err!=nil {
		return nil, err
	}
	overlayPtr := unsafe.Pointer(C.gtk_overlay_new())
	C.gtk_overlay_add_overlay((*C.GtkOverlay)(overlayPtr), spinWPtr(spin))
	obj := wrapObject(overlayPtr)
	return &OverlayWithSpinner{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}},spin}, nil
}

func (o *OverlayWithSpinner) StartSpin() {
	C.gtk_overlay_reorder_overlay(overlayPtr(o), spinWPtr(o.spin), -1)
	o.spin.Start()
}

func (o *OverlayWithSpinner) StopSpin() {
	o.spin.Stop()
	C.gtk_overlay_reorder_overlay(overlayPtr(o), spinWPtr(o.spin), 0)
}

func spinWPtr(ptr *gtk.Spinner) *C.GtkWidget {
	return (*C.GtkWidget)(unsafe.Pointer(ptr.Native()))
}

func overlayPtr(ptr *OverlayWithSpinner) *C.GtkOverlay {
	return (*C.GtkOverlay)(unsafe.Pointer(ptr.Container.Native()))
}

// Wrapper function for new objects with reference management.
func wrapObject(ptr unsafe.Pointer) *glib.Object {
	obj := &glib.Object{glib.ToGObject(ptr)}

	if obj.IsFloating() {
		obj.RefSink()
	} else {
		obj.Ref()
	}

	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return obj
}
