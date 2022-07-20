package implot

// #include "plot_wrapper/Contexts.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/inkyblackness/imgui-go/v4"
)

// Context specifies a scope; a global state for ImPlot.
type Context struct {
	handle C.igpContext
}

// CreateContext creates a new ImPlot context.
// It should be called right after imgui.CreateContext().
func CreateContext() *Context {
	return &Context{handle: C.igpCreateContext()}
}

// ErrNoContext is used when no context is current.
// It should not be the same with imgui.ErrNoContext.
var ErrNoContext = errors.New("no current context")

// CurrentContext returns the currently active context.
// Returns ErrNoContext if no context is available.
func CurrentContext() (*Context, error) {
	raw := C.igpCurrentContext()
	if raw == nil {
		return nil, ErrNoContext
	}
	return &Context{handle: raw}, nil
}

// ErrContextDestroyed is returned when trying to use an already destroyed context.
// It should not be the same with imgui.ErrContextDestroyed.
var ErrContextDestroyed = errors.New("context is destroyed")

// Destroy deletes the context.
// Destroying an already destroyed context does nothing.
func (c *Context) Destroy() {
	if c.handle != nil {
		C.igpDestroyContext(c.handle)
		c.handle = nil
	}
}

// SetCurrent activates this context as the current active one.
func (c *Context) SetCurrent() error {
	if c.handle == nil {
		return ErrContextDestroyed
	}
	C.igpSetCurrentContext(c.handle)
	return nil
}

// SetImGUIContext sets the current ImGUI context for ImPlot.
//
// It should not be used, at least with the current way Go links libraries.
func SetImGUIContext(ig *imgui.Context) {
	C.igpSetImGUIContext(C.iggContext(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ig)))))
}
