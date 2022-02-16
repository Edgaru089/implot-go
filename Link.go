package implot

// #cgo CPPFLAGS: -DIMGUI_DISABLE_WIN32_DEFAULT_IME_FUNCTIONS
// #cgo CXXFLAGS: -std=c++11
// #cgo CXXFLAGS: -Wno-subobject-linkage
// #cgo linux && amd64 LDFLAGS:  -L./implot -limplot -limgui
// #cgo linux && amd64 CPPFLAGS: -DIMPLOT_GO_PREBUILT
import "C"
