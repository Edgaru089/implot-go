package implot

// #include <stdlib.h>
// #include "wrapper/Plot.h"
import "C"
import (
	"reflect"
)

//-----------------------------------------------------------------------------
// [SECTION] Plot Items
//-----------------------------------------------------------------------------

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
// Since slices are so versatile, so the Count and Offset parameters are removed.
// You can just slice the data youself, if you have it.
//
// All the values parameters in the PlotXXXAny() functions should only be slices or
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
func PlotLine(label string, values interface{}, xscale, x0 float64) {
	PlotLineV(label, values, xscale, x0)
}

// PlotLineV plots a standard 2D line plot with all parameters.
func PlotLineV(label string, values interface{}, xscale, x0 float64) {
	vd := valueGet(values)
	C.igpPlotLine(wrapString(label), wrapDoubleSlice(vd), C.int(len(vd)), C.double(xscale), C.double(x0))
}

// PlotLineP plots a standard 2D line plot from a slice of points.
func PlotLineP(label string, points []Point) {
	xs, ys, count, stride := wrapPointSlice(points)
	C.igpPlotLineXY(wrapString(label), xs, ys, count, stride)
}

// PlotLineG plots a standard 2D line plot from a series of points obtained from a callback.
func PlotLineG(label string, getter DataGetter, userData interface{}, count int) {
	PlotLineP(label, DataGet(getter, userData, count))
}

// PlotScatter

// PlotStairs

// PlotShaded
