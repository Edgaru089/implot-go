package implot

// #include <stdlib.h>
// #include "wrapper/Setup.h"
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

//-----------------------------------------------------------------------------
// [SECTION] Setup
//-----------------------------------------------------------------------------
//
// The following API allows you to setup and customize various aspects of the
// current plot. The functions should be called immediately after BeginPlot
// and before any other API calls. Typical usage is as follows:
//
// if (BeginPlot(...)) {                     1) begin a new plot
//     SetupAxis(ImAxis_X1, "My X-Axis");    2) make Setup calls
//     SetupAxis(ImAxis_Y1, "My Y-Axis");
//     SetupLegend(ImPlotLocation_North);
//     ...
//     SetupFinish();                        3) [optional] explicitly finish setup
//     PlotLine(...);                        4) plot items
//     ...
//     EndPlot();                            5) end the plot
// }
//
// Important notes:
//
// - Always call Setup code at the top of your BeginPlot conditional statement.
// - Setup is locked once you start plotting or explicitly call SetupFinish.
//   Do NOT call Setup code after you begin plotting or after you make
//   any non-Setup API calls (e.g. utils like PlotToPixels also lock Setup)
// - Calling SetupFinish is OPTIONAL, but probably good practice. If you do not
//   call it yourself, then the first subsequent plotting or utility function will
//   call it for you.

// Formatter is a callback for axis tick label formatting.
type Formatter func(val float64, userData interface{}) string

// SetupAxis enables an axis or sets the label and/or flags for an existing axis.
func SetupAxis(axis Axis, label string, flags AxisFlags) {
	if len(label) == 0 {
		C.igpSetupAxis(C.igpAxis(axis), nil, C.igpAxisFlags(flags))
	} else {
		C.igpSetupAxis(C.igpAxis(axis), wrapString(label), C.igpAxisFlags(flags))
	}
}

// SetupAxisLimits sets an axis range limits.
// If ImPlotCond_Always is used, the axes limits will be locked.
//
// Note that SetupAxisLinks() is absent. I don't really know how to implement that.
func SetupAxisLimits(axis Axis, vmin, vmax float64, cond Condition) {
	C.igpSetupAxisLimits(C.igpAxis(axis), C.double(vmin), C.double(vmax), C.igpCondition(cond))
}

// SetupAxisFormat sets the format of numeric axis labels via formater specifier (default="%g").
// The formatted value will be C.double, and you can also use %f.
func SetupAxisFormat(axis Axis, fmt string) {
	C.igpSetupAxisFormat(C.igpAxis(axis), wrapString(fmt))
}

// For use by SetupAxisFormatCallback and callbacks
var (
	// Callbacks. MUST be cleaned AFTER every EndPlot().
	axisFormatCb map[uintptr]struct {
		fmt      Formatter
		userData interface{}
	} = make(map[uintptr]struct {
		fmt      Formatter
		userData interface{}
	})
	// Callbacks after every EndPlot(), for cleaning.
	axisEndPlotCb []func()
)

func addEndPlotCb(f func()) {
	axisEndPlotCb = append(axisEndPlotCb, f)
}

//export igpgoAxisFormatCb
func igpgoAxisFormatCb(value float64, buf *byte, size C.int, cbid uintptr) {
	// Construct a ByteSlice for conveience
	bhead := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(buf)),
		Len:  int(size),
		Cap:  int(size),
	}
	b := *((*[]byte)(unsafe.Pointer(bhead)))

	cb, ok := axisFormatCb[cbid]
	if !ok {
		panic(fmt.Errorf("igpgoAxisFormatCb() called with invalid callback ID (%d)", cbid))
	}
	copy(b, cb.fmt(value, cb.userData)+"\x00")
}

// SetupAxisFormatCallback sets the format of numeric axis labels via formatter callback.
//
// The userData value will be discarded on every EndPlot, so hopefully this will not
// cause a memory leak.
func SetupAxisFormatCallback(axis Axis, formatter Formatter, userData interface{}) {
	cbid := uintptr(len(axisFormatCb) + 1)
	axisFormatCb[cbid] = struct {
		fmt      Formatter
		userData interface{}
	}{fmt: formatter, userData: userData}

	C.igpSetupAxisFormatCallback(C.igpAxis(axis), C.uintptr_t(cbid))
}

// SetupAxisTickValues set an axis' tick values (as given in the slice) and labels.
// Labels are optional and can be set default with labels=nil.
//
// To keep the default ticks, set keep_default=true.
//
// Note that if len(values)!=len(labels), it takes len(values).
func SetupAxisTickValues(axis Axis, values []float64, labels []string, keepDefaults bool) {
	dsp, fin := wrapDoubleSliceAlloc(values)
	addEndPlotCb(fin)
	sp, fin := wrapStringSlice(labels)
	addEndPlotCb(fin)
	C.igpSetupAxisTickValues(C.igpAxis(axis), dsp, C.int(len(values)), sp, C.bool(keepDefaults))
}

// SetupAxisTickRange set an axis' tick values (n of them from [vmin, vmax]) and labels.
// Labels are optional and can be set default with labels=nil.
//
// To keep the default ticks, set keep_default=true.
func SetupAxisTickRange(axis Axis, vmin, vmax float64, n int, labels []string, keepDefaults bool) {
	sp, fin := wrapStringSlice(labels)
	addEndPlotCb(fin)
	C.igpSetupAxisTickRange(C.igpAxis(axis), C.double(vmin), C.double(vmax), C.int(n), sp, C.bool(keepDefaults))
}

// SetupAxes sets the label and/or flags for primary X and Y axes.
// (shorthand for two calls to SetupAxis)
func SetupAxes(xlabel, ylabel string, xflags, yflags AxisFlags) {
	C.igpSetupAxes(wrapString(xlabel), wrapString(ylabel), C.igpAxisFlags(xflags), C.igpAxisFlags(yflags))
}

// SetupAxesLimits sets the primary X and Y axes range limits.
// If ImPlotCond_Always is used, the axes limits will be locked.
// (shorthand for two calls to SetupAxisLimits)
func SetupAxesLimits(xmin, xmax, ymin, ymax float64, cond Condition) {
	C.igpSetupAxesLimits(C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.igpCondition(cond))
}

// SetupLegend sets up the position and flags of the plot legend.
func SetupLegend(location Location, flags LegendFlags) {
	C.igpSetupLegend(C.igpLocation(location), C.igpLegendFlags(flags))
}

// SetupMouseText sets up location of the current plot's mouse position text
// (the tiny xxx,yyy numbers on the plot).
// The default is South|East (so bottom-right).
func SetupMouseText(location Location, flags MouseTextFlags) {
	C.igpSetupMouseText(C.igpLocation(location), C.igpMouseTextFlags(flags))
}

// SetupFinish explicitly finalize plot setup. Once you call
// this, you cannot make anymore Setup calls for the current plot!
//
// Note that calling this function is OPTIONAL; it will be called
// by the first subsequent setup-locking API call.
func SetupFinish() {
	C.igpSetupFinish()
}
