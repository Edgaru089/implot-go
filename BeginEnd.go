package implot

// #include <stdlib.h>
// #include "wrapper/BeginEnd.h"
import "C"
import (
	"unsafe"

	"github.com/inkyblackness/imgui-go/v4"
)

// BeginPlot starts a 2D plotting context with only the title specified.
// It calls BeginPlotV(title, Vec2{-1, 0}, Flags_None).
//
// If this function returns true, then EndPlot() MUST be called! You can do:
//     if igwrap.BeginPlot(...) {
//         igwrap.PlotXXX(...)
//         ...
//         igwrap.EndPlot()
//     }
//
// Note that:
//  - title must be unique to the current ImGui ID scope. If you need
//    to avoid ID collisions or don't want to display a title in the plot,
//    use double hashes (e.g. "MyPlot##HiddenIdText" or "##NoTitle").
//  - size is the **frame** size of the plot widget, not the plot area.
//    The default size of plots (i.e. when ImVec2(0,0)) can be modified
//    in your ImPlotStyle.
func BeginPlot(title string) bool {
	return BeginPlotV(title, imgui.Vec2{X: -1, Y: 0}, Flags_None)
}

// BeginPlotV starts a 2D plotting context with all the parameters specified.
//
// If this function returns true, then EndPlot() MUST be called! You can do:
//     if igwrap.BeginPlot(...) {
//         igwrap.PlotXXX(...)
//         ...
//         igwrap.EndPlot()
//     }
//
// Note that:
//  - title must be unique to the current ImGui ID scope. If you need
//    to avoid ID collisions or don't want to display a title in the plot,
//    use double hashes (e.g. "MyPlot##HiddenIdText" or "##NoTitle").
//  - size is the **frame** size of the plot widget, not the plot area.
//    The default size of plots (i.e. when ImVec2(0,0)) can be modified
//    in your ImPlotStyle.
func BeginPlotV(title string, size imgui.Vec2, flags Flags) bool {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	return bool(C.igpBeginPlot(ctitle, wrapVec2(size), C.igpFlags(flags)))
}

// EndPlot marks the end of an active plot.
//
// Only call EndPlot() if BeginPlot() returns true! Typically called at the end
// of an if statement conditioned on BeginPlot(). See example above.
func EndPlot() {
	C.igpEndPlot()

	// Discard temp data
	// from Setup.go:69
	for k := range axisFormatCb {
		delete(axisFormatCb, k)
	}
	for _, f := range axisEndPlotCb {
		f()
	}
	axisEndPlotCb = axisEndPlotCb[0:0]
}

// BeginSubplots starts a subdivided plotting context with onle the required parameters.
// It calls BeginSubplotsV(title, rows, cols, Vec2{-1, 0}, SubplotFlags_None, nil, nil).
//
// If the function returns true, EndSubplots() MUST be called! Call BeginPlot/EndPlot
// AT MOST [rows*cols] times in between the begining and end of the subplot context.
// Plots are added in row major order. Example:
//
//   if BeginSubplots("My Subplot",2,3) {
//       for i := 0; i < 6; i++ {
//           if BeginPlot(...) {
//               PlotLine(...);
//               ...
//               EndPlot();
//           }
//       }
//       EndSubplots();
//   }
//
// Produces:
//
//   [0] | [1] | [2]
//   ----|-----|----
//   [3] | [4] | [5]
//
// Important notes:
//
//  - #title must be unique to the current ImGui ID scope. If you need to avoid ID
//    collisions or don't want to display a title in the plot, use double hashes
//    (e.g. "MySubplot##HiddenIdText" or "##NoTitle").
//  - #rows and #cols must be greater than 0.
//  - #size is the size of the entire grid of subplots, not the individual plots
//  - #row_ratios and #col_ratios must have AT LEAST #rows and #cols elements,
//    respectively. These are the sizes of the rows and columns expressed in ratios.
//    If the user adjusts the dimensions, the arrays are updated with new ratios.
//
// Important notes regarding BeginPlot from inside of BeginSubplots:
//
//  - The #title parameter of _BeginPlot_ (see above) does NOT have to be
//    unique when called inside of a subplot context. Subplot IDs are hashed
//    for your convenience so you don't have call PushID or generate unique title
//    strings. Simply pass an empty string to BeginPlot unless you want to title
//    each subplot.
//  - The #size parameter of _BeginPlot_ (see above) is ignored when inside of a
//    subplot context. The actual size of the subplot will be based on the
//    #size value you pass to _BeginSubplots_ and #row/#col_ratios if provided.
func BeginSubplots(title string, rows, cols int) bool {
	return BeginSubplotsV(title, rows, cols, imgui.Vec2{X: -1, Y: 0}, SubplotFlags_None, nil, nil)
}

// BeginSubplotsV starts a subdivided plotting context with all parameters.
//
// If the function returns true, EndSubplots() MUST be called! Call BeginPlot/EndPlot
// AT MOST [rows*cols] times in between the begining and end of the subplot context.
// Plots are added in row major order. Example:
//
//   if BeginSubplots("My Subplot",2,3) {
//       for i := 0; i < 6; i++ {
//           if BeginPlot(...) {
//               PlotLine(...);
//               ...
//               EndPlot();
//           }
//       }
//       EndSubplots();
//   }
//
// Produces:
//
//   [0] | [1] | [2]
//   ----|-----|----
//   [3] | [4] | [5]
//
// Important notes:
//
//  - #title must be unique to the current ImGui ID scope. If you need to avoid ID
//    collisions or don't want to display a title in the plot, use double hashes
//    (e.g. "MySubplot##HiddenIdText" or "##NoTitle").
//  - #rows and #cols must be greater than 0.
//  - #size is the size of the entire grid of subplots, not the individual plots
//  - #row_ratios and #col_ratios must have AT LEAST #rows and #cols elements,
//    respectively. These are the sizes of the rows and columns expressed in ratios.
//    If the user adjusts the dimensions, the arrays are updated with new ratios.
//
// Important notes regarding BeginPlot from inside of BeginSubplots:
//
//  - The #title parameter of _BeginPlot_ (see above) does NOT have to be
//    unique when called inside of a subplot context. Subplot IDs are hashed
//    for your convenience so you don't have call PushID or generate unique title
//    strings. Simply pass an empty string to BeginPlot unless you want to title
//    each subplot.
//  - The #size parameter of _BeginPlot_ (see above) is ignored when inside of a
//    subplot context. The actual size of the subplot will be based on the
//    #size value you pass to _BeginSubplots_ and #row/#col_ratios if provided.
func BeginSubplotsV(title string, rows, cols int, size imgui.Vec2, flags SubplotFlags, rowRatios, colRatios []float32) bool {
	var rf, cf *C.float
	if rowRatios != nil && len(rowRatios) >= rows {
		rf = (*C.float)(unsafe.Pointer(&rowRatios[0]))
	}
	if colRatios != nil && len(colRatios) >= cols {
		cf = (*C.float)(unsafe.Pointer(&colRatios[0]))
	}

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	return bool(C.igpBeginSubplots(ctitle, C.int(rows), C.int(cols), wrapVec2(size), C.igpSubplotFlags(flags), rf, cf))
}

// EndSubplots marks the end of a subdivided plotting area.
//
// Only call EndSubplots() if BeginSubplots() returns true! Typically called at the end
// of an if statement conditioned on BeginSublots(). See example above.
func EndSubplots() {
	C.igpEndSubplots()
}
