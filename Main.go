package implot

// #include "wrapper/Main.h"
import "C"

// Version returns a version string, e.g., "0.13 WIP".
func Version() string {
	return C.GoString(C.igpMain_Version())
}
