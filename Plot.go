package implot

// #include <stdlib.h>
// #include "plot_wrapper/Plot.h"
import "C"
import (
	"math"
	"reflect"
	"unsafe"
)

//-----------------------------------------------------------------------------
// [SECTION] Plot Items
//-----------------------------------------------------------------------------
//
// The main plotting API is provied below. Call these functions between
// Begin/EndPlot and after any Setup API calls. Each plots data on the current
// x and y axes, which can be changed with `SetAxis/Axes`.
//
// The templated functions are explicitly instantiated in implot_items.cpp.
// They are not intended to be used generically with custom types. You will get
// a linker error if you try! All functions support the following scalar types:
//
// float, double, ImS8, ImU8, ImS16, ImU16, ImS32, ImU32, ImS64, ImU64
//
//
// If you need to plot custom or non-homogenous data you have a few options:
//
// 1. If your data is a simple struct/class (e.g. Vector2f), you can use striding.
//    This is the most performant option if applicable.
//
//    struct Vector2f { float X, Y; };
//    ...
//    Vector2f data[42];
//    ImPlot::PlotLine("line", &data[0].x, &data[0].y, 42, 0, sizeof(Vector2f)); // or sizeof(float)*2
//
// 2. Write a custom getter C function or C++ lambda and pass it and optionally your data to
//    an ImPlot function post-fixed with a G (e.g. PlotScatterG). This has a slight performance
//    cost, but probably not enough to worry about unless your data is very large. Examples:
//
//    ImPlotPoint MyDataGetter(void* data, int idx) {
//        MyData* my_data = (MyData*)data;
//        ImPlotPoint p;
//        p.x = my_data->GetTime(idx);
//        p.y = my_data->GetValue(idx);
//        return p
//    }
//    ...
//    auto my_lambda = [](void*, int idx) {
//        double t = idx / 999.0;
//        return ImPlotPoint(t, 0.5+0.5*std::sin(2*PI*10*t));
//    };
//    ...
//    if (ImPlot::BeginPlot("MyPlot")) {
//        MyData my_data;
//        ImPlot::PlotScatterG("scatter", MyDataGetter, &my_data, my_data.Size());
//        ImPlot::PlotLineG("line", my_lambda, nullptr, 1000);
//        ImPlot::EndPlot();
//    }
//
// NB: All types are converted to double before plotting. You may lose information
// if you try plotting extremely large 64-bit integral types. Proceed with caution!

// Note only for the Go binding:
//
// For each type of ImPlot::PlotXXX/PlotXXXG, five functions are presented:
// PlotXXX & PlotXXXV plots a slice of any number (integer or float),
// PlotXXXXY plots separate X/Y slices, PlotXXXP plots a slice of points,
// and PlotXXXG plots a set of points from a given getter.
// They construct the data on the fly with no callbacks.
//
// Since slices are so versatile, the Count and Offset parameters are removed.
// You can just slice the data youself, if you have it.
//
// All the values parameters in the PlotXXX/PlotXXXV() functions should only be slices or
// arrays to numeric types, which is then converted to float64. If it is not, it panics.

// Callback signature for the data getter.
//
// It is called from within PlotXXXG() (the word "within" is subject to change),
// with idx ranging from 0 to N-1.
type DataGetter func(userData interface{}, idx int) Point

// DataGet generates a slice of Points from a given DataGetter.
//
// This puts a lot of stress on the allocator/GC so perhaps we need something else.
// This still can be useful for the end user however.
func DataGet(getter DataGetter, userData interface{}, count int) (ps []Point) {
	ps = make([]Point, count)
	for i := 0; i < count; i++ {
		ps[i] = getter(userData, i)
	}
	return
}

// valueGet converts any slice/array of numeric type to []float64.
// Again this is bad for allocator/GC so we should need something else.
func valueGet(values interface{}) (result []float64) {
	// fast path
	if f64s, ok := values.([]float64); ok {
		return f64s
	}

	// good old reflect
	ref := reflect.ValueOf(values)
	switch ref.Kind() {
	case reflect.Array, reflect.Slice:
		n := ref.Len()
		if n == 0 {
			return
		}
		result = make([]float64, n)

		// construct the values one-by-one
		for i := 0; i < n; i++ {
			e := ref.Index(i)
			switch e.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				result[i] = float64(e.Int())
			case reflect.Float32, reflect.Float64:
				result[i] = e.Float()
			default:
				panic("PlotXXX called with a slice/array containing non-scalar values")
			}
		}
	default:
		panic("PlotXXX called with a non-slice, non-array value")
	}
	return
}

// PlotLine

// PlotLine plots a standard 2D line plot with minimal parameters.
// It calls PlotLineV(label, values, 1, 0).
func PlotLine(label string, values interface{}) {
	PlotLineV(label, values, 1, 0)
}

// PlotLineV plots a standard 2D line plot with all parameters.
func PlotLineV(label string, values interface{}, xscale, x0 float64) {
	vd := valueGet(values)
	C.igpPlotLine(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(xscale), C.double(x0))
}

// PlotLineP plots a standard 2D line plot from a slice of points.
func PlotLineP(label string, points []Point) {
	xp, yp, count, stride := wrapPointSlice(points)
	C.igpPlotLineXY(wrapString(label), xp, yp, count, stride)
}

// PlotLineXY plots a standard 2D line plot from slices of X/Y coords.
func PlotLineXY(label string, xs, ys interface{}) {
	xp, yp, count, stride := wrapXYSlice(valueGet(xs), valueGet(ys))
	C.igpPlotLineXY(wrapString(label), xp, yp, count, stride)
}

// PlotLineG plots a standard 2D line plot from a series of points obtained from a callback.
func PlotLineG(label string, getter DataGetter, userData interface{}, count int) {
	PlotLineP(label, DataGet(getter, userData, count))
}

// PlotScatter

// PlotScatter plots a standard 2D scatter plot with minimal parameters.
// It calls PlotScatterV(label, values, 1, 0).
//
// Default marker is ImPlotMarker_Circle.
func PlotScatter(label string, values interface{}) {
	PlotScatterV(label, values, 1, 0)
}

// PlotScatterV plots a standard 2D scatter plot with all parameters.
//
// Default marker is ImPlotMarker_Circle.
func PlotScatterV(label string, values interface{}, xscale, x0 float64) {
	vd := valueGet(values)
	C.igpPlotScatter(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(xscale), C.double(x0))
}

// PlotScatterP plots a standard 2D scatter plot from a slice of points.
//
// Default marker is ImPlotMarker_Circle.
func PlotScatterP(label string, points []Point) {
	xp, yp, count, stride := wrapPointSlice(points)
	C.igpPlotScatterXY(wrapString(label), xp, yp, count, stride)
}

// PlotScatterXY plots a standard 2D scatter plot from slices of X/Y coords.
//
// Default marker is ImPlotMarker_Circle.
func PlotScatterXY(label string, xs, ys interface{}) {
	xp, yp, count, stride := wrapXYSlice(valueGet(xs), valueGet(ys))
	C.igpPlotScatterXY(wrapString(label), xp, yp, count, stride)
}

// PlotScatterG plots a standard 2D scatter plot from a series of points obtained from a callback.
//
// Default marker is ImPlotMarker_Circle.
func PlotScatterG(label string, getter DataGetter, userData interface{}, count int) {
	PlotScatterP(label, DataGet(getter, userData, count))
}

// PlotStairs

// PlotStairs plots a stairstep graph with minimal parameters.
// It calls PlotStairsV(label, values, 1, 0).
//
// The y value is continued constantly from every x position,
// i.e. the interval [x[i], x[i+1]) has the value y[i].
func PlotStairs(label string, values interface{}) {
	PlotStairsV(label, values, 1, 0)
}

// PlotStairsV plots a stairstep graph with all parameters.
// Default marker is ImPlotMarker_Circle.
//
// The y value is continued constantly from every x position,
// i.e. the interval [x[i], x[i+1]) has the value y[i].
func PlotStairsV(label string, values interface{}, xscale, x0 float64) {
	vd := valueGet(values)
	C.igpPlotStairs(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(xscale), C.double(x0))
}

// PlotStairsP plots a stairstep graph from a slice of points.
//
// The y value is continued constantly from every x position,
// i.e. the interval [x[i], x[i+1]) has the value y[i].
func PlotStairsP(label string, points []Point) {
	xp, yp, count, stride := wrapPointSlice(points)
	C.igpPlotStairsXY(wrapString(label), xp, yp, count, stride)
}

// PlotStairsXY plots a stairstep graph from slices of X/Y coords.
//
// The y value is continued constantly from every x position,
// i.e. the interval [x[i], x[i+1]) has the value y[i].
func PlotStairsXY(label string, xs, ys interface{}) {
	xp, yp, count, stride := wrapXYSlice(valueGet(xs), valueGet(ys))
	C.igpPlotStairsXY(wrapString(label), xp, yp, count, stride)
}

// PlotStairsG plots a stairstep graph from a series of points obtained from a callback.
//
// The y value is continued constantly from every x position,
// i.e. the interval [x[i], x[i+1]) has the value y[i].
func PlotStairsG(label string, getter DataGetter, userData interface{}, count int) {
	PlotStairsP(label, DataGet(getter, userData, count))
}

// PlotShaded
// PlotShadedRef

// PlotShadedRef plots a shaded (filled) region between a line and a horizontal reference.
// It calls PlotShadedV(label, values, 0, 1, 0).
func PlotShadedRef(label string, values interface{}) {
	PlotShadedRefV(label, values, 0, 1, 0)
}

// PlotShadedRefV plots a shaded (filled) region between a line and a horizontal reference.
//
// Set yref to +/-INFINITY for infinite fill extents.
func PlotShadedRefV(label string, values interface{}, yref, xscale, x0 float64) {
	vd := valueGet(values)
	C.igpPlotShadedRef(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(yref), C.double(xscale), C.double(x0))
}

// PlotShadedRefP plots a shaded (filled) region between a line and a horizontal reference.
//
// Set yref to +/-INFINITY for infinite fill extents.
func PlotShadedRefP(label string, points []Point, yref float64) {
	xp, yp, count, stride := wrapPointSlice(points)
	C.igpPlotShadedRefXY(wrapString(label), xp, yp, C.double(yref), count, stride)
}

// PlotShadedRefXY plots a shaded (filled) region between a line and a horizontal reference.
//
// Set yref to +/-INFINITY for infinite fill extents.
func PlotShadedRefXY(label string, xs, ys interface{}, yref float64) {
	xp, yp, count, stride := wrapXYSlice(valueGet(xs), valueGet(ys))
	C.igpPlotShadedRefXY(wrapString(label), xp, yp, C.double(yref), count, stride)
}

// PlotShadedRefG plots a shaded (filled) region between a line and a horizontal reference.
//
// Set yref to +/-INFINITY for infinite fill extents.
func PlotShadedRefG(label string, getter DataGetter, userData interface{}, count int, yref float64) {
	PlotShadedRefP(label, DataGet(getter, userData, count), yref)
}

// PlotShadedLines

// PlotShadedLines plots a shaded (filled) region between two lines, without the lines themselves.
// It calls PlotShadedLinesV(label, vs0, vs1, 1, 0).
func PlotShadedLines(label string, vs0, vs1 interface{}) {
	PlotShadedLinesV(label, vs0, vs1, 1, 0)
}

// PlotShadedLinesV plots a shaded (filled) region between two lines, without the lines themselves.
func PlotShadedLinesV(label string, vs0, vs1 interface{}, xscale, x0 float64) {
	vd0, vd1 := valueGet(vs0), valueGet(vs1)
	n := minint(len(vd0), len(vd1))
	// Constuct the X coords
	xs := make([]float64, n)
	for i := 0; i < n; i++ {
		xs[i] = x0 + (float64)(i)*xscale
	}
	PlotShadedLinesXY(label, xs, vd0, vd1)
}

// PlotShadedLinesXY plots a shaded (filled) region between two lines, without the lines themselves.
func PlotShadedLinesXY(label string, xs, ys1, ys2 interface{}) {
	xd, yd1, yd2 := valueGet(xs), valueGet(ys1), valueGet(ys2)
	C.igpPlotShadedLinesXY(
		wrapString(label),
		wrapDoubleSlice(xd),
		wrapDoubleSlice(yd1),
		wrapDoubleSlice(yd2),
		C.int(minint(len(xd), minint(len(yd1), len(yd2)))),
		C.int(unsafe.Sizeof(float64(0))),
	)
}

// PlotShadedLinesG plots a shaded (filled) region between two lines, without the lines themselves.
//
// The X component of the second getter is discarded.
func PlotShadedLinesG(label string, get1 DataGetter, data1 interface{}, get2 DataGetter, data2 interface{}, count int) {
	xs, ys1, ys2 := make([]float64, count), make([]float64, count), make([]float64, count)
	for i := 0; i < count; i++ {
		pt1 := get1(data1, i)
		xs[i], ys1[i] = pt1.X, pt1.Y
		pt2 := get2(data2, i)
		_, ys2[i] = pt2.X, pt2.Y
	}

	PlotShadedLinesXY(label, xs, ys1, ys2)
}

// PlotBars

// PlotBars plots a vertical bar graph, with every bar centering at X coords 0, 1, ..., N-1.
func PlotBars(label string, vs interface{}) {
	PlotBarsV(label, vs, 0.67, 0)
}

// PlotBarsV plots a vertical bar graph, with bars centering at
// x0, x0+1, x0+2, ... x0+N-1, Each taking up a fraction of the
// available width. #barWidth should be in (0, 1].
func PlotBarsV(label string, vs interface{}, barWidth, x0 float64) {
	vd := valueGet(vs)
	C.igpPlotBars(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(barWidth), C.double(x0))
}

// PlotBarsP plots a vertical bar graph, with bars each taking up a
// fraction of the available width. #barWidthFraction should be in (0, 1].
func PlotBarsP(label string, ps []Point, barWidth float64) {
	xp, yp, count, stride := wrapPointSlice(ps)
	C.igpPlotBarsXY(wrapString(label), xp, yp, count, C.double(barWidth), stride)
}

// PlotBarsXY plots a vertical bar graph, with bars each taking up a
// fraction of the available width. #barWidth should be in (0, 1].
func PlotBarsXY(label string, vx, vy interface{}, barWidth float64) {
	xp, yp, count, stride := wrapXYSlice(valueGet(vx), valueGet(vy))
	C.igpPlotBarsXY(wrapString(label), xp, yp, count, C.double(barWidth), stride)
}

// PlotBarsG plots a vertical bar graph, with bars each taking up a
// fraction of the available width. #barWidth should be in (0, 1].
func PlotBarsG(label string, getter DataGetter, userData interface{}, count int, barWidth float64) {
	PlotBarsP(label, DataGet(getter, userData, count), barWidth)
}

// PlotBarsH plots a horizontal bar graph, with every bar centering at Y coords 0, 1, ..., N-1.
func PlotBarsH(label string, vs interface{}) {
	PlotBarsHV(label, vs, 0.67, 0)
}

// PlotBarsHV plots a horizontal bar graph, with bars centering at
// y0, y0+1, y0+2, ... y0+N-1, Each taking up a fraction of the
// available height. #barHeight should be in (0, 1].
func PlotBarsHV(label string, vs interface{}, barHeight, y0 float64) {
	vd := valueGet(vs)
	C.igpPlotBarsH(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(barHeight), C.double(y0))
}

// PlotBarsHP plots a horizontal bar graph, with bars each taking up a
// fraction of the available height. #barHeight should be in (0, 1].
func PlotBarsHP(label string, ps []Point, barHeight float64) {
	xp, yp, count, stride := wrapPointSlice(ps)
	C.igpPlotBarsHXY(wrapString(label), xp, yp, count, C.double(barHeight), stride)
}

// PlotBarsHXY plots a horizontal bar graph, with bars each taking up a
// fraction of the available height. #barHeight should be in (0, 1].
func PlotBarsHXY(label string, vx, vy interface{}, barHeight float64) {
	xp, yp, count, stride := wrapXYSlice(valueGet(vx), valueGet(vy))
	C.igpPlotBarsHXY(wrapString(label), xp, yp, count, C.double(barHeight), stride)
}

// PlotBarsHG plots a horizontal bar graph, with bars each taking up a
// fraction of the available height. #barHeight should be in (0, 1].
func PlotBarsHG(label string, getter DataGetter, userData interface{}, count int, barHeight float64) {
	PlotBarsHP(label, DataGet(getter, userData, count), barHeight)
}

// PlotBarGroups plots a group of vertical bars.
//
// The I-th item in the J-th group is in #values[I][J].
// The I-th item has a legend label of #itemLabels[I].
//
// Item count  N = Min(len(itemLabels), len(values)).
// Group count M = Min(len(values[0]), len(values[1]), ... len(values[N-1])).
//
// The bar groups are centered at at x0, x0+1, x0+2, x0+M-1.
// If you want to put labels on the groups, use SetupAxisTickValues.
func PlotBarGroups(itemLabels []string, values [][]float64, groupWidth, x0 float64, flags BarGroupsFlags) {
	n, m := minint(len(itemLabels), len(values)), math.MaxInt
	for _, s := range values {
		m = minint(m, len(s))
	}

	// Construct the matrix
	stride := unsafe.Sizeof(C.double(0))
	vp := C.malloc(C.size_t(uintptr(n*m) * stride))
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			*((*float64)(unsafe.Add(vp, stride*uintptr(i*m+j)))) = values[i][j]
		}
	}
	addEndPlotCb(func() { C.free(vp) })

	// Copy the labels
	vplabels, fin := wrapStringSlice(itemLabels)
	addEndPlotCb(fin)

	// Make the call
	C.igpPlotBarGroups(vplabels, (*C.double)(vp), C.int(n), C.int(m), C.double(groupWidth), C.double(x0), C.igpBarGroupsFlags(flags))
}

// PlotBarGroupsH plots a group of horizontal bars.
//
// The I-th item in the J-th group is in #values[I][J].
// The I-th item has a legend label of #itemLabels[I].
//
// Item count  N = Min(len(itemLabels), len(values)).
// Group count M = Min(len(values[0]), len(values[1]), ... len(values[N-1])).
//
// The bar groups are centered at at y0, y0+1, y0+2, y0+M-1.
// If you want to put labels on the groups, use SetupAxisTickValues.
func PlotBarGroupsH(itemLabels []string, values [][]float64, groupWidth, y0 float64, flags BarGroupsFlags) {
	n, m := minint(len(itemLabels), len(values)), math.MaxInt
	for _, s := range values {
		m = minint(m, len(s))
	}

	// Construct the matrix
	stride := unsafe.Sizeof(C.double(0))
	vp := C.malloc(C.size_t(uintptr(n*m) * stride))
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			*((*float64)(unsafe.Add(vp, stride*uintptr(i*m+j)))) = values[i][j]
		}
	}
	addEndPlotCb(func() { C.free(vp) })

	// Copy the labels
	vplabels, fin := wrapStringSlice(itemLabels)
	addEndPlotCb(fin)

	// Make the call
	C.igpPlotBarGroupsH(vplabels, (*C.double)(vp), C.int(n), C.int(m), C.double(groupWidth), C.double(y0), C.igpBarGroupsFlags(flags))
}
